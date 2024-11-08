package api

import (
	_ "embed"
	"fmt"
	"github.com/mvoorberg/example/internal/tmdb_pg/store"
	"strings"
)

type ListActorsArgs struct {
	Limit  *int32  `db:"limit" json:"limit"`
	Offset *int32  `db:"offset" json:"offset"`
	Search *string `db:"search" json:"search"`
	Sort   *string `db:"sort" json:"sort"`
}

func (args *ListActorsArgs) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("Limit: %v", *args.Limit),
			fmt.Sprintf("Offset: %v", *args.Offset),
			fmt.Sprintf("Search: %v", *args.Search),
			fmt.Sprintf("Sort: %v", *args.Sort),
		},
		", ",
	)

	return fmt.Sprintf("ListActorsArgs{%s}", content)
}

func (args *ListActorsArgs) Query(db store.Database) ([]*ListActorsResult, error) {
	return store.Query[*ListActorsResult](db, args)
}

func (args *ListActorsArgs) Sql() string {
	return listActorsSql
}

type ListActorsResult struct {
	Id                *int32  `db:"id" json:"id"`
	Name              *string `db:"name" json:"name"`
	TotalRecordsCount *int64  `db:"totalRecordsCount" json:"totalRecordsCount"`
}

func (result *ListActorsResult) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("Id: %v", *result.Id),
			fmt.Sprintf("Name: %v", *result.Name),
			fmt.Sprintf("TotalRecordsCount: %v", *result.TotalRecordsCount),
		},
		", ",
	)

	return fmt.Sprintf("ListActorsResult{%s}", content)
}

//go:embed list-actors.sql
var listActorsSql string
