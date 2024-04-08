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

type MoviesActor struct {
	MovieId         *int64  `db:"movie_id" json:"movie_id,string"`
	ActorId         *int64  `db:"actor_id" json:"actor_id,string"`
	CastOrder       *int32  `db:"cast_order" json:"cast_order"`
	Character       *string `db:"character" json:"character"`
	CharacterSearch *string `db:"character_search" json:"character_search"`
}

func (m *MoviesActor) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("MovieId: %v", *m.MovieId),
			fmt.Sprintf("ActorId: %v", *m.ActorId),
			fmt.Sprintf("CastOrder: %v", *m.CastOrder),
			fmt.Sprintf("Character: %v", *m.Character),
			fmt.Sprintf("CharacterSearch: %v", *m.CharacterSearch),
		},
		", ",
	)

	return fmt.Sprintf("MoviesActor{%s}", content)
}

func (m *MoviesActor) TableName() string {
	return "public.movies_actors"
}

func (m *MoviesActor) PrimaryKey() []string {
	return []string{
		"movie_id",
		"actor_id",
	}
}

func (m *MoviesActor) InsertQuery() string {
	return moviesActorInsertSql
}

func (m *MoviesActor) CountQuery() string {
	return moviesActorModelCountSql
}

func (m *MoviesActor) FindAllQuery() string {
	return moviesActorFindAllSql
}

func (m *MoviesActor) FindFirstQuery() string {
	return moviesActorFindFirstSql
}

func (m *MoviesActor) FindByPkQuery() string {
	return moviesActorFindByPkSql
}

func (m *MoviesActor) DeleteByPkQuery() string {
	return moviesActorDeleteByPkSql
}

func (m *MoviesActor) DeleteAllQuery() string {
	return moviesActorDeleteAllSql
}

func (m *MoviesActor) GetPkWhere() string {
	return moviesActorPkFieldsWhere
}

func (m *MoviesActor) GetAllFieldsWhere() string {
	return moviesActorAllFieldsWhere
}

func (m *MoviesActor) GetReturning() string {
	return moviesActorReturningFields
}

// language=postgresql
var moviesActorAllFieldsWhere = `
WHERE TRUE
    AND (CAST(:movie_id AS INT8) IS NULL or movie_id = :movie_id)
    AND (CAST(:actor_id AS INT8) IS NULL or actor_id = :actor_id)
    AND (CAST(:cast_order AS INT4) IS NULL or cast_order = :cast_order)
    AND (CAST(:character AS TEXT) IS NULL or character = :character)
    AND (CAST(:character_search AS TSVECTOR) IS NULL or character_search = :character_search)
`

// language=postgresql
var moviesActorPkFieldsWhere = `
 WHERE movie_id = :movie_id
  AND actor_id = :actor_id
`

// language=postgresql
var moviesActorReturningFields = `
 RETURNING movie_id,
 actor_id,
 cast_order,
 character,
 character_search;
`

// language=postgresql
var moviesActorInsertSql = `
INSERT INTO public.movies_actors(
  movie_id,
  actor_id,
  cast_order,
  character
)
VALUES (
  :movie_id,
  :actor_id,
  :cast_order,
  :character
)` + moviesActorReturningFields + ";"

// language=postgresql
var moviesActorModelCountSql = `
SELECT count(*) as count
FROM public.movies_actors
` + moviesActorAllFieldsWhere + ";"

// language=postgresql
var moviesActorFindAllSql = `
SELECT
  movie_id,
  actor_id,
  cast_order,
  character,
  character_search
FROM public.movies_actors
` + moviesActorAllFieldsWhere + ";"

// language=postgresql
var moviesActorFindFirstSql = strings.TrimRight(moviesActorFindAllSql, ";") + `
LIMIT 1;`

// language=postgresql
var moviesActorFindByPkSql = `
SELECT
  movie_id,
  actor_id,
  cast_order,
  character,
  character_search
FROM public.movies_actors
` + moviesActorPkFieldsWhere + `
LIMIT 1;`

// language=postgresql
var moviesActorDeleteByPkSql = `
DELETE FROM public.movies_actors
` + moviesActorPkFieldsWhere + ";"

// language=postgresql
var moviesActorDeleteAllSql = `
DELETE FROM public.movies_actors
` + moviesActorAllFieldsWhere + ";"
