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

type MoviesLanguage struct {
	MovieId    *int64  `db:"movie_id" json:"movie_id"`
	LanguageId *string `db:"language_id" json:"language_id"`
}

func (m *MoviesLanguage) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("MovieId: %v", *m.MovieId),
			fmt.Sprintf("LanguageId: %v", *m.LanguageId),
		},
		", ",
	)

	return fmt.Sprintf("MoviesLanguage{%s}", content)
}

func (m *MoviesLanguage) TableName() string {
	return "app.movies_languages"
}

func (m *MoviesLanguage) PrimaryKey() []string {
	return []string{
		"movie_id",
		"language_id",
	}
}

func (m *MoviesLanguage) InsertQuery() string {
	return moviesLanguageInsertSql
}

func (m *MoviesLanguage) UpdateAllQuery() string {
	return moviesLanguageUpdateAllSql
}

func (m *MoviesLanguage) UpdateByPkQuery() string {
	return moviesLanguageUpdateByPkSql
}

func (m *MoviesLanguage) CountQuery() string {
	return moviesLanguageModelCountSql
}

func (m *MoviesLanguage) FindAllQuery() string {
	return moviesLanguageFindAllSql
}

func (m *MoviesLanguage) FindFirstQuery() string {
	return moviesLanguageFindFirstSql
}

func (m *MoviesLanguage) FindByPkQuery() string {
	return moviesLanguageFindByPkSql
}

func (m *MoviesLanguage) DeleteByPkQuery() string {
	return moviesLanguageDeleteByPkSql
}

func (m *MoviesLanguage) DeleteAllQuery() string {
	return moviesLanguageDeleteAllSql
}

// language=mysql
var moviesLanguageAllFieldsWhere = `
WHERE TRUE
    AND (CAST(:movie_id AS BIGINT) IS NULL or movie_id = :movie_id)
    AND (CAST(:language_id AS TEXT) IS NULL or language_id = :language_id)
`

// language=mysql
var moviesLanguagePkFieldsWhere = `
WHERE movie_id = :movie_id
  AND language_id = :language_id
`

// language=mysql
var moviesLanguageInsertSql = `
INSERT INTO app.movies_languages(
  movie_id,
  language_id
)
VALUES (
  :movie_id,
  :language_id
)
RETURNING
  movie_id,
  language_id;
`

// language=mysql
var moviesLanguageUpdateByPkSql = `
UPDATE app.movies_languages
SET
  movie_id = :movie_id,
  language_id = :language_id
` + moviesLanguagePkFieldsWhere + `
RETURNING
  movie_id,
  language_id;
`

// language=mysql
var moviesLanguageUpdateAllSql = `
UPDATE app.movies_languages
SET
  movie_id = :movie_id,
  language_id = :language_id
` + moviesLanguageAllFieldsWhere + `
RETURNING
  movie_id,
  language_id;
`

// language=mysql
var moviesLanguageModelCountSql = `
SELECT count(*) as count
FROM app.movies_languages
` + moviesLanguageAllFieldsWhere + ";"

// language=mysql
var moviesLanguageFindAllSql = `
SELECT
  movie_id,
  language_id
FROM app.movies_languages
` + moviesLanguageAllFieldsWhere + ";"

// language=mysql
var moviesLanguageFindFirstSql = strings.TrimRight(moviesLanguageFindAllSql, ";") + `
LIMIT 1;`

// language=mysql
var moviesLanguageFindByPkSql = `
SELECT
  movie_id,
  language_id
FROM app.movies_languages
` + moviesLanguagePkFieldsWhere + `
LIMIT 1;`

// language=mysql
var moviesLanguageDeleteByPkSql = `
DELETE FROM app.movies_languages
` + moviesLanguagePkFieldsWhere + ";"

// language=postgresql
var moviesLanguageDeleteAllSql = `
DELETE FROM app.movies_languages
` + moviesLanguageAllFieldsWhere + ";"
