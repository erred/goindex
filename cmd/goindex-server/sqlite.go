package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"io"

	"golang.org/x/mod/module"

	"go.seankhliao.com/goindex"
)

type Sqlite struct {
	db *sql.DB

	latestTS    *sql.Stmt
	addVersion  *sql.Stmt
	allVersions *sql.Stmt
}

func NewSqlite(ctx context.Context, dsn string) (*Sqlite, error) {
	var s Sqlite
	var err error
	s.db, err = sql.Open("sqlite", dsn)
	if err != nil {
		return nil, fmt.Errorf("NewSqlite open %s: %w", dsn, err)
	}

	err = s.db.PingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("NewSqlite ping: %w", err)
	}

	_, err = s.db.ExecContext(ctx, sqliteTable)
	if err != nil {
		return nil, fmt.Errorf("NewSqlite create table: %w", err)
	}
	_, err = s.db.ExecContext(ctx, sqliteIndexTS)
	if err != nil {
		return nil, fmt.Errorf("NewSqlite create index ts: %w", err)
	}
	_, err = s.db.ExecContext(ctx, sqliteIndexProject)
	if err != nil {
		return nil, fmt.Errorf("NewSqlite create index project: %w", err)
	}

	s.latestTS, err = s.db.PrepareContext(ctx, sqliteLatestTimestamp)
	if err != nil {
		return nil, fmt.Errorf("NewSqlite prepare stmt latestTs: %w", err)
	}
	s.addVersion, err = s.db.PrepareContext(ctx, sqliteAddVersion)
	if err != nil {
		return nil, fmt.Errorf("NewSqlite prepare stmt addVersion: %w", err)
	}
	s.allVersions, err = s.db.PrepareContext(ctx, sqliteAllVersions)
	if err != nil {
		return nil, fmt.Errorf("NewSqlite prepare stmt allVersions: %w", err)
	}

	return &s, nil
}

func (s *Sqlite) LatestTS(ctx context.Context) (string, error) {
	row := s.latestTS.QueryRowContext(ctx)
	err := row.Err()
	if errors.Is(err, sql.ErrNoRows) {
		return "", nil
	} else if err != nil {
		return "", fmt.Errorf("LatestTS query: %w", err)
	}

	var ts string
	err = row.Scan(&ts)
	if err != nil {
		return "", fmt.Errorf("LatestTS scan: %w", err)
	}
	return ts, nil
}

func (s *Sqlite) AddVersion(ctx context.Context, ir *goindex.IndexRecord) error {
	project, _, ok := module.SplitPathVersion(ir.Path)
	if !ok {
		project = ir.Path
	}

	// semver 0 = false, 1 = true
	var major, minor, patch, semver int
	var add string
	_, err := fmt.Sscanf(ir.Version, "v%d.%d.%d%s", &major, &minor, &patch, &add)
	if errors.Is(err, io.EOF) {
		semver = 1
	} else if err != nil {
		return fmt.Errorf("AddVersion parse version %s: %w", ir.Version, err)
	}

	_, err = s.addVersion.ExecContext(ctx, ir.Timestamp, project, ir.Path, ir.Version, semver, major, minor, patch)
	if err != nil {
		return fmt.Errorf("AddVersion exec: %w", err)
	}
	return nil
}

func (s *Sqlite) AllVersions(ctx context.Context, project string, semver bool) (*goindex.ProjectVersions, error) {
	var sv int
	if semver {
		sv = 1
	}

	rows, err := s.allVersions.QueryContext(ctx, project, sv)
	if err != nil {
		return nil, fmt.Errorf("AllVersions query project=%v semver=%v: %w", project, sv, err)
	}
	defer rows.Close()

	var irs []*goindex.IndexRecord
	for rows.Next() {
		var ir goindex.IndexRecord
		err = rows.Scan(&ir.Path, &ir.Version)
		if err != nil {
			return nil, fmt.Errorf("AllVersions scan: %w", err)
		}
		irs = append(irs, &ir)
	}

	pv := &goindex.ProjectVersions{
		Project:  project,
		Versions: irs,
	}
	return pv, nil
}

var (
	sqliteTable = `
CREATE TABLE IF NOT EXISTS goindex (
        timestamp       TEXT,
        project         TEXT,
        module          TEXT,
        version         TEXT,
        semver          INTEGER,
        major           INTEGER,
        minor           INTEGER,
        patch           INTEGER
)
`
	sqliteIndexTS = `
CREATE INDEX IF NOT EXISTS goindex_ts ON goindex (timestamp)`
	sqliteIndexProject = `
CREATE INDEX IF NOT EXISTS goindex_project ON goindex (project, semver)`
	sqliteLatestTimestamp = `
SELECT timestamp
FROM goindex
ORDER BY timestamp DESC
LIMIT 1
`

	sqliteAddVersion = `
INSERT INTO goindex (timestamp, project, module, version, semver, major, minor, patch)
VALUES (?, ?, ?, ?, ?, ?, ?, ?)
`

	sqliteAllVersions = `
SELECT module, version
FROM goindex
WHERE project = ? AND semver = ?
ORDER BY major, minor, patch, version
`
)
