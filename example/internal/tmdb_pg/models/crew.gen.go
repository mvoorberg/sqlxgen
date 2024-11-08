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

type Crew struct {
	Id         *int32  `db:"id" json:"id"`
	Name       *string `db:"name" json:"name"`
	NameSearch *string `db:"name_search" json:"name_search"`
}

func (c *Crew) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("Id: %v", *c.Id),
			fmt.Sprintf("Name: %v", *c.Name),
			fmt.Sprintf("NameSearch: %v", *c.NameSearch),
		},
		", ",
	)

	return fmt.Sprintf("Crew{%s}", content)
}

func (c *Crew) TableName() string {
	return "public.crew"
}

func (c *Crew) PrimaryKey() []string {
	return []string{
		"id",
	}
}

func (c *Crew) InsertQuery() string {
	return crewInsertSql
}

func (c *Crew) CountQuery() string {
	return crewModelCountSql
}

func (c *Crew) FindAllQuery() string {
	return crewFindAllSql
}

func (c *Crew) FindFirstQuery() string {
	return crewFindFirstSql
}

func (c *Crew) FindByPkQuery() string {
	return crewFindByPkSql
}

func (c *Crew) DeleteByPkQuery() string {
	return crewDeleteByPkSql
}

func (c *Crew) DeleteAllQuery() string {
	return crewDeleteAllSql
}

func (c *Crew) GetPkWhere() string {
	return crewPkFieldsWhere
}

func (c *Crew) GetAllFieldsWhere() string {
	return crewAllFieldsWhere
}

func (c *Crew) GetReturning() string {
	return crewReturningFields
}

// language=postgresql
var crewAllFieldsWhere = `
WHERE TRUE
    AND (CAST(:id AS INT4) IS NULL or id = :id)
    AND (CAST(:name AS TEXT) IS NULL or name = :name)
    AND (CAST(:name_search AS TSVECTOR) IS NULL or name_search = :name_search)
`

// language=postgresql
var crewPkFieldsWhere = `
 WHERE id = :id
`

// language=postgresql
var crewReturningFields = `
 RETURNING id,
 name,
 name_search;
`

// language=postgresql
var crewInsertSql = `
INSERT INTO public.crew(
  name
)
VALUES (
  :name
)` + crewReturningFields + ";"

// language=postgresql
var crewModelCountSql = `
SELECT count(*) as count
FROM public.crew
` + crewAllFieldsWhere + ";"

// language=postgresql
var crewFindAllSql = `
SELECT
  id,
  name,
  name_search
FROM public.crew
` + crewAllFieldsWhere + ";"

// language=postgresql
var crewFindFirstSql = strings.TrimRight(crewFindAllSql, ";") + `
LIMIT 1;`

// language=postgresql
var crewFindByPkSql = `
SELECT
  id,
  name,
  name_search
FROM public.crew
` + crewPkFieldsWhere + `
LIMIT 1;`

// language=postgresql
var crewDeleteByPkSql = `
DELETE FROM public.crew
` + crewPkFieldsWhere + ";"

// language=postgresql
var crewDeleteAllSql = `
DELETE FROM public.crew
` + crewAllFieldsWhere + ";"
