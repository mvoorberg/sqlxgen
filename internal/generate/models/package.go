package models

import (
	"os"
	"path/filepath"

	"github.com/joomcode/errorx"
	"github.com/mvoorberg/sqlxgen/internal/generate/types"
	"github.com/mvoorberg/sqlxgen/internal/introspect"
	"github.com/mvoorberg/sqlxgen/internal/utils/casing"
	"github.com/mvoorberg/sqlxgen/internal/utils/writer"
)

type Package struct {
	ModelTemplate    string
	StorePackageDir  string
	StorePackageName string
	PackageName      string
	PackageDir       string
	GenDir           string
	Models           []model
	Options          map[string]string
}

func (p Package) Generate() error {
	err := os.MkdirAll(p.GenDir, 0755)

	if err != nil {
		return errorx.InitializationFailed.Wrap(err, "unable to create directory for models")
	}

	for _, m := range p.Models {
		err := m.generate(p.ModelTemplate, p.PackageName, p.GenDir, p.Options)

		if err != nil {
			return err
		}
	}

	return nil
}

func NewPackage(
	writerCreator writer.Creator,
	translate types.Translate,
	storePackageDir string,
	storePackageName string,
	packageDir string,
	genDir string,
	tables []introspect.Table,
	options map[string]string,
) (Package, error) {
	parentDir := filepath.Base(packageDir)

	packageName, err := casing.SnakeCase(parentDir)

	if err != nil {
		return Package{}, errorx.IllegalState.Wrap(err, "unable to generate package name")
	}

	models := make([]model, len(tables))

	for index, table := range tables {
		m, err := newModel(
			writerCreator,
			translate,
			storePackageDir,
			storePackageName,
			table,
		)

		if err != nil {
			return Package{}, errorx.IllegalState.Wrap(err, "unable to generate model from table")
		}

		models[index] = m
	}

	p := Package{
		ModelTemplate:    translate.ModelTemplate(),
		StorePackageDir:  storePackageDir,
		StorePackageName: storePackageName,
		PackageDir:       packageDir,
		PackageName:      packageName,
		Models:           models,
		GenDir:           genDir,
		Options:          options,
	}

	return p, nil
}
