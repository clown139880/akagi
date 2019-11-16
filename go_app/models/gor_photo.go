// Package models includes the functions on the model Photo.
package models

import (
	"errors"
	"fmt"
	"log"
	"math"
	"strings"
	"time"

	"github.com/asaskevich/govalidator"
)

// set flags to output more detailed log
func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

type Photo struct {
	Id            int64     `json:"id,omitempty" db:"id" valid:"-"`
	Key           string    `json:"key,omitempty" db:"key" valid:"-"`
	IsLogo        bool      `json:"is_logo,omitempty" db:"is_logo" valid:"-"`
	Url           string    `json:"url,omitempty" db:"url" valid:"-"`
	PhotoableId   int64     `json:"photoable_id,omitempty" db:"photoable_id" valid:"-"`
	PhotoableType string    `json:"photoable_type,omitempty" db:"photoable_type" valid:"-"`
	CreatedAt     time.Time `json:"created_at,omitempty" db:"created_at" valid:"-"`
	UpdatedAt     time.Time `json:"updated_at,omitempty" db:"updated_at" valid:"-"`
}

// DataStruct for the pagination
type PhotoPage struct {
	WhereString string
	WhereParams []interface{}
	Order       map[string]string
	FirstId     int64
	LastId      int64
	PageNum     int
	PerPage     int
	TotalPages  int
	TotalItems  int64
	orderStr    string
}

// Current get the current page of PhotoPage object for pagination.
func (_p *PhotoPage) Current() ([]Photo, error) {
	if _, exist := _p.Order["id"]; !exist {
		return nil, errors.New("No id order specified in Order map")
	}
	err := _p.buildPageCount()
	if err != nil {
		return nil, fmt.Errorf("Calculate page count error: %v", err)
	}
	if _p.orderStr == "" {
		_p.buildOrder()
	}
	idStr, idParams := _p.buildIdRestrict("current")
	whereStr := fmt.Sprintf("%s %s %s LIMIT %v", _p.WhereString, idStr, _p.orderStr, _p.PerPage)
	whereParams := []interface{}{}
	whereParams = append(append(whereParams, _p.WhereParams...), idParams...)
	photos, err := FindPhotosWhere(whereStr, whereParams...)
	if err != nil {
		return nil, err
	}
	if len(photos) != 0 {
		_p.FirstId, _p.LastId = photos[0].Id, photos[len(photos)-1].Id
	}
	return photos, nil
}

// Previous get the previous page of PhotoPage object for pagination.
func (_p *PhotoPage) Previous() ([]Photo, error) {
	if _p.PageNum == 0 {
		return nil, errors.New("This's the first page, no previous page yet")
	}
	if _, exist := _p.Order["id"]; !exist {
		return nil, errors.New("No id order specified in Order map")
	}
	err := _p.buildPageCount()
	if err != nil {
		return nil, fmt.Errorf("Calculate page count error: %v", err)
	}
	if _p.orderStr == "" {
		_p.buildOrder()
	}
	idStr, idParams := _p.buildIdRestrict("previous")
	whereStr := fmt.Sprintf("%s %s %s LIMIT %v", _p.WhereString, idStr, _p.orderStr, _p.PerPage)
	whereParams := []interface{}{}
	whereParams = append(append(whereParams, _p.WhereParams...), idParams...)
	photos, err := FindPhotosWhere(whereStr, whereParams...)
	if err != nil {
		return nil, err
	}
	if len(photos) != 0 {
		_p.FirstId, _p.LastId = photos[0].Id, photos[len(photos)-1].Id
	}
	_p.PageNum -= 1
	return photos, nil
}

// Next get the next page of PhotoPage object for pagination.
func (_p *PhotoPage) Next() ([]Photo, error) {
	if _p.PageNum == _p.TotalPages-1 {
		return nil, errors.New("This's the last page, no next page yet")
	}
	if _, exist := _p.Order["id"]; !exist {
		return nil, errors.New("No id order specified in Order map")
	}
	err := _p.buildPageCount()
	if err != nil {
		return nil, fmt.Errorf("Calculate page count error: %v", err)
	}
	if _p.orderStr == "" {
		_p.buildOrder()
	}
	idStr, idParams := _p.buildIdRestrict("next")
	whereStr := fmt.Sprintf("%s %s %s LIMIT %v", _p.WhereString, idStr, _p.orderStr, _p.PerPage)
	whereParams := []interface{}{}
	whereParams = append(append(whereParams, _p.WhereParams...), idParams...)
	photos, err := FindPhotosWhere(whereStr, whereParams...)
	if err != nil {
		return nil, err
	}
	if len(photos) != 0 {
		_p.FirstId, _p.LastId = photos[0].Id, photos[len(photos)-1].Id
	}
	_p.PageNum += 1
	return photos, nil
}

// GetPage is a helper function for the PhotoPage object to return a corresponding page due to
// the parameter passed in, i.e. one of "previous, current or next".
func (_p *PhotoPage) GetPage(direction string) (ps []Photo, err error) {
	switch direction {
	case "previous":
		ps, _ = _p.Previous()
	case "next":
		ps, _ = _p.Next()
	case "current":
		ps, _ = _p.Current()
	default:
		return nil, errors.New("Error: wrong dircetion! None of previous, current or next!")
	}
	return
}

// buildOrder is for PhotoPage object to build a SQL ORDER BY clause.
func (_p *PhotoPage) buildOrder() {
	tempList := []string{}
	for k, v := range _p.Order {
		tempList = append(tempList, fmt.Sprintf("%v %v", k, v))
	}
	_p.orderStr = " ORDER BY " + strings.Join(tempList, ", ")
}

// buildIdRestrict is for PhotoPage object to build a SQL clause for ID restriction,
// implementing a simple keyset style pagination.
func (_p *PhotoPage) buildIdRestrict(direction string) (idStr string, idParams []interface{}) {
	switch direction {
	case "previous":
		if strings.ToLower(_p.Order["id"]) == "desc" {
			idStr += "id > ? "
			idParams = append(idParams, _p.FirstId)
		} else {
			idStr += "id < ? "
			idParams = append(idParams, _p.FirstId)
		}
	case "current":
		// trick to make Where function work
		if _p.PageNum == 0 && _p.FirstId == 0 && _p.LastId == 0 {
			idStr += "id > ? "
			idParams = append(idParams, 0)
		} else {
			if strings.ToLower(_p.Order["id"]) == "desc" {
				idStr += "id <= ? AND id >= ? "
				idParams = append(idParams, _p.FirstId, _p.LastId)
			} else {
				idStr += "id >= ? AND id <= ? "
				idParams = append(idParams, _p.FirstId, _p.LastId)
			}
		}
	case "next":
		if strings.ToLower(_p.Order["id"]) == "desc" {
			idStr += "id < ? "
			idParams = append(idParams, _p.LastId)
		} else {
			idStr += "id > ? "
			idParams = append(idParams, _p.LastId)
		}
	}
	if _p.WhereString != "" {
		idStr = " AND " + idStr
	}
	return
}

// buildPageCount calculate the TotalItems/TotalPages for the PhotoPage object.
func (_p *PhotoPage) buildPageCount() error {
	count, err := PhotoCountWhere(_p.WhereString, _p.WhereParams...)
	if err != nil {
		return err
	}
	_p.TotalItems = count
	if _p.PerPage == 0 {
		_p.PerPage = 10
	}
	_p.TotalPages = int(math.Ceil(float64(_p.TotalItems) / float64(_p.PerPage)))
	return nil
}

// FindPhoto find a single photo by an ID.
func FindPhoto(id int64) (*Photo, error) {
	if id == 0 {
		return nil, errors.New("Invalid ID: it can't be zero")
	}
	_photo := Photo{}
	err := DB.Get(&_photo, DB.Rebind(`SELECT COALESCE(photos.key, '') AS 'key', COALESCE(photos.is_logo, FALSE) AS is_logo, COALESCE(photos.url, '') AS url, COALESCE(photos.photoable_id, 0) AS photoable_id, COALESCE(photos.photoable_type, '') AS photoable_type, photos.id, photos.created_at, photos.updated_at FROM photos WHERE photos.id = ? LIMIT 1`), id)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return nil, err
	}
	return &_photo, nil
}

// FirstPhoto find the first one photo by ID ASC order.
func FirstPhoto() (*Photo, error) {
	_photo := Photo{}
	err := DB.Get(&_photo, DB.Rebind(`SELECT COALESCE(photos.key, '') AS 'key', COALESCE(photos.is_logo, FALSE) AS is_logo, COALESCE(photos.url, '') AS url, COALESCE(photos.photoable_id, 0) AS photoable_id, COALESCE(photos.photoable_type, '') AS photoable_type, photos.id, photos.created_at, photos.updated_at FROM photos ORDER BY photos.id ASC LIMIT 1`))
	if err != nil {
		log.Printf("Error: %v\n", err)
		return nil, err
	}
	return &_photo, nil
}

// FirstPhotos find the first N photos by ID ASC order.
func FirstPhotos(n uint32) ([]Photo, error) {
	_photos := []Photo{}
	sql := fmt.Sprintf("SELECT COALESCE(photos.key, '') AS 'key', COALESCE(photos.is_logo, FALSE) AS is_logo, COALESCE(photos.url, '') AS url, COALESCE(photos.photoable_id, 0) AS photoable_id, COALESCE(photos.photoable_type, '') AS photoable_type, photos.id, photos.created_at, photos.updated_at FROM photos ORDER BY photos.id ASC LIMIT %v", n)
	err := DB.Select(&_photos, DB.Rebind(sql))
	if err != nil {
		log.Printf("Error: %v\n", err)
		return nil, err
	}
	return _photos, nil
}

// LastPhoto find the last one photo by ID DESC order.
func LastPhoto() (*Photo, error) {
	_photo := Photo{}
	err := DB.Get(&_photo, DB.Rebind(`SELECT COALESCE(photos.key, '') AS 'key', COALESCE(photos.is_logo, FALSE) AS is_logo, COALESCE(photos.url, '') AS url, COALESCE(photos.photoable_id, 0) AS photoable_id, COALESCE(photos.photoable_type, '') AS photoable_type, photos.id, photos.created_at, photos.updated_at FROM photos ORDER BY photos.id DESC LIMIT 1`))
	if err != nil {
		log.Printf("Error: %v\n", err)
		return nil, err
	}
	return &_photo, nil
}

// LastPhotos find the last N photos by ID DESC order.
func LastPhotos(n uint32) ([]Photo, error) {
	_photos := []Photo{}
	sql := fmt.Sprintf("SELECT COALESCE(photos.key, '') AS 'key', COALESCE(photos.is_logo, FALSE) AS is_logo, COALESCE(photos.url, '') AS url, COALESCE(photos.photoable_id, 0) AS photoable_id, COALESCE(photos.photoable_type, '') AS photoable_type, photos.id, photos.created_at, photos.updated_at FROM photos ORDER BY photos.id DESC LIMIT %v", n)
	err := DB.Select(&_photos, DB.Rebind(sql))
	if err != nil {
		log.Printf("Error: %v\n", err)
		return nil, err
	}
	return _photos, nil
}

// FindPhotos find one or more photos by the given ID(s).
func FindPhotos(ids ...int64) ([]Photo, error) {
	if len(ids) == 0 {
		msg := "At least one or more ids needed"
		log.Println(msg)
		return nil, errors.New(msg)
	}
	_photos := []Photo{}
	idsHolder := strings.Repeat(",?", len(ids)-1)
	sql := DB.Rebind(fmt.Sprintf(`SELECT COALESCE(photos.key, '') AS 'key', COALESCE(photos.is_logo, FALSE) AS is_logo, COALESCE(photos.url, '') AS url, COALESCE(photos.photoable_id, 0) AS photoable_id, COALESCE(photos.photoable_type, '') AS photoable_type, photos.id, photos.created_at, photos.updated_at FROM photos WHERE photos.id IN (?%s)`, idsHolder))
	idsT := []interface{}{}
	for _, id := range ids {
		idsT = append(idsT, interface{}(id))
	}
	err := DB.Select(&_photos, sql, idsT...)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return nil, err
	}
	return _photos, nil
}

// FindPhotoBy find a single photo by a field name and a value.
func FindPhotoBy(field string, val interface{}) (*Photo, error) {
	_photo := Photo{}
	sqlFmt := `SELECT COALESCE(photos.key, '') AS 'key', COALESCE(photos.is_logo, FALSE) AS is_logo, COALESCE(photos.url, '') AS url, COALESCE(photos.photoable_id, 0) AS photoable_id, COALESCE(photos.photoable_type, '') AS photoable_type, photos.id, photos.created_at, photos.updated_at FROM photos WHERE %s = ? LIMIT 1`
	sqlStr := fmt.Sprintf(sqlFmt, field)
	err := DB.Get(&_photo, DB.Rebind(sqlStr), val)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return nil, err
	}
	return &_photo, nil
}

// FindPhotosBy find all photos by a field name and a value.
func FindPhotosBy(field string, val interface{}) (_photos []Photo, err error) {
	sqlFmt := `SELECT COALESCE(photos.key, '') AS 'key', COALESCE(photos.is_logo, FALSE) AS is_logo, COALESCE(photos.url, '') AS url, COALESCE(photos.photoable_id, 0) AS photoable_id, COALESCE(photos.photoable_type, '') AS photoable_type, photos.id, photos.created_at, photos.updated_at FROM photos WHERE %s = ?`
	sqlStr := fmt.Sprintf(sqlFmt, field)
	err = DB.Select(&_photos, DB.Rebind(sqlStr), val)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return nil, err
	}
	return _photos, nil
}

// AllPhotos get all the Photo records.
func AllPhotos() (photos []Photo, err error) {
	err = DB.Select(&photos, "SELECT COALESCE(photos.key, '') AS 'key', COALESCE(photos.is_logo, FALSE) AS is_logo, COALESCE(photos.url, '') AS url, COALESCE(photos.photoable_id, 0) AS photoable_id, COALESCE(photos.photoable_type, '') AS photoable_type, photos.id, photos.created_at, photos.updated_at FROM photos")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return photos, nil
}

// PhotoCount get the count of all the Photo records.
func PhotoCount() (c int64, err error) {
	err = DB.Get(&c, "SELECT count(*) FROM photos")
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return c, nil
}

// PhotoCountWhere get the count of all the Photo records with a where clause.
func PhotoCountWhere(where string, args ...interface{}) (c int64, err error) {
	sql := "SELECT count(*) FROM photos"
	if len(where) > 0 {
		sql = sql + " WHERE " + where
	}
	stmt, err := DB.Preparex(DB.Rebind(sql))
	if err != nil {
		log.Println(err)
		return 0, err
	}
	err = stmt.Get(&c, args...)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return c, nil
}

// PhotoIncludesWhere get the Photo associated models records, currently it's not same as the corresponding "includes" function but "preload" instead in Ruby on Rails. It means that the "sql" should be restricted on Photo model.
func PhotoIncludesWhere(assocs []string, sql string, args ...interface{}) (_photos []Photo, err error) {
	_photos, err = FindPhotosWhere(sql, args...)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if len(assocs) == 0 {
		log.Println("No associated fields ard specified")
		return _photos, err
	}
	if len(_photos) <= 0 {
		return nil, errors.New("No results available")
	}
	ids := make([]interface{}, len(_photos))
	for _, v := range _photos {
		ids = append(ids, interface{}(v.Id))
	}
	return _photos, nil
}

// PhotoIds get all the IDs of Photo records.
func PhotoIds() (ids []int64, err error) {
	err = DB.Select(&ids, "SELECT id FROM photos")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return ids, nil
}

// PhotoIdsWhere get all the IDs of Photo records by where restriction.
func PhotoIdsWhere(where string, args ...interface{}) ([]int64, error) {
	ids, err := PhotoIntCol("id", where, args...)
	return ids, err
}

// PhotoIntCol get some int64 typed column of Photo by where restriction.
func PhotoIntCol(col, where string, args ...interface{}) (intColRecs []int64, err error) {
	sql := "SELECT " + col + " FROM photos"
	if len(where) > 0 {
		sql = sql + " WHERE " + where
	}
	stmt, err := DB.Preparex(DB.Rebind(sql))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	err = stmt.Select(&intColRecs, args...)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return intColRecs, nil
}

// PhotoStrCol get some string typed column of Photo by where restriction.
func PhotoStrCol(col, where string, args ...interface{}) (strColRecs []string, err error) {
	sql := "SELECT " + col + " FROM photos"
	if len(where) > 0 {
		sql = sql + " WHERE " + where
	}
	stmt, err := DB.Preparex(DB.Rebind(sql))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	err = stmt.Select(&strColRecs, args...)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return strColRecs, nil
}

// FindPhotosWhere query use a partial SQL clause that usually following after WHERE
// with placeholders, eg: FindUsersWhere("first_name = ? AND age > ?", "John", 18)
// will return those records in the table "users" whose first_name is "John" and age elder than 18.
func FindPhotosWhere(where string, args ...interface{}) (photos []Photo, err error) {
	sql := "SELECT COALESCE(photos.key, '') AS `key`, COALESCE(photos.is_logo, FALSE) AS is_logo, COALESCE(photos.url, '') AS url, COALESCE(photos.photoable_id, 0) AS photoable_id, COALESCE(photos.photoable_type, '') AS photoable_type, photos.id, photos.created_at, photos.updated_at FROM photos"
	if len(where) > 0 {
		sql = sql + " WHERE " + where
	}
	stmt, err := DB.Preparex(DB.Rebind(sql))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	err = stmt.Select(&photos, args...)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return photos, nil
}

// FindPhotoBySql query use a complete SQL clause
// with placeholders, eg: FindUserBySql("SELECT * FROM users WHERE first_name = ? AND age > ? ORDER BY DESC LIMIT 1", "John", 18)
// will return only One record in the table "users" whose first_name is "John" and age elder than 18.
func FindPhotoBySql(sql string, args ...interface{}) (*Photo, error) {
	stmt, err := DB.Preparex(DB.Rebind(sql))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	_photo := &Photo{}
	err = stmt.Get(_photo, args...)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return _photo, nil
}

// FindPhotosBySql query use a complete SQL clause
// with placeholders, eg: FindUsersBySql("SELECT * FROM users WHERE first_name = ? AND age > ?", "John", 18)
// will return those records in the table "users" whose first_name is "John" and age elder than 18.
func FindPhotosBySql(sql string, args ...interface{}) (photos []Photo, err error) {
	stmt, err := DB.Preparex(DB.Rebind(sql))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	err = stmt.Select(&photos, args...)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return photos, nil
}

// CreatePhoto use a named params to create a single Photo record.
// A named params is key-value map like map[string]interface{}{"first_name": "John", "age": 23} .
func CreatePhoto(am map[string]interface{}) (int64, error) {
	if len(am) == 0 {
		return 0, fmt.Errorf("Zero key in the attributes map!")
	}
	t := time.Now()
	for _, v := range []string{"created_at", "updated_at"} {
		if am[v] == nil {
			am[v] = t
		}
	}
	keys := allKeys(am)
	sqlFmt := `INSERT INTO photos (%s) VALUES (%s)`
	sql := fmt.Sprintf(sqlFmt, strings.Join(keys, ","), ":"+strings.Join(keys, ",:"))
	result, err := DB.NamedExec(sql, am)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	lastId, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return lastId, nil
}

// Create is a method for Photo to create a record.
func (_photo *Photo) Create() (int64, error) {
	ok, err := govalidator.ValidateStruct(_photo)
	if !ok {
		errMsg := "Validate Photo struct error: Unknown error"
		if err != nil {
			errMsg = "Validate Photo struct error: " + err.Error()
		}
		log.Println(errMsg)
		return 0, errors.New(errMsg)
	}
	t := time.Now()
	_photo.CreatedAt = t
	_photo.UpdatedAt = t
	sql := "INSERT INTO photos (`key`,is_logo,url,photoable_id,photoable_type,created_at,updated_at) VALUES (:key,:is_logo,:url,:photoable_id,:photoable_type,:created_at,:updated_at)"
	result, err := DB.NamedExec(sql, _photo)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	lastId, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return lastId, nil
}

// Destroy is method used for a Photo object to be destroyed.
func (_photo *Photo) Destroy() error {
	if _photo.Id == 0 {
		return errors.New("Invalid Id field: it can't be a zero value")
	}
	err := DestroyPhoto(_photo.Id)
	return err
}

// DestroyPhoto will destroy a Photo record specified by the id parameter.
func DestroyPhoto(id int64) error {
	stmt, err := DB.Preparex(DB.Rebind(`DELETE FROM photos WHERE id = ?`))
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

// DestroyPhotos will destroy Photo records those specified by the ids parameters.
func DestroyPhotos(ids ...int64) (int64, error) {
	if len(ids) == 0 {
		msg := "At least one or more ids needed"
		log.Println(msg)
		return 0, errors.New(msg)
	}
	idsHolder := strings.Repeat(",?", len(ids)-1)
	sql := fmt.Sprintf(`DELETE FROM photos WHERE id IN (?%s)`, idsHolder)
	idsT := []interface{}{}
	for _, id := range ids {
		idsT = append(idsT, interface{}(id))
	}
	stmt, err := DB.Preparex(DB.Rebind(sql))
	result, err := stmt.Exec(idsT...)
	if err != nil {
		return 0, err
	}
	cnt, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return cnt, nil
}

// DestroyPhotosWhere delete records by a where clause restriction.
// e.g. DestroyPhotosWhere("name = ?", "John")
// And this func will not call the association dependent action
func DestroyPhotosWhere(where string, args ...interface{}) (int64, error) {
	sql := `DELETE FROM photos WHERE `
	if len(where) > 0 {
		sql = sql + where
	} else {
		return 0, errors.New("No WHERE conditions provided")
	}
	stmt, err := DB.Preparex(DB.Rebind(sql))
	result, err := stmt.Exec(args...)
	if err != nil {
		return 0, err
	}
	cnt, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return cnt, nil
}

// Save method is used for a Photo object to update an existed record mainly.
// If no id provided a new record will be created. FIXME: A UPSERT action will be implemented further.
func (_photo *Photo) Save() error {
	ok, err := govalidator.ValidateStruct(_photo)
	if !ok {
		errMsg := "Validate Photo struct error: Unknown error"
		if err != nil {
			errMsg = "Validate Photo struct error: " + err.Error()
		}
		log.Println(errMsg)
		return errors.New(errMsg)
	}
	if _photo.Id == 0 {
		_, err = _photo.Create()
		return err
	}
	_photo.UpdatedAt = time.Now()
	sqlFmt := `UPDATE photos SET %s WHERE id = %v`
	sqlStr := fmt.Sprintf(sqlFmt, "key = :key, is_logo = :is_logo, url = :url, photoable_id = :photoable_id, photoable_type = :photoable_type, updated_at = :updated_at", _photo.Id)
	_, err = DB.NamedExec(sqlStr, _photo)
	return err
}

// UpdatePhoto is used to update a record with a id and map[string]interface{} typed key-value parameters.
func UpdatePhoto(id int64, am map[string]interface{}) error {
	if len(am) == 0 {
		return errors.New("Zero key in the attributes map!")
	}
	am["updated_at"] = time.Now()
	keys := allKeys(am)
	sqlFmt := `UPDATE photos SET %s WHERE id = %v`
	setKeysArr := []string{}
	for _, v := range keys {
		s := fmt.Sprintf(" %s = :%s", v, v)
		setKeysArr = append(setKeysArr, s)
	}
	sqlStr := fmt.Sprintf(sqlFmt, strings.Join(setKeysArr, ", "), id)
	_, err := DB.NamedExec(sqlStr, am)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// Update is a method used to update a Photo record with the map[string]interface{} typed key-value parameters.
func (_photo *Photo) Update(am map[string]interface{}) error {
	if _photo.Id == 0 {
		return errors.New("Invalid Id field: it can't be a zero value")
	}
	err := UpdatePhoto(_photo.Id, am)
	return err
}

// UpdateAttributes method is supposed to be used to update Photo records as corresponding update_attributes in Ruby on Rails.
func (_photo *Photo) UpdateAttributes(am map[string]interface{}) error {
	if _photo.Id == 0 {
		return errors.New("Invalid Id field: it can't be a zero value")
	}
	err := UpdatePhoto(_photo.Id, am)
	return err
}

// UpdateColumns method is supposed to be used to update Photo records as corresponding update_columns in Ruby on Rails.
func (_photo *Photo) UpdateColumns(am map[string]interface{}) error {
	if _photo.Id == 0 {
		return errors.New("Invalid Id field: it can't be a zero value")
	}
	err := UpdatePhoto(_photo.Id, am)
	return err
}

// UpdatePhotosBySql is used to update Photo records by a SQL clause
// using the '?' binding syntax.
func UpdatePhotosBySql(sql string, args ...interface{}) (int64, error) {
	if sql == "" {
		return 0, errors.New("A blank SQL clause")
	}
	stmt, err := DB.Preparex(DB.Rebind(sql))
	result, err := stmt.Exec(args...)
	if err != nil {
		return 0, err
	}
	cnt, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return cnt, nil
}
