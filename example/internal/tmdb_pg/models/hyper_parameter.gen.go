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

type HyperParameter struct {
	Type               *string `db:"type" json:"type"`
	Value              *string `db:"value" json:"value"`
	FriendlyName       *string `db:"friendly_name" json:"friendly_name"`
	FriendlyNameSearch *string `db:"friendly_name_search" json:"friendly_name_search"`
}

func (h *HyperParameter) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("Type: %v", *h.Type),
			fmt.Sprintf("Value: %v", *h.Value),
			fmt.Sprintf("FriendlyName: %v", *h.FriendlyName),
			fmt.Sprintf("FriendlyNameSearch: %v", *h.FriendlyNameSearch),
		},
		", ",
	)

	return fmt.Sprintf("HyperParameter{%s}", content)
}

func (h *HyperParameter) TableName() string {
	return "public.hyper_parameters"
}

func (h *HyperParameter) PrimaryKey() []string {
	return []string{
		"type",
		"value",
	}
}

func (h *HyperParameter) InsertQuery() string {
	return hyperParameterInsertSql
}

func (h *HyperParameter) CountQuery() string {
	return hyperParameterModelCountSql
}

func (h *HyperParameter) FindAllQuery() string {
	return hyperParameterFindAllSql
}

func (h *HyperParameter) FindFirstQuery() string {
	return hyperParameterFindFirstSql
}

func (h *HyperParameter) FindByPkQuery() string {
	return hyperParameterFindByPkSql
}

func (h *HyperParameter) DeleteByPkQuery() string {
	return hyperParameterDeleteByPkSql
}

func (h *HyperParameter) DeleteAllQuery() string {
	return hyperParameterDeleteAllSql
}

func (h *HyperParameter) GetPkWhere() string {
	return hyperParameterPkFieldsWhere
}

func (h *HyperParameter) GetAllFieldsWhere() string {
	return hyperParameterAllFieldsWhere
}

func (h *HyperParameter) GetReturning() string {
	return hyperParameterReturningFields
}

// language=postgresql
var hyperParameterAllFieldsWhere = `
WHERE TRUE
    AND (CAST(:type AS TEXT) IS NULL or type = :type)
    AND (CAST(:value AS TEXT) IS NULL or value = :value)
    AND (CAST(:friendly_name AS TEXT) IS NULL or friendly_name = :friendly_name)
    AND (CAST(:friendly_name_search AS TSVECTOR) IS NULL or friendly_name_search = :friendly_name_search)
`

// language=postgresql
var hyperParameterPkFieldsWhere = `
 WHERE type = :type
  AND value = :value
`

// language=postgresql
var hyperParameterReturningFields = `
 RETURNING type,
 value,
 friendly_name,
 friendly_name_search;
`

// language=postgresql
var hyperParameterInsertSql = `
INSERT INTO public.hyper_parameters(
  type,
  value,
  friendly_name
)
VALUES (
  :type,
  :value,
  :friendly_name
)` + hyperParameterReturningFields + ";"

// language=postgresql
var hyperParameterModelCountSql = `
SELECT count(*) as count
FROM public.hyper_parameters
` + hyperParameterAllFieldsWhere + ";"

// language=postgresql
var hyperParameterFindAllSql = `
SELECT
  type,
  value,
  friendly_name,
  friendly_name_search
FROM public.hyper_parameters
` + hyperParameterAllFieldsWhere + ";"

// language=postgresql
var hyperParameterFindFirstSql = strings.TrimRight(hyperParameterFindAllSql, ";") + `
LIMIT 1;`

// language=postgresql
var hyperParameterFindByPkSql = `
SELECT
  type,
  value,
  friendly_name,
  friendly_name_search
FROM public.hyper_parameters
` + hyperParameterPkFieldsWhere + `
LIMIT 1;`

// language=postgresql
var hyperParameterDeleteByPkSql = `
DELETE FROM public.hyper_parameters
` + hyperParameterPkFieldsWhere + ";"

// language=postgresql
var hyperParameterDeleteAllSql = `
DELETE FROM public.hyper_parameters
` + hyperParameterAllFieldsWhere + ";"
