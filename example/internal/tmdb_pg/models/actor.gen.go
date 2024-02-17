package models

import (
	"fmt"
	"strings"
)

type Actor struct {
	Id         *int32  `db:"id" json:"id"`
	Name       *string `db:"name" json:"name"`
	NameSearch *string `db:"name_search" json:"name_search"`
}

func (a *Actor) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("Id: %v", *a.Id),
			fmt.Sprintf("Name: %v", *a.Name),
			fmt.Sprintf("NameSearch: %v", *a.NameSearch),
		},
		", ",
	)

	return fmt.Sprintf("Actor{%s}", content)
}

func (a *Actor) TableName() string {
	return "public.actors"
}

func (a *Actor) PrimaryKey() []string {
	return []string{
		"id",
	}
}

func (a *Actor) InsertQuery() string {
	return actorInsertSql
}

func (a *Actor) CountQuery() string {
	return actorModelCountSql
}

func (a *Actor) FindAllQuery() string {
	return actorFindAllSql
}

func (a *Actor) FindFirstQuery() string {
	return actorFindFirstSql
}

func (a *Actor) FindByPkQuery() string {
	return actorFindByPkSql
}

func (a *Actor) DeleteByPkQuery() string {
	return actorDeleteByPkSql
}

func (a *Actor) DeleteAllQuery() string {
	return actorDeleteAllSql
}

func (a *Actor) GetPkWhere() string {
	return actorPkFieldsWhere
}

func (a *Actor) GetAllFieldsWhere() string {
	return actorAllFieldsWhere
}

func (a *Actor) GetReturning() string {
	return actorReturningFields
}

// language=postgresql
var actorAllFieldsWhere = `
WHERE TRUE
    AND (CAST(:id AS INT4) IS NULL or id = :id)
    AND (CAST(:name AS TEXT) IS NULL or name = :name)
    AND (CAST(:name_search AS TSVECTOR) IS NULL or name_search = :name_search)
`

// language=postgresql
var actorPkFieldsWhere = `
 WHERE id = :id
`

// language=postgresql
var actorReturningFields = `
 RETURNING id,
 name,
 name_search;
`

// language=postgresql
var actorInsertSql = `
INSERT INTO public.actors(
  name
)
VALUES (
  :name
)` + actorReturningFields + ";"

// language=postgresql
var actorModelCountSql = `
SELECT count(*) as count
FROM public.actors
` + actorAllFieldsWhere + ";"

// language=postgresql
var actorFindAllSql = `
SELECT
  id,
  name,
  name_search
FROM public.actors
` + actorAllFieldsWhere + ";"

// language=postgresql
var actorFindFirstSql = strings.TrimRight(actorFindAllSql, ";") + `
LIMIT 1;`

// language=postgresql
var actorFindByPkSql = `
SELECT
  id,
  name,
  name_search
FROM public.actors
` + actorPkFieldsWhere + `
LIMIT 1;`

// language=postgresql
var actorDeleteByPkSql = `
DELETE FROM public.actors
` + actorPkFieldsWhere + ";"

// language=postgresql
var actorDeleteAllSql = `
DELETE FROM public.actors
` + actorAllFieldsWhere + ";"
