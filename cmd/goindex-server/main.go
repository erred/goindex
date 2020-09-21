package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/api/metric"
	"go.seankhliao.com/goindex"
	"go.seankhliao.com/usvc"

	_ "modernc.org/sqlite"
)

var (
	exampleSQLiteDSN = `file:test.db?cache=shared&mode=memory`
	IndexURL         = "ttps://index.golang.org/index"
)

func main() {
	var s Server

	srvc := usvc.DefaultConf(&s)
	s.log = srvc.Logger()

	ctx := context.Background()
	err := s.setup(ctx)
	if err != nil {
		s.log.Fatal().Err(err).Msg("setup database")
	}

	s.addedvers = metric.Must(global.Meter(os.Args[0])).NewInt64Counter(
		"added_module_versions",
		metric.WithDescription("versions added to index"),
	)
	s.requests = metric.Must(global.Meter(os.Args[0])).NewInt64Counter(
		"requests",
		metric.WithDescription("total requests"),
	)

	_, grpcServer, run, err := srvc.Server(nil)
	if err != nil {
		s.log.Fatal().Err(err).Msg("setup server")
	}

	// register
	goindex.RegisterGoindexService(grpcServer, &goindex.GoindexService{
		Versions: s.Versions,
	})

	err = run(ctx)
	if err != nil {
		s.log.Fatal().Err(err).Msg("run server")
	}
}

type Server struct {
	sqliteDSN string
	sqlite    *Sqlite

	requests  metric.Int64Counter
	addedvers metric.Int64Counter

	tick time.Duration

	log zerolog.Logger
}

func (s *Server) RegisterFlags(fs *flag.FlagSet) {
	fs.StringVar(&s.sqliteDSN, "sqlite", "", exampleSQLiteDSN)
	fs.DurationVar(&s.tick, "tick", 1*time.Minute, "background update tick")
}

func (s *Server) setup(ctx context.Context) error {
	var err error
	s.sqlite, err = NewSqlite(ctx, s.sqliteDSN)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) background(ctx context.Context) {
	err := s.updateIndex(ctx)
	if err != nil {
		s.log.Error().Err(err).Msg("background init index")
	}
	for range time.NewTicker(s.tick).C {
		err := s.updateIndex(ctx)
		if err != nil {
			s.log.Error().Err(err).Msg("background update index")
		}
	}

}

func (s *Server) updateIndex(ctx context.Context) error {
	ts, err := s.sqlite.LatestTS(ctx)
	if err != nil {
		return fmt.Errorf("updateIndex: %w", err)
	}
	prev := 2000
	for prev == 2000 {
		u := IndexURL
		if ts != "" {
			u += "?since=" + ts
		}
		r, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
		if err != nil {
			return fmt.Errorf("updateIndex create request: %w", err)
		}
		res, err := http.DefaultClient.Do(r)
		if err != nil {
			return fmt.Errorf("updateIndex get: %w", err)
		} else if res.StatusCode != 200 {
			return fmt.Errorf("updateIndex status: %d %s", res.StatusCode, res.Status)
		}
		prev = 0
		d := json.NewDecoder(res.Body)
		for d.More() {
			var ir goindex.IndexRecord
			err = d.Decode(&ir)
			if err != nil {
				return fmt.Errorf("updateIndex decode: %w", err)
			}
			err = s.sqlite.AddVersion(ctx, &ir)
			if err != nil {
				return fmt.Errorf("updateIndex save: %w", err)
			}
			s.addedvers.Add(ctx, 1)
			prev++
			ts = ir.Timestamp
		}
		s.log.Info().Int("versions", prev).Msg("updateIndex added")
	}
	return nil
}

func (s *Server) Versions(ctx context.Context, vr *goindex.VersionsRequest) (*goindex.ProjectVersions, error) {
	pv, err := s.sqlite.AllVersions(ctx, vr.Project, vr.Semver)
	if err != nil {
		err = fmt.Errorf("query versions: %w", err)
		s.log.Error().Err(err).Str("project", vr.Project).Bool("semver", vr.Semver).Msg("versions")
		return nil, err
	}
	s.requests.Add(ctx, 1)
	return pv, nil
}
