package models

// ************************************************************
// This is an example MySql generated model.
// ************************************************************
// Options:
//   postgresInt64JsonString:

import (
	"fmt"
	"strings"
)

type Crew struct {
	Name *string `db:"name" json:"name"`
	Id   *int64  `db:"id" json:"id"`
}

func (c *Crew) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("Name: %v", *c.Name),
			fmt.Sprintf("Id: %v", *c.Id),
		},
		", ",
	)

	return fmt.Sprintf("Crew{%s}", content)
}

func (c *Crew) TableName() string {
	return "app.crew"
}

func (c *Crew) PrimaryKey() []string {
	return []string{
		"id",
	}
}

func (c *Crew) InsertQuery() string {
	return crewInsertSql
}

func (c *Crew) UpdateAllQuery() string {
	return crewUpdateAllSql
}

func (c *Crew) UpdateByPkQuery() string {
	return crewUpdateByPkSql
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

// language=mysql
var crewAllFieldsWhere = `
WHERE TRUE
    AND (CAST(:name AS TEXT) IS NULL or name = :name)
    AND (CAST(:id AS BIGINT) IS NULL or id = :id)
`

// language=mysql
var crewPkFieldsWhere = `
WHERE id = :id
`

// language=mysql
var crewInsertSql = `
INSERT INTO app.crew(
  name
)
VALUES (
  :name
)
RETURNING
  name,
  id;
`

// language=mysql
var crewUpdateByPkSql = `
UPDATE app.crew
SET
  name = :name,
  id = :id
` + crewPkFieldsWhere + `
RETURNING
  name,
  id;
`

// language=mysql
var crewUpdateAllSql = `
UPDATE app.crew
SET
  name = :name,
  id = :id
` + crewAllFieldsWhere + `
RETURNING
  name,
  id;
`

// language=mysql
var crewModelCountSql = `
SELECT count(*) as count
FROM app.crew
` + crewAllFieldsWhere + ";"

// language=mysql
var crewFindAllSql = `
SELECT
  name,
  id
FROM app.crew
` + crewAllFieldsWhere + ";"

// language=mysql
var crewFindFirstSql = strings.TrimRight(crewFindAllSql, ";") + `
LIMIT 1;`

// language=mysql
var crewFindByPkSql = `
SELECT
  name,
  id
FROM app.crew
` + crewPkFieldsWhere + `
LIMIT 1;`

// language=mysql
var crewDeleteByPkSql = `
DELETE FROM app.crew
` + crewPkFieldsWhere + ";"

// language=postgresql
var crewDeleteAllSql = `
DELETE FROM app.crew
` + crewAllFieldsWhere + ";"
