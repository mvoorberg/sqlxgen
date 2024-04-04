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

type MoviesCountry struct {
	MovieId   *int64  `db:"movie_id" json:"movie_id"`
	CountryId *string `db:"country_id" json:"country_id"`
}

func (m *MoviesCountry) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("MovieId: %v", *m.MovieId),
			fmt.Sprintf("CountryId: %v", *m.CountryId),
		},
		", ",
	)

	return fmt.Sprintf("MoviesCountry{%s}", content)
}

func (m *MoviesCountry) TableName() string {
	return "app.movies_countries"
}

func (m *MoviesCountry) PrimaryKey() []string {
	return []string{
		"movie_id",
		"country_id",
	}
}

func (m *MoviesCountry) InsertQuery() string {
	return moviesCountryInsertSql
}

func (m *MoviesCountry) UpdateAllQuery() string {
	return moviesCountryUpdateAllSql
}

func (m *MoviesCountry) UpdateByPkQuery() string {
	return moviesCountryUpdateByPkSql
}

func (m *MoviesCountry) CountQuery() string {
	return moviesCountryModelCountSql
}

func (m *MoviesCountry) FindAllQuery() string {
	return moviesCountryFindAllSql
}

func (m *MoviesCountry) FindFirstQuery() string {
	return moviesCountryFindFirstSql
}

func (m *MoviesCountry) FindByPkQuery() string {
	return moviesCountryFindByPkSql
}

func (m *MoviesCountry) DeleteByPkQuery() string {
	return moviesCountryDeleteByPkSql
}

func (m *MoviesCountry) DeleteAllQuery() string {
	return moviesCountryDeleteAllSql
}

// language=mysql
var moviesCountryAllFieldsWhere = `
WHERE TRUE
    AND (CAST(:movie_id AS BIGINT) IS NULL or movie_id = :movie_id)
    AND (CAST(:country_id AS TEXT) IS NULL or country_id = :country_id)
`

// language=mysql
var moviesCountryPkFieldsWhere = `
WHERE movie_id = :movie_id
  AND country_id = :country_id
`

// language=mysql
var moviesCountryInsertSql = `
INSERT INTO app.movies_countries(
  movie_id,
  country_id
)
VALUES (
  :movie_id,
  :country_id
)
RETURNING
  movie_id,
  country_id;
`

// language=mysql
var moviesCountryUpdateByPkSql = `
UPDATE app.movies_countries
SET
  movie_id = :movie_id,
  country_id = :country_id
` + moviesCountryPkFieldsWhere + `
RETURNING
  movie_id,
  country_id;
`

// language=mysql
var moviesCountryUpdateAllSql = `
UPDATE app.movies_countries
SET
  movie_id = :movie_id,
  country_id = :country_id
` + moviesCountryAllFieldsWhere + `
RETURNING
  movie_id,
  country_id;
`

// language=mysql
var moviesCountryModelCountSql = `
SELECT count(*) as count
FROM app.movies_countries
` + moviesCountryAllFieldsWhere + ";"

// language=mysql
var moviesCountryFindAllSql = `
SELECT
  movie_id,
  country_id
FROM app.movies_countries
` + moviesCountryAllFieldsWhere + ";"

// language=mysql
var moviesCountryFindFirstSql = strings.TrimRight(moviesCountryFindAllSql, ";") + `
LIMIT 1;`

// language=mysql
var moviesCountryFindByPkSql = `
SELECT
  movie_id,
  country_id
FROM app.movies_countries
` + moviesCountryPkFieldsWhere + `
LIMIT 1;`

// language=mysql
var moviesCountryDeleteByPkSql = `
DELETE FROM app.movies_countries
` + moviesCountryPkFieldsWhere + ";"

// language=postgresql
var moviesCountryDeleteAllSql = `
DELETE FROM app.movies_countries
` + moviesCountryAllFieldsWhere + ";"
