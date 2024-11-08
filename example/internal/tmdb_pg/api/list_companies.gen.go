package api

import (
	_ "embed"
	"fmt"
	"github.com/mvoorberg/example/internal/tmdb_pg/store"
	"strings"
)

type ListCompaniesArgs struct {
	Limit  *int32  `db:"limit" json:"limit"`
	Offset *int32  `db:"offset" json:"offset"`
	Search *string `db:"search" json:"search"`
	Sort   *string `db:"sort" json:"sort"`
}

func (args *ListCompaniesArgs) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("Limit: %v", *args.Limit),
			fmt.Sprintf("Offset: %v", *args.Offset),
			fmt.Sprintf("Search: %v", *args.Search),
			fmt.Sprintf("Sort: %v", *args.Sort),
		},
		", ",
	)

	return fmt.Sprintf("ListCompaniesArgs{%s}", content)
}

func (args *ListCompaniesArgs) Query(db store.Database) ([]*ListCompaniesResult, error) {
	return store.Query[*ListCompaniesResult](db, args)
}

func (args *ListCompaniesArgs) Sql() string {
	return listCompaniesSql
}

type ListCompaniesResult struct {
	Id                *int64  `db:"id" json:"id"`
	Name              *string `db:"name" json:"name"`
	TotalRecordsCount *int64  `db:"totalRecordsCount" json:"totalRecordsCount"`
}

func (result *ListCompaniesResult) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("Id: %v", *result.Id),
			fmt.Sprintf("Name: %v", *result.Name),
			fmt.Sprintf("TotalRecordsCount: %v", *result.TotalRecordsCount),
		},
		", ",
	)

	return fmt.Sprintf("ListCompaniesResult{%s}", content)
}

//go:embed list-companies.sql
var listCompaniesSql string
