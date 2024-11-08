package pg

import (
	_ "embed"
	"encoding/json"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bradleyjkemp/cupaloy"
	"github.com/jmoiron/sqlx"
	"github.com/mvoorberg/sqlxgen/internal/utils"
	"github.com/mvoorberg/sqlxgen/internal/utils/array"
	"github.com/mvoorberg/sqlxgen/internal/utils/fs"
	"github.com/stretchr/testify/assert"
)

func TestMsgWithFilename(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		filename string
		msg      string
		want     string
	}{
		{
			name:     "empty filename",
			filename: "",
			msg:      "failed to parse query params",
			want:     ": failed to parse query params",
		},
		{
			name:     "valid filename 1",
			filename: "foo.sql",
			msg:      "failed to parse query params",
			want:     "foo.sql: failed to parse query params",
		},
		{
			name:     "valid filename 2",
			filename: "foo/bar.sql",
			msg:      "failed to parse query params",
			want:     "foo/bar.sql: failed to parse query params",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := msgWithFilename(testCase.filename, testCase.msg)

			assert.Equal(t, testCase.want, got)
		})
	}
}

func TestGenerateIntrospectQuery(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		query string
		want  string
		err   error
	}{
		{
			name: "query 1",
			//language=PostgreSQL
			query: `select * from users`,
			//language=PostgreSQL
			want: `--
drop table if exists sample_query_introspection;
--
create temp table if not exists sample_query_introspection as
select * from users;
--
select
attr.attname as column_name,
regexp_replace(tp.typname, '^_(\w+)$', '\1') as type,
tp.oid as type_id,
tp.typcategory = 'A' as is_array,
false as is_sequence,
not attr.attnotnull as nullable,
attr.attgenerated = 's' as generated
from pg_attribute attr
inner join pg_type tp on tp.oid = attr.atttypid
where true
and attr.attrelid = cast('sample_query_introspection' as regclass)
and attr.attnum > 0
and not attr.attisdropped
order by attr.attname;
`,
			err: nil,
		},
		{
			name: "query 2",
			//language=PostgreSQL
			query: `select * from users where id = :id; -- :id type: int`,
			//language=PostgreSQL
			want: `--
drop table if exists sample_query_introspection;
--
create temp table if not exists sample_query_introspection as
select * from users where false and (
id = :id
)
;;
--
select
attr.attname as column_name,
regexp_replace(tp.typname, '^_(\w+)$', '\1') as type,
tp.oid as type_id,
tp.typcategory = 'A' as is_array,
false as is_sequence,
not attr.attnotnull as nullable,
attr.attgenerated = 's' as generated
from pg_attribute attr
inner join pg_type tp on tp.oid = attr.atttypid
where true
and attr.attrelid = cast('sample_query_introspection' as regclass)
and attr.attnum > 0
and not attr.attisdropped
order by attr.attname;
`,
			err: nil,
		},
		{
			name: "custom 1",
			//language=PostgreSQL
			query: `
select
count(*) over () as "totalRecordsCount",
m.id as "id",
m.title as "title",
m.release_date as "releaseDate",
m.status as "status",
m.popularity as "popularity"
from movies m
where true
and (
  false
  or cast(:search as text) is null
  or m.title_search @@ to_tsquery(:search)
  or m.keywords_search @@ to_tsquery(:search)
) -- :search type: text
and (
  false
  or cast(:genre_id as text) is null
  or m.id in (
    select
    g.movie_id
    from movies_genres g
    where true
    and g.genre_id = :genre_id -- :genre_id type: text
    order by g.movie_id
  )
)
order by (case when :sort = 'desc' then m.id end) desc, m.id -- :sort type: text
limit :limit -- :limit type: int
offset :offset; -- :offset type: int`,
			//language=PostgreSQL
			want: `--
drop table if exists sample_query_introspection;
--
create temp table if not exists sample_query_introspection as
select
count(*) over () as "totalRecordsCount",
m.id as "id",
m.title as "title",
m.release_date as "releaseDate",
m.status as "status",
m.popularity as "popularity"
from movies m
where true
and (
  false
  or cast(:search as text) is null
  or m.title_search @@ to_tsquery(:search)
  or m.keywords_search @@ to_tsquery(:search)
) 
and (
  false
  or cast(:genre_id as text) is null
  or m.id in (
    select
    g.movie_id
    from movies_genres g
    where false and (
true
    and g.genre_id = :genre_id
)
order by g.movie_id
  )
)
order by (case when :sort = 'desc' then m.id end) desc, m.id 
limit :limit 
offset :offset;;
--
select
attr.attname as column_name,
regexp_replace(tp.typname, '^_(\w+)$', '\1') as type,
tp.oid as type_id,
tp.typcategory = 'A' as is_array,
false as is_sequence,
not attr.attnotnull as nullable,
attr.attgenerated = 's' as generated
from pg_attribute attr
inner join pg_type tp on tp.oid = attr.atttypid
where true
and attr.attrelid = cast('sample_query_introspection' as regclass)
and attr.attnum > 0
and not attr.attisdropped
order by attr.attname;
`,
			err: nil,
		},
		{
			name: "custom 2",
			//language=PostgreSQL
			query: `
select
m."id" as "id",
m."title" as "title",
m."original_title" as "originalTitle",
m."original_language" as "originalLanguage",
m."overview" as "overview",
m."runtime" as "runtime",
m."release_date" as "releaseDate",
m."tagline" as "tagline",
m."status" as "status",
m."homepage" as "homepage",
m."popularity" as "popularity",
m."vote_average" as "voteAverage",
m."vote_count" as "voteCount",
m."budget" as "budget",
m."revenue" as "revenue",
m."keywords" as "keywords",
coalesce(
  (
    select
    jsonb_agg(
      jsonb_build_object(
        'id', g.genre_id,
        'name', hp.friendly_name
      ) order by hp.friendly_name
    )
    from movies_genres g
    inner join hyper_parameter hp on (
      true
      and hp.type = 'genre'
      and hp.value = g.genre_id
    )
    where true
    and g.movie_id = m.id
  ),
  '[]'
) as "genres",
coalesce(
  (
    select
    jsonb_agg(
      jsonb_build_object(
        'id', c.country_id,
        'name', hp.friendly_name
      ) order by hp.friendly_name
    )
    from movies_countries c
    inner join hyper_parameter hp on (
      true
      and hp.type = 'country'
      and hp.value = c.country_id
    )
    where true
    and c.movie_id = m.id
  ),
  '[]'
) as "countries",
coalesce(
  (
    select
    jsonb_agg(
      jsonb_build_object(
        'id', l.language_id,
        'name', hp.friendly_name
      ) order by hp.friendly_name
    )
    from movies_languages l
    inner join hyper_parameter hp on (
      true
      and hp.type = 'language'
      and hp.value = l.language_id
    )
    where true
    and l.movie_id = m.id
  ),
  '[]'
) as "languages",
coalesce(
  (
    select
    jsonb_agg(
      jsonb_build_object(
        'id', mc.company_id,
        'name', c.name
      ) order by c.name
    )
    from movies_companies mc
    inner join companies c on mc.company_id = c.id
    where true
    and mc.movie_id = m.id
  ),
  '[]'
) as "companies",
coalesce(
  (
    select
    jsonb_agg(
      jsonb_build_object(
        'id', ma.actor_id,
        'name', a.name,
        'character', ma.character,
        'order', ma.cast_order
      ) order by ma.cast_order
    )
    from movies_actors ma
    inner join actors a on ma.actor_id = a.id
    where true
    and ma.movie_id = m.id
  ),
  '[]'
) as "actors",
coalesce(
  (
    select
    jsonb_agg(
      jsonb_build_object(
        'id', mc.crew_id,
        'name', c.name,
        'job', j.friendly_name,
        'department', d.friendly_name
      ) order by j.name
    )
    from movies_crew mc
    inner join crew c on mc.crew_id = c.id
    inner join hyper_parameter d on (
      true
      and d.type = 'department'
      and d.value = mc.department_id
    )
    inner join hyper_parameter j on (
      true
      and j.type = 'job'
      and j.value = mc.job_id
    )
    where true
    and mc.movie_id = m.id
  ),
  '[]'
) as "crews",
1
from movies m
where true
and m.id = :id; -- :id type: bigint
`,
			//language=PostgreSQL
			want: `--
drop table if exists sample_query_introspection;
--
create temp table if not exists sample_query_introspection as
select
m."id" as "id",
m."title" as "title",
m."original_title" as "originalTitle",
m."original_language" as "originalLanguage",
m."overview" as "overview",
m."runtime" as "runtime",
m."release_date" as "releaseDate",
m."tagline" as "tagline",
m."status" as "status",
m."homepage" as "homepage",
m."popularity" as "popularity",
m."vote_average" as "voteAverage",
m."vote_count" as "voteCount",
m."budget" as "budget",
m."revenue" as "revenue",
m."keywords" as "keywords",
coalesce(
  (
    select
    jsonb_agg(
      jsonb_build_object(
        'id', g.genre_id,
        'name', hp.friendly_name
      ) order by hp.friendly_name
    )
    from movies_genres g
    inner join hyper_parameter hp on (
      true
      and hp.type = 'genre'
      and hp.value = g.genre_id
    )
    where false and (
true
    and g.movie_id = m.id
)
),
  '[]'
) as "genres",
coalesce(
  (
    select
    jsonb_agg(
      jsonb_build_object(
        'id', c.country_id,
        'name', hp.friendly_name
      ) order by hp.friendly_name
    )
    from movies_countries c
    inner join hyper_parameter hp on (
      true
      and hp.type = 'country'
      and hp.value = c.country_id
    )
    where false and (
true
    and c.movie_id = m.id
)
),
  '[]'
) as "countries",
coalesce(
  (
    select
    jsonb_agg(
      jsonb_build_object(
        'id', l.language_id,
        'name', hp.friendly_name
      ) order by hp.friendly_name
    )
    from movies_languages l
    inner join hyper_parameter hp on (
      true
      and hp.type = 'language'
      and hp.value = l.language_id
    )
    where false and (
true
    and l.movie_id = m.id
)
),
  '[]'
) as "languages",
coalesce(
  (
    select
    jsonb_agg(
      jsonb_build_object(
        'id', mc.company_id,
        'name', c.name
      ) order by c.name
    )
    from movies_companies mc
    inner join companies c on mc.company_id = c.id
    where false and (
true
    and mc.movie_id = m.id
)
),
  '[]'
) as "companies",
coalesce(
  (
    select
    jsonb_agg(
      jsonb_build_object(
        'id', ma.actor_id,
        'name', a.name,
        'character', ma.character,
        'order', ma.cast_order
      ) order by ma.cast_order
    )
    from movies_actors ma
    inner join actors a on ma.actor_id = a.id
    where false and (
true
    and ma.movie_id = m.id
)
),
  '[]'
) as "actors",
coalesce(
  (
    select
    jsonb_agg(
      jsonb_build_object(
        'id', mc.crew_id,
        'name', c.name,
        'job', j.friendly_name,
        'department', d.friendly_name
      ) order by j.name
    )
    from movies_crew mc
    inner join crew c on mc.crew_id = c.id
    inner join hyper_parameter d on (
      true
      and d.type = 'department'
      and d.value = mc.department_id
    )
    inner join hyper_parameter j on (
      true
      and j.type = 'job'
      and j.value = mc.job_id
    )
    where false and (
true
    and mc.movie_id = m.id
)
),
  '[]'
) as "crews",
1
from movies m
where true
and m.id = :id;;
--
select
attr.attname as column_name,
regexp_replace(tp.typname, '^_(\w+)$', '\1') as type,
tp.oid as type_id,
tp.typcategory = 'A' as is_array,
false as is_sequence,
not attr.attnotnull as nullable,
attr.attgenerated = 's' as generated
from pg_attribute attr
inner join pg_type tp on tp.oid = attr.atttypid
where true
and attr.attrelid = cast('sample_query_introspection' as regclass)
and attr.attnum > 0
and not attr.attisdropped
order by attr.attname;
`,
			err: nil,
		},
		{
			name: "custom 3",
			//language=PostgreSQL
			query: `
select
c."id" as "id",
c."name" as "name",
coalesce(
  (
    select
    jsonb_agg(
      jsonb_build_object(
        'id', mc.movie_id,
        'title', m.title,
        'releaseDate', m.release_date,
        'job', j.friendly_name,
        'department', d.friendly_name
      ) order by m.release_date desc
    )
    from movies_crew mc
    inner join movies m on mc.movie_id = m.id
    inner join hyper_parameter d on (
      true
      and d.type = 'department'
      and d.value = mc.department_id
    )
    inner join hyper_parameter j on (
      true
      and j.type = 'job'
      and j.value = mc.job_id
    )
    where true
    and mc.crew_id = c.id
  ),
  '[]'
) as "movies"
from crew c
where c.id = :id; -- :id type: bigint`,
			//language=PostgreSQL
			want: `--
drop table if exists sample_query_introspection;
--
create temp table if not exists sample_query_introspection as
select
c."id" as "id",
c."name" as "name",
coalesce(
  (
    select
    jsonb_agg(
      jsonb_build_object(
        'id', mc.movie_id,
        'title', m.title,
        'releaseDate', m.release_date,
        'job', j.friendly_name,
        'department', d.friendly_name
      ) order by m.release_date desc
    )
    from movies_crew mc
    inner join movies m on mc.movie_id = m.id
    inner join hyper_parameter d on (
      true
      and d.type = 'department'
      and d.value = mc.department_id
    )
    inner join hyper_parameter j on (
      true
      and j.type = 'job'
      and j.value = mc.job_id
    )
    where false and (
true
    and mc.crew_id = c.id
)
),
  '[]'
) as "movies"
from crew c
where c.id = :id;;
--
select
attr.attname as column_name,
regexp_replace(tp.typname, '^_(\w+)$', '\1') as type,
tp.oid as type_id,
tp.typcategory = 'A' as is_array,
false as is_sequence,
not attr.attnotnull as nullable,
attr.attgenerated = 's' as generated
from pg_attribute attr
inner join pg_type tp on tp.oid = attr.atttypid
where true
and attr.attrelid = cast('sample_query_introspection' as regclass)
and attr.attnum > 0
and not attr.attisdropped
order by attr.attname;
`,
			err: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := generateIntrospectQuery(testCase.query)

			assert.Equal(t, testCase.err, err)

			assert.Equal(t, testCase.want, got)
		})
	}
}

func TestIntrospectQuery(t *testing.T) {
	testCases := []struct {
		name       string
		args       QueryArgs
		resultsCsv string
	}{
		{
			name: "list actors",
			args: QueryArgs{
				Query:    listActorsQuery,
				Filename: "list-actors.sql",
				GenDir:   "fixtures",
			},
			resultsCsv: listActorsResultCsv,
		},
		{
			name: "get actor",
			args: QueryArgs{
				Query:    getActorQuery,
				Filename: "get-actor.sql",
				GenDir:   "fixtures",
			},
			resultsCsv: getActorResultCsv,
		},
		{
			name: "list movies",
			args: QueryArgs{
				Query:    listMoviesQuery,
				Filename: "list-movies.sql",
				GenDir:   "fixtures",
			},
			resultsCsv: listMoviesResultCsv,
		},
		{
			name: "get movie",
			args: QueryArgs{
				Query:    getMovieQuery,
				Filename: "get-movie.sql",
				GenDir:   "fixtures",
			},
			resultsCsv: getMoviesResultCsv,
		},
	}

	db, mock, err := utils.NewMockSqlx()

	if err != nil {
		t.Fatalf("failed to create mock db: %v", err)
	}

	defer func(db *sqlx.DB) {
		err := db.Close()

		if err != nil {
			t.Fatalf("failed to close mock db: %v", err)
		}
	}(db)

	mock.ExpectBegin()

	for _, testCase := range testCases {
		mock.ExpectExec("drop table if exists sample_query_introspection").
			WillReturnResult(
				sqlmock.NewResult(0, 0),
			)

		mock.ExpectExec("create temp table if not exists sample_query_introspection (.+)").
			WillReturnResult(
				sqlmock.NewResult(0, 0),
			)

		mock.ExpectQuery("select (.+) from pg_attribute attr (.+)").
			WillReturnRows(
				sqlmock.NewRows([]string{"column_name", "type", "type_id", "is_array", "is_sequence", "nullable", "generated"}).
					FromCSVString(testCase.resultsCsv),
			)
	}

	mock.ExpectRollback()

	mock.ExpectClose()

	tx, err := db.Beginx()

	if err != nil {
		t.Fatalf("failed to create transaction: %v", err)
	}

	defer func(tx *sqlx.Tx) {
		err := tx.Rollback()

		if err != nil {
			t.Fatalf("failed to rollback transaction: %v", err)
		}
	}(tx)

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := introspectQuery(tx, testCase.args)

			assert.NoError(t, err)

			assert.Equal(t, testCase.args.Filename, got.Filename)

			gotJson, err := json.MarshalIndent(got, "", "  ")

			if err != nil {
				t.Fatalf("failed to marshal query: %v", err)
			}

			cupaloy.SnapshotT(t, gotJson)
		})
	}
}

func TestIntrospectQueries(t *testing.T) {
	type testCase struct {
		name     string
		query    string
		result   string
		filename string
	}

	testCases := []testCase{
		{
			name:     "list actors",
			filename: "list-actors.sql",
			query:    listActorsQuery,
			result:   listActorsResultCsv,
		},
		{
			name:     "get actor",
			filename: "get-actor.sql",
			query:    getActorQuery,
			result:   getActorResultCsv,
		},
		{
			name:     "list movies",
			filename: "list-movies.sql",
			query:    listMoviesQuery,
			result:   listMoviesResultCsv,
		},
		{
			name:     "get movie",
			filename: "get-movie.sql",
			query:    getMovieQuery,
			result:   getMoviesResultCsv,
		},
	}

	fd := fs.NewFakeFileDiscovery(
		array.Map(
			testCases,
			func(testCase testCase, i int) fs.FakeDiscover {
				return fs.FakeDiscover{
					Content:  testCase.query,
					Dir:      "fixtures",
					Filename: testCase.filename,
					FullPath: "fixtures/" + testCase.filename,
				}
			},
		),
	)

	s := NewIntrospect(fd, IntrospectArgs{QueryDirs: []string{"fixtures"}})

	db, mock, err := utils.NewMockSqlx()

	if err != nil {
		t.Fatalf("failed to create mock db: %v", err)
	}

	defer func(db *sqlx.DB) {
		err := db.Close()

		if err != nil {
			t.Fatalf("failed to close mock db: %v", err)
		}
	}(db)

	mock.ExpectBegin()

	for _, testCase := range testCases {
		mock.ExpectExec("drop table if exists sample_query_introspection").
			WillReturnResult(
				sqlmock.NewResult(0, 0),
			)

		mock.ExpectExec("create temp table if not exists sample_query_introspection (.+)").
			WillReturnResult(
				sqlmock.NewResult(0, 0),
			)

		mock.ExpectQuery("select (.+) from pg_attribute attr (.+)").
			WillReturnRows(
				sqlmock.NewRows([]string{"column_name", "type", "type_id", "is_array", "is_sequence", "nullable", "generated"}).
					FromCSVString(testCase.result),
			)
	}

	mock.ExpectRollback()

	mock.ExpectClose()

	tx, err := db.Beginx()

	if err != nil {
		t.Fatalf("failed to create transaction: %v", err)
	}

	defer func(tx *sqlx.Tx) {
		err := tx.Rollback()

		if err != nil {
			t.Fatalf("failed to rollback transaction: %v", err)
		}
	}(tx)

	queries, err := s.IntrospectQueries(tx)

	assert.NoError(t, err)

	for i, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := queries[i]

			assert.Equal(t, testCase.filename, got.Filename)

			gotJson, err := json.MarshalIndent(got, "", "  ")

			if err != nil {
				t.Fatalf("failed to marshal query: %v", err)
			}

			cupaloy.SnapshotT(t, gotJson)
		})
	}
}

//go:embed fixtures/list-actors.sql
var listActorsQuery string

//go:embed fixtures/list-actors.csv
var listActorsResultCsv string

//go:embed fixtures/get-actor.sql
var getActorQuery string

//go:embed fixtures/get-actor.csv
var getActorResultCsv string

//go:embed fixtures/list-movies.sql
var listMoviesQuery string

//go:embed fixtures/list-movies.csv
var listMoviesResultCsv string

//go:embed fixtures/get-movie.sql
var getMovieQuery string

//go:embed fixtures/get-movie.csv
var getMoviesResultCsv string
