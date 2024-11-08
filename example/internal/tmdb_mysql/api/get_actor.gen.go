package api

import (
	_ "embed"
	"fmt"
	"github.com/mvoorberg/example/internal/tmdb_mysql/store"
	"strings"
)

type GetActorArgs struct {
	Id *int64 `db:"id" json:"id"`
}

func (args *GetActorArgs) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("Id: %v", *args.Id),
		},
		", ",
	)

	return fmt.Sprintf("GetActorArgs{%s}", content)
}

func (args *GetActorArgs) Query(db store.Database) ([]*GetActorResult, error) {
	return store.Query[*GetActorResult](db, args)
}

func (args *GetActorArgs) Sql() string {
	return getActorSql
}

type GetActorResult struct {
}

func (result *GetActorResult) String() string {
	content := strings.Join(
		[]string{},
		", ",
	)

	return fmt.Sprintf("GetActorResult{%s}", content)
}

//go:embed get-actor.sql
var getActorSql string
