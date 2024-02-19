package tmdb_pg

import (
	"log"
	"testing"

	"github.com/jmoiron/sqlx"
	models "github.com/mvoorberg/example/internal/tmdb_pg/models"
	store "github.com/mvoorberg/example/internal/tmdb_pg/store"
	"gotest.tools/assert"

	// "models"
	// "store"

	_ "github.com/lib/pq"
)

type StoreTest struct {
	Db *sqlx.DB
}

// You can use testing.T, if you want to test the code without benchmarking
func setupSuite(t testing.T) (func(t testing.T), StoreTest) {
	log.Println("setup suite")

	st := StoreTest{}

	engine := "postgres"
	connectionUrl := "postgres://app:app@localhost:54320/app?sslmode=disable"

	var err error
	st.Db, err = sqlx.Open(engine, connectionUrl)
	if err != nil {
		t.Errorf("unable to connect to database: %v", err)
	}

	// Return a function to teardown the test
	return func(t testing.T) {
		log.Println("teardown suite")

		st.Db.Close()
	}, st
}

func TestMoviesCount(t *testing.T) {

	teardownSuite, st := setupSuite(*t)
	defer teardownSuite(*t)

	db := st.Db

	var id int32 = 24
	m := models.Movie{
		Id: &id,
	}

	// Count
	countMovies, err := store.Count[*models.Movie](db, &m)
	if err != nil {
		t.Errorf("unable : %v", err)
	}
	assert.Equal(t, countMovies, 1)

	// CountPtr
	countMoviesPtr, err := store.CountPtr[*models.Movie](db, &m)
	if err != nil {
		t.Errorf("unable : %v", err)
	}
	assert.Equal(t, *countMoviesPtr, int64(1))

	// CountSql
	countSql := `SELECT COUNT(*) as hello FROM public.movies
					WHERE original_language = :original_language`

	oLang := "zh"
	zhLang := struct {
		OriginalLanguage *string `db:"original_language"`
	}{
		OriginalLanguage: &oLang,
	}
	countMoviesSql, err := store.CountSql(db, countSql, zhLang)
	if err != nil {
		t.Errorf("unable : %v", err)
	}
	assert.Equal(t, countMoviesSql, 27)

}

func TestMoviesFind(t *testing.T) {

	teardownSuite, st := setupSuite(*t)
	defer teardownSuite(*t)

	db := st.Db

	var id int32 = 24
	m := models.Movie{
		Id: &id,
	}

	// Find by PK
	killBillVol1, err := store.FindByPk[*models.Movie](db, &m)
	if err != nil {
		t.Errorf("unable : %v", err)
	}
	assert.Equal(t, "Kill Bill: Vol. 1", *killBillVol1.Title)
	assert.Equal(t, id, *killBillVol1.Id)

	// Find One
	var title string = "Kill Bill: Vol. 1"
	m2 := models.Movie{
		Title: &title,
	}
	killBillVol1, err = store.FindOne[*models.Movie](db, &m2)
	if err != nil {
		t.Errorf("unable : %v", err)
	}
	assert.Equal(t, "Kill Bill: Vol. 1", *killBillVol1.Title)
	assert.Equal(t, id, *killBillVol1.Id)

	// Find First
	killBillVol1, err = store.FindFirst[*models.Movie](db, &m2)
	if err != nil {
		t.Errorf("unable : %v", err)
	}
	assert.Equal(t, "Kill Bill: Vol. 1", *killBillVol1.Title)
	assert.Equal(t, id, *killBillVol1.Id)

	// Find Many
	var lang string = "zh"
	m3 := models.Movie{
		OriginalLanguage: &lang,
	}
	manyMovies, err := store.FindMany[*models.Movie](db, &m3)
	if err != nil {
		t.Errorf("unable : %v", err)
	}
	assert.Equal(t, 27, len(manyMovies))
	assert.Equal(t, "zh", *manyMovies[0].OriginalLanguage)

	println("store find done")
}

func TestMoviesUpdate(t *testing.T) {

	teardownSuite, st := setupSuite(*t)
	defer teardownSuite(*t)

	db := st.Db

	var id int32 = 24
	var title string = "Kill Bill: Vol. 1"

	kbLang := "en"
	kb := models.Movie{}
	kb.Id = &id
	kb.Title = &title
	kb.OriginalLanguage = &kbLang

	// Update will only update the fields that are set!
	result, err := store.UpdateByPk[*models.Movie](db, &kb)
	if err != nil {
		t.Errorf("unable : %v", err)
	}
	assert.Equal(t, 24, int(*result.Id))

	badAkCols := []string{"dummy", "original_language"}
	_, err = store.UpdateByAk[*models.Movie](db, &kb, badAkCols)
	assert.ErrorContains(t, err, "alternate key dummy not found in Movie")

	// Update will only update the fields that are set!
	akCols := []string{"id", "original_language"}
	akResult, err := store.UpdateByAk[*models.Movie](db, &kb, akCols)
	if err != nil {
		t.Errorf("unable : %v", err)
	}
	assert.Equal(t, 24, int(*akResult.Id))

	// Update by AK will only update one record!
	akCols2 := []string{"original_language"}
	_, err = store.UpdateByAk[*models.Movie](db, &kb, akCols2)
	assert.ErrorContains(t, err, "update-by-AK Movie would have matched")

	// Update Many
	someMovies := []*models.Movie{}
	someMovies = append(someMovies, &kb)
	someMovies = append(someMovies, &kb)

	someResults, err := store.Update[*models.Movie](db, someMovies...)
	if err != nil {
		t.Errorf("unable : %v", err)
	}
	assert.Equal(t, 24, int(*someResults[0].Id))
	// })

}
