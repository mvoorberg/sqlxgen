package models

// ************************************************************
// This is an example Postgres generated model.
// ************************************************************
// Options:
//   postgresInt64JsonString: true
//   createdDateFields: created_at
//   updatedDateFields: updated_at

import (
	"fmt"
	"strings"
)

type Company struct {
	Id         *int64  `db:"id" json:"id,string"`
	Name       *string `db:"name" json:"name"`
	NameSearch *string `db:"name_search" json:"name_search"`
}

func (c *Company) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("Id: %v", *c.Id),
			fmt.Sprintf("Name: %v", *c.Name),
			fmt.Sprintf("NameSearch: %v", *c.NameSearch),
		},
		", ",
	)

	return fmt.Sprintf("Company{%s}", content)
}

func (c *Company) TableName() string {
	return "public.companies"
}

func (c *Company) PrimaryKey() []string {
	return []string{
		"id",
	}
}

func (c *Company) InsertQuery() string {
	return companyInsertSql
}

func (c *Company) CountQuery() string {
	return companyModelCountSql
}

func (c *Company) FindAllQuery() string {
	return companyFindAllSql
}

func (c *Company) FindFirstQuery() string {
	return companyFindFirstSql
}

func (c *Company) FindByPkQuery() string {
	return companyFindByPkSql
}

func (c *Company) DeleteByPkQuery() string {
	return companyDeleteByPkSql
}

func (c *Company) DeleteAllQuery() string {
	return companyDeleteAllSql
}

func (c *Company) GetPkWhere() string {
	return companyPkFieldsWhere
}

func (c *Company) GetAllFieldsWhere() string {
	return companyAllFieldsWhere
}

func (c *Company) GetReturning() string {
	return companyReturningFields
}

// language=postgresql
var companyAllFieldsWhere = `
WHERE TRUE
    AND (CAST(:id AS INT8) IS NULL or id = :id)
    AND (CAST(:name AS TEXT) IS NULL or name = :name)
    AND (CAST(:name_search AS TSVECTOR) IS NULL or name_search = :name_search)
`

// language=postgresql
var companyPkFieldsWhere = `
 WHERE id = :id
`

// language=postgresql
var companyReturningFields = `
 RETURNING id,
 name,
 name_search;
`

// language=postgresql
var companyInsertSql = `
INSERT INTO public.companies(
  id,
  name
)
VALUES (
  :id,
  :name
)` + companyReturningFields + ";"

// language=postgresql
var companyModelCountSql = `
SELECT count(*) as count
FROM public.companies
` + companyAllFieldsWhere + ";"

// language=postgresql
var companyFindAllSql = `
SELECT
  id,
  name,
  name_search
FROM public.companies
` + companyAllFieldsWhere + ";"

// language=postgresql
var companyFindFirstSql = strings.TrimRight(companyFindAllSql, ";") + `
LIMIT 1;`

// language=postgresql
var companyFindByPkSql = `
SELECT
  id,
  name,
  name_search
FROM public.companies
` + companyPkFieldsWhere + `
LIMIT 1;`

// language=postgresql
var companyDeleteByPkSql = `
DELETE FROM public.companies
` + companyPkFieldsWhere + ";"

// language=postgresql
var companyDeleteAllSql = `
DELETE FROM public.companies
` + companyAllFieldsWhere + ";"
