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

type MoviesCrew struct {
	DepartmentId *string `db:"department_id" json:"department_id"`
	JobId        *string `db:"job_id" json:"job_id"`
	MovieId      *int64  `db:"movie_id" json:"movie_id"`
	CrewId       *int64  `db:"crew_id" json:"crew_id"`
}

func (m *MoviesCrew) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("DepartmentId: %v", *m.DepartmentId),
			fmt.Sprintf("JobId: %v", *m.JobId),
			fmt.Sprintf("MovieId: %v", *m.MovieId),
			fmt.Sprintf("CrewId: %v", *m.CrewId),
		},
		", ",
	)

	return fmt.Sprintf("MoviesCrew{%s}", content)
}

func (m *MoviesCrew) TableName() string {
	return "app.movies_crew"
}

func (m *MoviesCrew) PrimaryKey() []string {
	return []string{
		"movie_id",
		"crew_id",
	}
}

func (m *MoviesCrew) InsertQuery() string {
	return moviesCrewInsertSql
}

func (m *MoviesCrew) UpdateAllQuery() string {
	return moviesCrewUpdateAllSql
}

func (m *MoviesCrew) UpdateByPkQuery() string {
	return moviesCrewUpdateByPkSql
}

func (m *MoviesCrew) CountQuery() string {
	return moviesCrewModelCountSql
}

func (m *MoviesCrew) FindAllQuery() string {
	return moviesCrewFindAllSql
}

func (m *MoviesCrew) FindFirstQuery() string {
	return moviesCrewFindFirstSql
}

func (m *MoviesCrew) FindByPkQuery() string {
	return moviesCrewFindByPkSql
}

func (m *MoviesCrew) DeleteByPkQuery() string {
	return moviesCrewDeleteByPkSql
}

func (m *MoviesCrew) DeleteAllQuery() string {
	return moviesCrewDeleteAllSql
}

// language=mysql
var moviesCrewAllFieldsWhere = `
WHERE TRUE
    AND (CAST(:department_id AS TEXT) IS NULL or department_id = :department_id)
    AND (CAST(:job_id AS TEXT) IS NULL or job_id = :job_id)
    AND (CAST(:movie_id AS BIGINT) IS NULL or movie_id = :movie_id)
    AND (CAST(:crew_id AS BIGINT) IS NULL or crew_id = :crew_id)
`

// language=mysql
var moviesCrewPkFieldsWhere = `
WHERE movie_id = :movie_id
  AND crew_id = :crew_id
`

// language=mysql
var moviesCrewInsertSql = `
INSERT INTO app.movies_crew(
  department_id,
  job_id,
  movie_id,
  crew_id
)
VALUES (
  :department_id,
  :job_id,
  :movie_id,
  :crew_id
)
RETURNING
  department_id,
  job_id,
  movie_id,
  crew_id;
`

// language=mysql
var moviesCrewUpdateByPkSql = `
UPDATE app.movies_crew
SET
  department_id = :department_id,
  job_id = :job_id,
  movie_id = :movie_id,
  crew_id = :crew_id
` + moviesCrewPkFieldsWhere + `
RETURNING
  department_id,
  job_id,
  movie_id,
  crew_id;
`

// language=mysql
var moviesCrewUpdateAllSql = `
UPDATE app.movies_crew
SET
  department_id = :department_id,
  job_id = :job_id,
  movie_id = :movie_id,
  crew_id = :crew_id
` + moviesCrewAllFieldsWhere + `
RETURNING
  department_id,
  job_id,
  movie_id,
  crew_id;
`

// language=mysql
var moviesCrewModelCountSql = `
SELECT count(*) as count
FROM app.movies_crew
` + moviesCrewAllFieldsWhere + ";"

// language=mysql
var moviesCrewFindAllSql = `
SELECT
  department_id,
  job_id,
  movie_id,
  crew_id
FROM app.movies_crew
` + moviesCrewAllFieldsWhere + ";"

// language=mysql
var moviesCrewFindFirstSql = strings.TrimRight(moviesCrewFindAllSql, ";") + `
LIMIT 1;`

// language=mysql
var moviesCrewFindByPkSql = `
SELECT
  department_id,
  job_id,
  movie_id,
  crew_id
FROM app.movies_crew
` + moviesCrewPkFieldsWhere + `
LIMIT 1;`

// language=mysql
var moviesCrewDeleteByPkSql = `
DELETE FROM app.movies_crew
` + moviesCrewPkFieldsWhere + ";"

// language=postgresql
var moviesCrewDeleteAllSql = `
DELETE FROM app.movies_crew
` + moviesCrewAllFieldsWhere + ";"
