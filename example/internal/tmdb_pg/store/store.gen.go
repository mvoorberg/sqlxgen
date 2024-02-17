package store

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
)

// orm

// Get the type name of a model.
func GetTypeName[T any](instance T) string {
	t := reflect.TypeOf(instance)
	typeName := t.Name()
	if t.Kind() == reflect.Pointer {
		typeName = t.Elem().Name()
	}
	return typeName
}

// Insert a single record and reselect it.
func InsertOne[T model[P], P any](db Database, instance T) (T, error) {

	inserted, err := Insert[T](db, instance)
	if err != nil {
		return nil, err
	}

	return inserted[0], nil
}

// Insert a slice of records, one at a time. Return the inserted records.
func Insert[T model[P], P any](db Database, instances ...T) ([]T, error) {
	inserts := make([]T, 0)

	for _, instance := range instances {
		insertSql := instance.InsertQuery()
		rows, err := db.NamedQuery(insertSql, instance)

		if err != nil {
			return nil, err
		}

		hasNext := rows.Next()

		if !hasNext {
			return nil, fmt.Errorf("unable to insert %s", GetTypeName(instance))
		}

		inserted := new(P)

		err = rows.StructScan(inserted)
		if err != nil {
			return nil, err
		}

		rows.Close()

		inserts = append(inserts, inserted)
	}

	return inserts, nil
}

// Update a single record by Primary Key.
func UpdateByPk[T model[P], P any](db Database, instance T) (T, error) {

	pkCols := instance.PrimaryKey()
	if len(pkCols) == 0 {
		return nil, fmt.Errorf("primary key not defined for %s", GetTypeName(instance))
	}

	updateSql, err := getUpdateSql(instance, pkCols)
	if err != nil {
		return nil, err
	}

	*updateSql += instance.GetPkWhere()
	*updateSql += instance.GetReturning()

	return updateSingle[T, P](db, *updateSql, instance)
}

// Update a single record by an Alternate Key.
func UpdateByAk[T model[P], P any](db Database, instance T, altKeys []string) (T, error) {

	updateSql, err := getUpdateSql(instance, altKeys)
	if err != nil {
		return nil, err
	}

	altWhereSql, err := getAltKeyWhere(instance, altKeys)
	if err != nil {
		return nil, err
	}
	*updateSql += *altWhereSql
	*updateSql += instance.GetReturning()

	countSql := `SELECT COUNT(*) FROM ` + instance.TableName() + *altWhereSql
	result, err := CountSql(db, countSql, instance)
	if err != nil {
		return nil, err
	}
	if result != 1 {
		return nil, fmt.Errorf("update-by-AK %s would have matched %d rows", GetTypeName(instance), result)
	}

	return updateSingle[T, P](db, *updateSql, instance)
}

func getAltKeyWhere[T model[P], P any](model T, altKeys []string) (*string, error) {

	fields := getDbFieldMeta(model)

	where := " WHERE "
	for i, k := range altKeys {
		_, ok := fields[k]
		if !ok {
			return nil, fmt.Errorf("alternate key %s not found in %s", k, GetTypeName(model))
		}
		if i > 0 {
			where += " AND "
		}
		where += fmt.Sprintf("%s = :%s", k, k)
	}
	return &where, nil
}

func getUpdateSql[T model[P], P any](instance T, keyCols []string) (*string, error) {

	if len(keyCols) == 0 {
		return nil, fmt.Errorf("key columns not defined for %s", GetTypeName(instance))
	}

	tableName := instance.TableName()
	meta := getFieldMetaForUpdate(instance)

	updateSql := fmt.Sprintf("UPDATE %s SET ", tableName)
	setCols := 0
	for _, v := range meta {

		isKey := false // Don't update the Primary/Alternate Key cols!
		for _, k := range keyCols {
			if v.DbName == k {
				isKey = true
				break
			}
		}
		if !isKey && v.FieldHasValue {
			if setCols != 0 {
				updateSql += ","
			}
			updateSql += fmt.Sprintf("\n  %s = :%s", v.DbName, v.DbName)
			setCols++
		}
	}
	if setCols == 0 {
		return nil, fmt.Errorf("no fields to update on %s", GetTypeName(instance))
	}
	return &updateSql, nil
}

func updateSingle[T model[P], P any](db Database, updateSql string, instance T) (T, error) {

	rows, err := db.NamedQuery(updateSql, instance)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	hasNext := rows.Next()
	if !hasNext {
		return nil, fmt.Errorf("unable to update %s", GetTypeName(instance))
	}

	updated := new(P)
	err = rows.StructScan(updated)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

// Update a slice of records, one at a time. Return the updated records.
func Update[T model[P], P any](db Database, instances ...T) ([]T, error) {
	updates := make([]T, 0)

	// TODO: put this in a transaction and fail them all together
	for _, instance := range instances {
		updated, err := UpdateByPk[T](db, instance)
		if err != nil {
			return nil, err
		}
		updates = append(updates, updated)
	}

	return updates, nil
}

// Count the number of records that match the instance. Return count as a pointer.
func CountPtr[T model[P], P any](db Database, instance T) (*int64, error) {
	countSql := instance.CountQuery()
	return count(db, countSql, instance)
}

// Count the number of records that match the instance.
func Count[T model[P], P any](db Database, instance T) (int, error) {
	result, err := CountPtr[T](db, instance)
	if err != nil {
		return -1, err
	}
	return int(*result), nil
}

// Count records with a sql query. Bind struct fields to the query.
func CountSql(db Database, countSql string, args interface{}) (int, error) {
	result, err := count(db, countSql, args)
	if err != nil {
		return -1, err
	}
	return int(*result), nil
}

func count(db Database, countSql string, instance interface{}) (*int64, error) {
	result := new(int64)

	rows, err := db.NamedQuery(countSql, instance)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	hasNext := rows.Next()
	if !hasNext {
		return nil, fmt.Errorf("count %s failed", GetTypeName(instance))
	}

	err = rows.Scan(result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type QueryOptions struct {
	SelectList *[]string
	Paginator  *Paginator
	OrderBy    *string
}

type Paginator struct {
	Page     int
	PageSize int
}

func FindMany[T model[P], P any](db Database, instance T) ([]T, error) {
	return FindPage[T](db, instance, nil)
}

func FindPage[T model[P], P any](db Database, instance T, queryOpts *QueryOptions) ([]T, error) {

	findAllSql := instance.FindAllQuery()
	if queryOpts != nil {
		if queryOpts.SelectList != nil && len(*queryOpts.SelectList) > 0 {
			selectList := *queryOpts.SelectList
			fromSql := fmt.Sprintf("FROM %s", instance.TableName())
			whereSql := strings.Split(findAllSql, fromSql)[1]
			findAllSql = fmt.Sprintf("SELECT %s %s %s", strings.Join(selectList, ", "), fromSql, whereSql)
		}

		findAllSql = strings.TrimRight(findAllSql, "\r\n;")
		if queryOpts.OrderBy != nil {
			findAllSql += fmt.Sprintf(" ORDER BY %s", *queryOpts.OrderBy)
		} else {
			findAllSql += " ORDER BY 1"
		}

		pager := *queryOpts.Paginator
		if queryOpts.Paginator != nil && pager.PageSize > 0 && pager.Page > 0 {
			findAllSql += fmt.Sprintf(" LIMIT %d OFFSET %d", pager.PageSize, (pager.Page-1)*pager.PageSize)
		}
	}
	return findMany[T](db, instance, findAllSql, false)
}

func FindManySql[T model[P], P any](db Database, querySQL string, args interface{}) ([]T, error) {
	return findMany[T](db, args, querySQL, false)
}

func findMany[T model[P], P any](db Database, instance interface{}, sqlQuery string, failOnMulti bool) ([]T, error) {
	if instance == nil {
		instance = struct{}{}
	}
	rows, err := db.NamedQuery(sqlQuery, instance)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	idx := 0
	result := make([]T, 0)

	for rows.Next() {
		if failOnMulti && idx > 0 {
			return nil, ErrFoundMultiple
		}
		rowInstance := new(P)

		err = rows.StructScan(rowInstance)

		if err != nil {
			return nil, err
		}

		idx++
		result = append(result, rowInstance)
	}

	return result, nil
}

// Find limit 1
func FindFirst[T model[P], P any](db Database, instance T) (T, error) {
	return findSingle[T](db, instance, instance.FindFirstQuery())
}

// Find and return 1, err if > 1
func FindOne[T model[P], P any](db Database, instance T) (T, error) {
	querySql := instance.FindAllQuery()

	result, err := findMany[T](db, instance, querySql, true)
	if err != nil {
		return nil, err
	}
	return result[0], nil
}

func FindByPk[T model[P], P any](db Database, instance T) (T, error) {
	return findSingle[T](db, instance, instance.FindByPkQuery())
}

func FindOneSql[T model[P], P any](db Database, querySQL string, args interface{}) (T, error) {
	result, err := findMany[T](db, args, querySQL, true)
	if err != nil {
		return nil, err
	}
	return result[0], nil
}

func FindFirstSql[T model[P], P any](db Database, querySQL string, args interface{}) (T, error) {
	return findSingle[T](db, args, querySQL)
}

func findSingle[T model[P], P any](db Database, instance interface{}, sqlQuery string) (T, error) {
	if instance == nil {
		instance = struct{}{}
	}

	result := new(P)

	rows, err := db.NamedQuery(sqlQuery, instance)
	if err != nil {
		return result, err
	}

	defer rows.Close()

	hasNext := rows.Next()

	if !hasNext {
		msg := fmt.Sprintf("%s not found", GetTypeName(instance))

		return result, fmt.Errorf(msg)
	}

	err = rows.StructScan(result)

	if err != nil {
		return result, err
	}

	return result, nil
}

// Delete by Pk, err if not found
func DeleteByPk[T model[P], P any](db Database, instance T) error {

	result, err := db.NamedExec(instance.DeleteByPkQuery(), instance)
	if err != nil {
		return err
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAff != 1 {
		return fmt.Errorf("unable to delete %s", instance)
	}
	return nil
}

func DeleteOne[T model[P], P any](db Database, instance T) error {
	count, err := Count[T](db, instance)
	if err != nil {
		return err
	}
	if count != 1 {
		return fmt.Errorf("delete-one %s would have matched %d rows", GetTypeName(instance), count)
	}

	_, err = DeleteAll[T](db, instance)

	return err
}

func DeleteAll[T model[P], P any](db Database, instance T) (*int64, error) {

	result, err := db.NamedExec(instance.DeleteAllQuery(), instance)
	if err != nil {
		return nil, err
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	return &rowsAff, nil
}

type model[P any] interface {
	*P

	TableName() string
	PrimaryKey() []string

	InsertQuery() string

	CountQuery() string
	FindFirstQuery() string
	FindByPkQuery() string
	FindAllQuery() string

	DeleteByPkQuery() string
	DeleteAllQuery() string

	GetReturning() string
	GetPkWhere() string
	GetAllFieldsWhere() string
}

func Query[R result[pR], Q queryable[pQ], pR, pQ any](db Database, args Q) ([]R, error) {
	re, err := regexp.Compile(`-{2,}\s*([\w\W\s\S]*?)(\n|\z)`)

	if err != nil {
		return nil, err
	}

	query := re.ReplaceAllString(args.Sql(), "$2")
	rows, err := db.NamedQuery(query, args)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	result := make([]R, 0)

	for rows.Next() {
		instance := new(pR)
		err = rows.StructScan(instance)

		if err != nil {
			return nil, err
		}

		result = append(result, instance)
	}
	return result, nil
}

type queryable[P any] interface {
	*P

	Sql() string
}

type result[P any] interface {
	*P
}

// supplementary types

type Database interface {
	NamedExec(query string, arg interface{}) (sql.Result, error)

	NamedQuery(query string, arg interface{}) (*sqlx.Rows, error)
}

type JsonObject map[string]interface{}

func (j *JsonObject) Scan(src any) error {
	jsonBytes, ok := src.([]byte)

	if !ok {
		return fmt.Errorf("expected []byte, got %T", src)
	}

	err := json.Unmarshal(jsonBytes, &j)

	if err != nil {
		return err
	}

	return nil
}

func (j *JsonObject) Value() (driver.Value, error) {
	return json.Marshal(j)
}

type JsonArray []map[string]interface{}

func (j *JsonArray) Scan(src any) error {
	jsonBytes, ok := src.([]byte)

	if !ok {
		return fmt.Errorf("expected []byte, got %T", src)
	}

	err := json.Unmarshal(jsonBytes, &j)

	if err != nil {
		return err
	}

	return nil
}

func (j *JsonArray) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func countFields(v any) int {
	return rvCountFields(reflect.ValueOf(v))
}

func rvCountFields(rv reflect.Value) (count int) {

	if rv.Kind() != reflect.Struct {
		return
	}

	fs := rv.NumField()
	count += fs

	for i := 0; i < fs; i++ {
		f := rv.Field(i)
		if rv.Type().Field(i).Anonymous {
			count-- // don't count embedded structs (listed as anonymous fields) as a field
		}

		// recurse each field to see if that field is also an embedded struct
		count += rvCountFields(f)
	}

	return
}

func BulkInsert[T model[P], P any](db *sqlx.DB, ctx context.Context, itemsToSave ...T) ([]*P, error) {
	tx, err := db.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}

	pgMaxParameterCount := 500 // TODO: offer as an option

	// we need to batch the inserts so num `items` * `item` struct field
	// count is less than postgres MaxParameterCount - this is conservative
	firstItem := itemsToSave[0]
	itemPropertyCount := countFields(*firstItem)

	maxBatch := pgMaxParameterCount / itemPropertyCount
	items := make([]*P, 0, len(itemsToSave))
	for i := 0; i < len(itemsToSave); i += maxBatch {
		end := i + maxBatch

		if end > len(itemsToSave) {
			end = len(itemsToSave)
		}

		// Using an anonymous function to ensure that the rows are closed as we go!
		err := func() error {
			firstItemSql := firstItem.InsertQuery()
			rows, err := sqlx.NamedQueryContext(ctx, tx, firstItemSql, itemsToSave[i:end])
			if err != nil {
				tx.Rollback()
				return err
			}
			defer rows.Close()

			insertedRowIdx := 0
			batchItems := make([]*P, len(itemsToSave[i:end]))
			for rows.Next() {
				newItem := new(P)
				if err := rows.StructScan(newItem); err != nil {
					tx.Rollback()
					return err
				}
				batchItems[insertedRowIdx] = newItem
				insertedRowIdx++
			}
			items = append(items, batchItems...)
			return nil
		}()
		if err != nil {
			return nil, err
		}
	}
	tx.Commit()

	return items, nil
}

// *************************
// starting partial Update!
// *************************

var SET_NULL = struct {
	STRING      string
	INT         int
	INT16       int16
	INT32       int32
	INT64       int64
	FLOAT32     float32
	FLOAT64     float64
	BOOL        bool
	TIME        time.Time
	BYTE        byte
	JSON_RAW    string // json.RawMessage
	JSON_ARRAY  string // lib.JsonArray
	JSON_OBJECT string // lib.JsonObject
}{
	STRING:      "",
	INT:         0,
	INT16:       0,
	INT32:       0,
	INT64:       0,
	FLOAT32:     0.0,
	FLOAT64:     0.0,
	BOOL:        false,
	TIME:        time.Time{},
	BYTE:        0,
	JSON_RAW:    "", // json.RawMessage{},
	JSON_ARRAY:  "", // lib.JsonArray{},
	JSON_OBJECT: "", // lib.JsonObject{},
}

type UpdateObjectMetadata struct {
	DbName        string
	FieldValue    any
	FieldType     string
	ShouldSetNull bool
	FieldHasValue bool
}

func rUpdateMeta(rv reflect.Value) (fields map[string]UpdateObjectMetadata) {

	if rv.Kind() != reflect.Struct {
		return
	}

	fields = make(map[string]UpdateObjectMetadata)

	for i := 0; i < rv.NumField(); i++ {

		fieldName := rv.Type().Field(i).Name
		f := reflect.Indirect(rv).FieldByName(fieldName)

		if f.Kind() != reflect.Struct {
			fieldType := rv.Type().Field(i).Type.String()
			fieldTag := rv.Type().Field(i).Tag

			var shouldUpdate bool = false
			var shouldSetNull bool = false

			if !f.IsNil() {
				fieldVal := f.Interface()
				if fieldVal == &SET_NULL.STRING || fieldVal == &SET_NULL.INT || fieldVal == &SET_NULL.INT32 || fieldVal == &SET_NULL.INT64 || fieldVal == &SET_NULL.FLOAT32 || fieldVal == &SET_NULL.FLOAT64 || fieldVal == &SET_NULL.BOOL || fieldVal == &SET_NULL.TIME || fieldVal == &SET_NULL.BYTE || fieldVal == &SET_NULL.JSON_RAW {
					shouldSetNull = true
				}
				shouldUpdate = true
			}

			fields[fieldName] = UpdateObjectMetadata{
				DbName:        fieldTag.Get("db"),
				FieldType:     fieldType,
				ShouldSetNull: shouldSetNull,
				FieldHasValue: shouldUpdate,
			}
		}
	}

	for i := 0; i < rv.NumField(); i++ {
		f := rv.Field(i)

		// recurse each field to see if that field is also an embedded struct
		newfields := rUpdateMeta(f)

		// merge the new fields into the fields map
		for k, v := range newfields {
			fields[k] = v
		}
	}
	return fields
}

func getFieldMetaForUpdate[T model[P], P any](instance T) (fields map[string]UpdateObjectMetadata) {

	instancePtr := *instance
	fields = rUpdateMeta(reflect.ValueOf(instancePtr))

	return fields
}

func getDbFieldMeta[T model[P], P any](instance T) (fields map[string]UpdateObjectMetadata) {

	dbFields := getFieldMetaForUpdate[T](instance)
	fields = make(map[string]UpdateObjectMetadata)

	for _, v := range dbFields {
		if v.FieldHasValue {
			fields[v.DbName] = v
		}
	}
	return fields
}

// *************************
// errors
// *************************

var ErrNotFound = errors.New("entity not found")
var ErrFoundMultiple = errors.New("multiple matching entities")
