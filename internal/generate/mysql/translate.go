package mysql

import (
	_ "embed"

	"github.com/mvoorberg/sqlxgen/internal/generate/types"
	"github.com/mvoorberg/sqlxgen/internal/introspect"
)

type Mysql struct{}

func (mysql Mysql) Infer(
	storePackageDir string,
	_ string,
	column introspect.Column,
) (types.GoType, error) {
	return infer(storePackageDir, column)
}

func (mysql Mysql) ModelTemplate() string {
	return modelTemplate
}

func (mysql Mysql) QueryTemplate() string {
	return queryTemplate
}

//go:embed model.go.tmpl
var modelTemplate string

//go:embed query.go.tmpl
var queryTemplate string

func NewTranslate() types.Translate {
	return Mysql{}
}
