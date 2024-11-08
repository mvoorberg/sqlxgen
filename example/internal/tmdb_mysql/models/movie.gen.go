package models

// ************************************************************
// This is an example MySql generated model.
// ************************************************************
// Options:
//   postgresInt64JsonString:

import (
	"fmt"
	"strings"
	"time"
)

type Movie struct {
	Budget           *int64     `db:"budget" json:"budget"`
	Homepage         *string    `db:"homepage" json:"homepage"`
	Keywords         *string    `db:"keywords" json:"keywords"`
	OriginalLanguage *string    `db:"original_language" json:"original_language"`
	OriginalTitle    *string    `db:"original_title" json:"original_title"`
	Overview         *string    `db:"overview" json:"overview"`
	Popularity       *float32   `db:"popularity" json:"popularity"`
	ReleaseDate      *time.Time `db:"release_date" json:"release_date"`
	Revenue          *int64     `db:"revenue" json:"revenue"`
	Runtime          *int32     `db:"runtime" json:"runtime"`
	Status           *string    `db:"status" json:"status"`
	Tagline          *string    `db:"tagline" json:"tagline"`
	Title            *string    `db:"title" json:"title"`
	VoteAverage      *float32   `db:"vote_average" json:"vote_average"`
	VoteCount        *int32     `db:"vote_count" json:"vote_count"`
	Id               *int64     `db:"id" json:"id"`
}

func (m *Movie) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("Budget: %v", *m.Budget),
			fmt.Sprintf("Homepage: %v", *m.Homepage),
			fmt.Sprintf("Keywords: %v", *m.Keywords),
			fmt.Sprintf("OriginalLanguage: %v", *m.OriginalLanguage),
			fmt.Sprintf("OriginalTitle: %v", *m.OriginalTitle),
			fmt.Sprintf("Overview: %v", *m.Overview),
			fmt.Sprintf("Popularity: %v", *m.Popularity),
			fmt.Sprintf("ReleaseDate: %v", *m.ReleaseDate),
			fmt.Sprintf("Revenue: %v", *m.Revenue),
			fmt.Sprintf("Runtime: %v", *m.Runtime),
			fmt.Sprintf("Status: %v", *m.Status),
			fmt.Sprintf("Tagline: %v", *m.Tagline),
			fmt.Sprintf("Title: %v", *m.Title),
			fmt.Sprintf("VoteAverage: %v", *m.VoteAverage),
			fmt.Sprintf("VoteCount: %v", *m.VoteCount),
			fmt.Sprintf("Id: %v", *m.Id),
		},
		", ",
	)

	return fmt.Sprintf("Movie{%s}", content)
}

func (m *Movie) TableName() string {
	return "app.movies"
}

func (m *Movie) PrimaryKey() []string {
	return []string{
		"id",
	}
}

func (m *Movie) InsertQuery() string {
	return movieInsertSql
}

func (m *Movie) UpdateAllQuery() string {
	return movieUpdateAllSql
}

func (m *Movie) UpdateByPkQuery() string {
	return movieUpdateByPkSql
}

func (m *Movie) CountQuery() string {
	return movieModelCountSql
}

func (m *Movie) FindAllQuery() string {
	return movieFindAllSql
}

func (m *Movie) FindFirstQuery() string {
	return movieFindFirstSql
}

func (m *Movie) FindByPkQuery() string {
	return movieFindByPkSql
}

func (m *Movie) DeleteByPkQuery() string {
	return movieDeleteByPkSql
}

func (m *Movie) DeleteAllQuery() string {
	return movieDeleteAllSql
}

// language=mysql
var movieAllFieldsWhere = `
WHERE TRUE
    AND (CAST(:budget AS BIGINT) IS NULL or budget = :budget)
    AND (CAST(:homepage AS TEXT) IS NULL or homepage = :homepage)
    AND (CAST(:keywords AS TEXT) IS NULL or keywords = :keywords)
    AND (CAST(:original_language AS TEXT) IS NULL or original_language = :original_language)
    AND (CAST(:original_title AS TEXT) IS NULL or original_title = :original_title)
    AND (CAST(:overview AS TEXT) IS NULL or overview = :overview)
    AND (CAST(:popularity AS FLOAT) IS NULL or popularity = :popularity)
    AND (CAST(:release_date AS DATE) IS NULL or release_date = :release_date)
    AND (CAST(:revenue AS BIGINT) IS NULL or revenue = :revenue)
    AND (CAST(:runtime AS INT) IS NULL or runtime = :runtime)
    AND (CAST(:status AS TEXT) IS NULL or status = :status)
    AND (CAST(:tagline AS TEXT) IS NULL or tagline = :tagline)
    AND (CAST(:title AS TEXT) IS NULL or title = :title)
    AND (CAST(:vote_average AS FLOAT) IS NULL or vote_average = :vote_average)
    AND (CAST(:vote_count AS INT) IS NULL or vote_count = :vote_count)
    AND (CAST(:id AS BIGINT) IS NULL or id = :id)
`

// language=mysql
var moviePkFieldsWhere = `
WHERE id = :id
`

// language=mysql
var movieInsertSql = `
INSERT INTO app.movies(
  budget,
  homepage,
  keywords,
  original_language,
  original_title,
  overview,
  popularity,
  release_date,
  revenue,
  runtime,
  status,
  tagline,
  title,
  vote_average,
  vote_count
)
VALUES (
  :budget,
  :homepage,
  :keywords,
  :original_language,
  :original_title,
  :overview,
  :popularity,
  :release_date,
  :revenue,
  :runtime,
  :status,
  :tagline,
  :title,
  :vote_average,
  :vote_count
)
RETURNING
  budget,
  homepage,
  keywords,
  original_language,
  original_title,
  overview,
  popularity,
  release_date,
  revenue,
  runtime,
  status,
  tagline,
  title,
  vote_average,
  vote_count,
  id;
`

// language=mysql
var movieUpdateByPkSql = `
UPDATE app.movies
SET
  budget = :budget,
  homepage = :homepage,
  keywords = :keywords,
  original_language = :original_language,
  original_title = :original_title,
  overview = :overview,
  popularity = :popularity,
  release_date = :release_date,
  revenue = :revenue,
  runtime = :runtime,
  status = :status,
  tagline = :tagline,
  title = :title,
  vote_average = :vote_average,
  vote_count = :vote_count,
  id = :id
` + moviePkFieldsWhere + `
RETURNING
  budget,
  homepage,
  keywords,
  original_language,
  original_title,
  overview,
  popularity,
  release_date,
  revenue,
  runtime,
  status,
  tagline,
  title,
  vote_average,
  vote_count,
  id;
`

// language=mysql
var movieUpdateAllSql = `
UPDATE app.movies
SET
  budget = :budget,
  homepage = :homepage,
  keywords = :keywords,
  original_language = :original_language,
  original_title = :original_title,
  overview = :overview,
  popularity = :popularity,
  release_date = :release_date,
  revenue = :revenue,
  runtime = :runtime,
  status = :status,
  tagline = :tagline,
  title = :title,
  vote_average = :vote_average,
  vote_count = :vote_count,
  id = :id
` + movieAllFieldsWhere + `
RETURNING
  budget,
  homepage,
  keywords,
  original_language,
  original_title,
  overview,
  popularity,
  release_date,
  revenue,
  runtime,
  status,
  tagline,
  title,
  vote_average,
  vote_count,
  id;
`

// language=mysql
var movieModelCountSql = `
SELECT count(*) as count
FROM app.movies
` + movieAllFieldsWhere + ";"

// language=mysql
var movieFindAllSql = `
SELECT
  budget,
  homepage,
  keywords,
  original_language,
  original_title,
  overview,
  popularity,
  release_date,
  revenue,
  runtime,
  status,
  tagline,
  title,
  vote_average,
  vote_count,
  id
FROM app.movies
` + movieAllFieldsWhere + ";"

// language=mysql
var movieFindFirstSql = strings.TrimRight(movieFindAllSql, ";") + `
LIMIT 1;`

// language=mysql
var movieFindByPkSql = `
SELECT
  budget,
  homepage,
  keywords,
  original_language,
  original_title,
  overview,
  popularity,
  release_date,
  revenue,
  runtime,
  status,
  tagline,
  title,
  vote_average,
  vote_count,
  id
FROM app.movies
` + moviePkFieldsWhere + `
LIMIT 1;`

// language=mysql
var movieDeleteByPkSql = `
DELETE FROM app.movies
` + moviePkFieldsWhere + ";"

// language=postgresql
var movieDeleteAllSql = `
DELETE FROM app.movies
` + movieAllFieldsWhere + ";"
