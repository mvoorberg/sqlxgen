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

type MoviesGenre struct {
	MovieId *int64  `db:"movie_id" json:"movie_id"`
	GenreId *string `db:"genre_id" json:"genre_id"`
}

func (m *MoviesGenre) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("MovieId: %v", *m.MovieId),
			fmt.Sprintf("GenreId: %v", *m.GenreId),
		},
		", ",
	)

	return fmt.Sprintf("MoviesGenre{%s}", content)
}

func (m *MoviesGenre) TableName() string {
	return "app.movies_genres"
}

func (m *MoviesGenre) PrimaryKey() []string {
	return []string{
		"movie_id",
		"genre_id",
	}
}

func (m *MoviesGenre) InsertQuery() string {
	return moviesGenreInsertSql
}

func (m *MoviesGenre) UpdateAllQuery() string {
	return moviesGenreUpdateAllSql
}

func (m *MoviesGenre) UpdateByPkQuery() string {
	return moviesGenreUpdateByPkSql
}

func (m *MoviesGenre) CountQuery() string {
	return moviesGenreModelCountSql
}

func (m *MoviesGenre) FindAllQuery() string {
	return moviesGenreFindAllSql
}

func (m *MoviesGenre) FindFirstQuery() string {
	return moviesGenreFindFirstSql
}

func (m *MoviesGenre) FindByPkQuery() string {
	return moviesGenreFindByPkSql
}

func (m *MoviesGenre) DeleteByPkQuery() string {
	return moviesGenreDeleteByPkSql
}

func (m *MoviesGenre) DeleteAllQuery() string {
	return moviesGenreDeleteAllSql
}

// language=mysql
var moviesGenreAllFieldsWhere = `
WHERE TRUE
    AND (CAST(:movie_id AS BIGINT) IS NULL or movie_id = :movie_id)
    AND (CAST(:genre_id AS TEXT) IS NULL or genre_id = :genre_id)
`

// language=mysql
var moviesGenrePkFieldsWhere = `
WHERE movie_id = :movie_id
  AND genre_id = :genre_id
`

// language=mysql
var moviesGenreInsertSql = `
INSERT INTO app.movies_genres(
  movie_id,
  genre_id
)
VALUES (
  :movie_id,
  :genre_id
)
RETURNING
  movie_id,
  genre_id;
`

// language=mysql
var moviesGenreUpdateByPkSql = `
UPDATE app.movies_genres
SET
  movie_id = :movie_id,
  genre_id = :genre_id
` + moviesGenrePkFieldsWhere + `
RETURNING
  movie_id,
  genre_id;
`

// language=mysql
var moviesGenreUpdateAllSql = `
UPDATE app.movies_genres
SET
  movie_id = :movie_id,
  genre_id = :genre_id
` + moviesGenreAllFieldsWhere + `
RETURNING
  movie_id,
  genre_id;
`

// language=mysql
var moviesGenreModelCountSql = `
SELECT count(*) as count
FROM app.movies_genres
` + moviesGenreAllFieldsWhere + ";"

// language=mysql
var moviesGenreFindAllSql = `
SELECT
  movie_id,
  genre_id
FROM app.movies_genres
` + moviesGenreAllFieldsWhere + ";"

// language=mysql
var moviesGenreFindFirstSql = strings.TrimRight(moviesGenreFindAllSql, ";") + `
LIMIT 1;`

// language=mysql
var moviesGenreFindByPkSql = `
SELECT
  movie_id,
  genre_id
FROM app.movies_genres
` + moviesGenrePkFieldsWhere + `
LIMIT 1;`

// language=mysql
var moviesGenreDeleteByPkSql = `
DELETE FROM app.movies_genres
` + moviesGenrePkFieldsWhere + ";"

// language=postgresql
var moviesGenreDeleteAllSql = `
DELETE FROM app.movies_genres
` + moviesGenreAllFieldsWhere + ";"
