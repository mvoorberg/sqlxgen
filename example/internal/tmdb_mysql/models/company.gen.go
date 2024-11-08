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

type Company struct {
	Name *string `db:"name" json:"name"`
	Id   *int64  `db:"id" json:"id"`
}

func (c *Company) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("Name: %v", *c.Name),
			fmt.Sprintf("Id: %v", *c.Id),
		},
		", ",
	)

	return fmt.Sprintf("Company{%s}", content)
}

func (c *Company) TableName() string {
	return "app.companies"
}

func (c *Company) PrimaryKey() []string {
	return []string{
		"id",
	}
}

func (c *Company) InsertQuery() string {
	return companyInsertSql
}

func (c *Company) UpdateAllQuery() string {
	return companyUpdateAllSql
}

func (c *Company) UpdateByPkQuery() string {
	return companyUpdateByPkSql
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

// language=mysql
var companyAllFieldsWhere = `
WHERE TRUE
    AND (CAST(:name AS TEXT) IS NULL or name = :name)
    AND (CAST(:id AS BIGINT) IS NULL or id = :id)
`

// language=mysql
var companyPkFieldsWhere = `
WHERE id = :id
`

// language=mysql
var companyInsertSql = `
INSERT INTO app.companies(
  name,
  id
)
VALUES (
  :name,
  :id
)
RETURNING
  name,
  id;
`

// language=mysql
var companyUpdateByPkSql = `
UPDATE app.companies
SET
  name = :name,
  id = :id
` + companyPkFieldsWhere + `
RETURNING
  name,
  id;
`

// language=mysql
var companyUpdateAllSql = `
UPDATE app.companies
SET
  name = :name,
  id = :id
` + companyAllFieldsWhere + `
RETURNING
  name,
  id;
`

// language=mysql
var companyModelCountSql = `
SELECT count(*) as count
FROM app.companies
` + companyAllFieldsWhere + ";"

// language=mysql
var companyFindAllSql = `
SELECT
  name,
  id
FROM app.companies
` + companyAllFieldsWhere + ";"

// language=mysql
var companyFindFirstSql = strings.TrimRight(companyFindAllSql, ";") + `
LIMIT 1;`

// language=mysql
var companyFindByPkSql = `
SELECT
  name,
  id
FROM app.companies
` + companyPkFieldsWhere + `
LIMIT 1;`

// language=mysql
var companyDeleteByPkSql = `
DELETE FROM app.companies
` + companyPkFieldsWhere + ";"

// language=postgresql
var companyDeleteAllSql = `
DELETE FROM app.companies
` + companyAllFieldsWhere + ";"
