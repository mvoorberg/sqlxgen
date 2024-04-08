package store

import (
	_ "embed"
	"path"
	"testing"

	"github.com/bradleyjkemp/cupaloy"
	"github.com/mvoorberg/sqlxgen/internal/utils/writer"
	"github.com/stretchr/testify/assert"
)

func TestPackage_Generate(t *testing.T) {
	tmpDir := t.TempDir()

	genDir := path.Join(tmpDir, "gen/tmdb_pg/store")

	mw := writer.NewMemoryWriters()

	opts := map[string]string{
		"mysqlModelBanner":    "This is my MySql banner",
		"postgresModelBanner": "This is my Postgres banner",
	}

	storePackage, err := NewPackage(
		mw.Creator,
		"github.com/mvoorberg/sqlxgen/gen/tmdb_pg/store",
		genDir,
		opts,
	)

	if err != nil {
		t.Fatalf("unexpected error: %s", err.Error())
	}

	err = storePackage.Generate()

	assert.Nil(t, err)

	assert.Equal(t, 1, len(mw.Writers))

	pen := mw.Writers[0]

	assert.Equal(t, path.Join(genDir, "store.gen.go"), pen.FullPath)

	cupaloy.SnapshotT(t, pen.Content)
}

func TestNewPackage(t *testing.T) {
	tmpDir := t.TempDir()

	genDir := path.Join(tmpDir, "gen/tmdb_pg/store")

	mw := writer.NewMemoryWriters()

	opts := map[string]string{
		"mysqlModelBanner":    "This is my MySql banner",
		"postgresModelBanner": "This is my Postgres banner",
	}

	got, err := NewPackage(
		mw.Creator,
		"github.com/mvoorberg/sqlxgen/gen/tmdb_pg/store",
		genDir,
		opts,
	)

	assert.Nil(t, err)

	want := Package{
		WriterCreator: mw.Creator,
		PackageName:   "store",
		PackageDir:    "github.com/mvoorberg/sqlxgen/gen/tmdb_pg/store",
		GenDir:        genDir,
		Options:       opts,
	}

	assert.Equal(t, want.PackageName, got.PackageName)

	assert.Equal(t, want.PackageDir, got.PackageDir)

	assert.Equal(t, want.GenDir, got.GenDir)
}
