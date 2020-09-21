package main

import (
	"context"
	"testing"
)

func TestCreateTable(t *testing.T) {
	td := t.TempDir()

	dsn := "file:" + td + "/TestCreateTable.db"

	_, err := NewSqlite(context.Background(), dsn)
	if err != nil {
		t.Errorf("TestCreateTable NewSqlite: %w", err)
	}
}
