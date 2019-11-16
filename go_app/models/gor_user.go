

// Package models includes the functions on the model User.
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

type User struct {
	Id int64 `json:"id,omitempty" db:"id" valid:"-"`
Name string `json:"name,omitempty" db:"name" valid:"-"`
Email string `json:"email,omitempty" db:"email" valid:"-"`
CreatedAt time.Time `json:"created_at,omitempty" db:"created_at" valid:"-"`
UpdatedAt time.Time `json:"updated_at,omitempty" db:"updated_at" valid:"-"`
PasswordDigest string `json:"password_digest,omitempty" db:"password_digest" valid:"-"`
RememberDigest string `json:"remember_digest,omitempty" db:"remember_digest" valid:"-"`
Admin bool `json:"admin,omitempty" db:"admin" valid:"-"`
ActivationDigest string `json:"activation_digest,omitempty" db:"activation_digest" valid:"-"`
Activated bool `json:"activated,omitempty" db:"activated" valid:"-"`
ActivatedAt time.Time `json:"activated_at,omitempty" db:"activated_at" valid:"-"`
ResetDigest string `json:"reset_digest,omitempty" db:"reset_digest" valid:"-"`
ResetSentAt time.Time `json:"reset_sent_at,omitempty" db:"reset_sent_at" valid:"-"`
Weibo string `json:"weibo,omitempty" db:"weibo" valid:"-"`
WeixinOpenid string `json:"weixin_openid,omitempty" db:"weixin_openid" valid:"-"`
Avatar string `json:"avatar,omitempty" db:"avatar" valid:"-"`
Posts []Post `json:"posts,omitempty" db:"posts" valid:"-"`
}

// DataStruct for the pagination
type UserPage struct {
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

// Current get the current page of UserPage object for pagination.
func (_p *UserPage) Current() ([]User, error) {
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
	users, err := FindUsersWhere(whereStr, whereParams...)
	if err != nil {
		return nil, err
	}
	if len(users) != 0 {
		_p.FirstId, _p.LastId = users[0].Id, users[len(users)-1].Id
	}
	return users, nil
}

// Previous get the previous page of UserPage object for pagination.
func (_p *UserPage) Previous() ([]User, error) {
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
	users, err := FindUsersWhere(whereStr, whereParams...)
	if err != nil {
		return nil, err
	}
	if len(users) != 0 {
		_p.FirstId, _p.LastId = users[0].Id, users[len(users)-1].Id
	}
	_p.PageNum -= 1
	return users, nil
}

// Next get the next page of UserPage object for pagination.
func (_p *UserPage) Next() ([]User, error) {
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
	users, err := FindUsersWhere(whereStr, whereParams...)
	if err != nil {
		return nil, err
	}
	if len(users) != 0 {
		_p.FirstId, _p.LastId = users[0].Id, users[len(users)-1].Id
	}
	_p.PageNum += 1
	return users, nil
}

// GetPage is a helper function for the UserPage object to return a corresponding page due to
// the parameter passed in, i.e. one of "previous, current or next".
func (_p *UserPage) GetPage(direction string) (ps []User, err error) {
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

// buildOrder is for UserPage object to build a SQL ORDER BY clause.
func (_p *UserPage) buildOrder() {
	tempList := []string{}
	for k, v := range _p.Order {
		tempList = append(tempList, fmt.Sprintf("%v %v", k, v))
	}
	_p.orderStr = " ORDER BY " + strings.Join(tempList, ", ")
}

// buildIdRestrict is for UserPage object to build a SQL clause for ID restriction,
// implementing a simple keyset style pagination.
func (_p *UserPage) buildIdRestrict(direction string) (idStr string, idParams []interface{}) {
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

// buildPageCount calculate the TotalItems/TotalPages for the UserPage object.
func (_p *UserPage) buildPageCount() error {
	count, err := UserCountWhere(_p.WhereString, _p.WhereParams...)
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


// FindUser find a single user by an ID.
func FindUser(id int64) (*User, error) {
	if id == 0 {
		return nil, errors.New("Invalid ID: it can't be zero")
	}
	_user := User{}
	err := DB.Get(&_user, DB.Rebind(`SELECT COALESCE(users.name, '') AS name, COALESCE(users.email, '') AS email, COALESCE(users.password_digest, '') AS password_digest, COALESCE(users.remember_digest, '') AS remember_digest, COALESCE(users.admin, FALSE) AS admin, COALESCE(users.activation_digest, '') AS activation_digest, COALESCE(users.activated, FALSE) AS activated, COALESCE(users.activated_at, CONVERT_TZ('0001-01-01 00:00:00','+00:00','UTC')) AS activated_at, COALESCE(users.reset_digest, '') AS reset_digest, COALESCE(users.reset_sent_at, CONVERT_TZ('0001-01-01 00:00:00','+00:00','UTC')) AS reset_sent_at, COALESCE(users.weibo, '') AS weibo, COALESCE(users.weixin_openid, '') AS weixin_openid, COALESCE(users.avatar, '') AS avatar, users.id, users.created_at, users.updated_at FROM users WHERE users.id = ? LIMIT 1`), id)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return nil, err
	}
	return &_user, nil
}

// FirstUser find the first one user by ID ASC order.
func FirstUser() (*User, error) {
	_user := User{}
	err := DB.Get(&_user, DB.Rebind(`SELECT COALESCE(users.name, '') AS name, COALESCE(users.email, '') AS email, COALESCE(users.password_digest, '') AS password_digest, COALESCE(users.remember_digest, '') AS remember_digest, COALESCE(users.admin, FALSE) AS admin, COALESCE(users.activation_digest, '') AS activation_digest, COALESCE(users.activated, FALSE) AS activated, COALESCE(users.activated_at, CONVERT_TZ('0001-01-01 00:00:00','+00:00','UTC')) AS activated_at, COALESCE(users.reset_digest, '') AS reset_digest, COALESCE(users.reset_sent_at, CONVERT_TZ('0001-01-01 00:00:00','+00:00','UTC')) AS reset_sent_at, COALESCE(users.weibo, '') AS weibo, COALESCE(users.weixin_openid, '') AS weixin_openid, COALESCE(users.avatar, '') AS avatar, users.id, users.created_at, users.updated_at FROM users ORDER BY users.id ASC LIMIT 1`))
	if err != nil {
		log.Printf("Error: %v\n", err)
		return nil, err
	}
	return &_user, nil
}

// FirstUsers find the first N users by ID ASC order.
func FirstUsers(n uint32) ([]User, error) {
	_users := []User{}
	sql := fmt.Sprintf("SELECT COALESCE(users.name, '') AS name, COALESCE(users.email, '') AS email, COALESCE(users.password_digest, '') AS password_digest, COALESCE(users.remember_digest, '') AS remember_digest, COALESCE(users.admin, FALSE) AS admin, COALESCE(users.activation_digest, '') AS activation_digest, COALESCE(users.activated, FALSE) AS activated, COALESCE(users.activated_at, CONVERT_TZ('0001-01-01 00:00:00','+00:00','UTC')) AS activated_at, COALESCE(users.reset_digest, '') AS reset_digest, COALESCE(users.reset_sent_at, CONVERT_TZ('0001-01-01 00:00:00','+00:00','UTC')) AS reset_sent_at, COALESCE(users.weibo, '') AS weibo, COALESCE(users.weixin_openid, '') AS weixin_openid, COALESCE(users.avatar, '') AS avatar, users.id, users.created_at, users.updated_at FROM users ORDER BY users.id ASC LIMIT %v", n)
	err := DB.Select(&_users, DB.Rebind(sql))
	if err != nil {
		log.Printf("Error: %v\n", err)
		return nil, err
	}
	return _users, nil
}

// LastUser find the last one user by ID DESC order.
func LastUser() (*User, error) {
	_user := User{}
	err := DB.Get(&_user, DB.Rebind(`SELECT COALESCE(users.name, '') AS name, COALESCE(users.email, '') AS email, COALESCE(users.password_digest, '') AS password_digest, COALESCE(users.remember_digest, '') AS remember_digest, COALESCE(users.admin, FALSE) AS admin, COALESCE(users.activation_digest, '') AS activation_digest, COALESCE(users.activated, FALSE) AS activated, COALESCE(users.activated_at, CONVERT_TZ('0001-01-01 00:00:00','+00:00','UTC')) AS activated_at, COALESCE(users.reset_digest, '') AS reset_digest, COALESCE(users.reset_sent_at, CONVERT_TZ('0001-01-01 00:00:00','+00:00','UTC')) AS reset_sent_at, COALESCE(users.weibo, '') AS weibo, COALESCE(users.weixin_openid, '') AS weixin_openid, COALESCE(users.avatar, '') AS avatar, users.id, users.created_at, users.updated_at FROM users ORDER BY users.id DESC LIMIT 1`))
	if err != nil {
		log.Printf("Error: %v\n", err)
		return nil, err
	}
	return &_user, nil
}

// LastUsers find the last N users by ID DESC order.
func LastUsers(n uint32) ([]User, error) {
	_users := []User{}
	sql := fmt.Sprintf("SELECT COALESCE(users.name, '') AS name, COALESCE(users.email, '') AS email, COALESCE(users.password_digest, '') AS password_digest, COALESCE(users.remember_digest, '') AS remember_digest, COALESCE(users.admin, FALSE) AS admin, COALESCE(users.activation_digest, '') AS activation_digest, COALESCE(users.activated, FALSE) AS activated, COALESCE(users.activated_at, CONVERT_TZ('0001-01-01 00:00:00','+00:00','UTC')) AS activated_at, COALESCE(users.reset_digest, '') AS reset_digest, COALESCE(users.reset_sent_at, CONVERT_TZ('0001-01-01 00:00:00','+00:00','UTC')) AS reset_sent_at, COALESCE(users.weibo, '') AS weibo, COALESCE(users.weixin_openid, '') AS weixin_openid, COALESCE(users.avatar, '') AS avatar, users.id, users.created_at, users.updated_at FROM users ORDER BY users.id DESC LIMIT %v", n)
	err := DB.Select(&_users, DB.Rebind(sql))
	if err != nil {
		log.Printf("Error: %v\n", err)
		return nil, err
	}
	return _users, nil
}

// FindUsers find one or more users by the given ID(s).
func FindUsers(ids ...int64) ([]User, error) {
	if len(ids) == 0 {
		msg := "At least one or more ids needed"
		log.Println(msg)
		return nil, errors.New(msg)
	}
	_users := []User{}
	idsHolder := strings.Repeat(",?", len(ids)-1)
	sql := DB.Rebind(fmt.Sprintf(`SELECT COALESCE(users.name, '') AS name, COALESCE(users.email, '') AS email, COALESCE(users.password_digest, '') AS password_digest, COALESCE(users.remember_digest, '') AS remember_digest, COALESCE(users.admin, FALSE) AS admin, COALESCE(users.activation_digest, '') AS activation_digest, COALESCE(users.activated, FALSE) AS activated, COALESCE(users.activated_at, CONVERT_TZ('0001-01-01 00:00:00','+00:00','UTC')) AS activated_at, COALESCE(users.reset_digest, '') AS reset_digest, COALESCE(users.reset_sent_at, CONVERT_TZ('0001-01-01 00:00:00','+00:00','UTC')) AS reset_sent_at, COALESCE(users.weibo, '') AS weibo, COALESCE(users.weixin_openid, '') AS weixin_openid, COALESCE(users.avatar, '') AS avatar, users.id, users.created_at, users.updated_at FROM users WHERE users.id IN (?%s)`, idsHolder))
	idsT := []interface{}{}
	for _,id := range ids {
		idsT = append(idsT, interface{}(id))
	}
	err := DB.Select(&_users, sql, idsT...)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return nil, err
	}
	return _users, nil
}

// FindUserBy find a single user by a field name and a value.
func FindUserBy(field string, val interface{}) (*User, error) {
	_user := User{}
	sqlFmt := `SELECT COALESCE(users.name, '') AS name, COALESCE(users.email, '') AS email, COALESCE(users.password_digest, '') AS password_digest, COALESCE(users.remember_digest, '') AS remember_digest, COALESCE(users.admin, FALSE) AS admin, COALESCE(users.activation_digest, '') AS activation_digest, COALESCE(users.activated, FALSE) AS activated, COALESCE(users.activated_at, CONVERT_TZ('0001-01-01 00:00:00','+00:00','UTC')) AS activated_at, COALESCE(users.reset_digest, '') AS reset_digest, COALESCE(users.reset_sent_at, CONVERT_TZ('0001-01-01 00:00:00','+00:00','UTC')) AS reset_sent_at, COALESCE(users.weibo, '') AS weibo, COALESCE(users.weixin_openid, '') AS weixin_openid, COALESCE(users.avatar, '') AS avatar, users.id, users.created_at, users.updated_at FROM users WHERE %s = ? LIMIT 1`
	sqlStr := fmt.Sprintf(sqlFmt, field)
	err := DB.Get(&_user, DB.Rebind(sqlStr), val)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return nil, err
	}
	return &_user, nil
}

// FindUsersBy find all users by a field name and a value.
func FindUsersBy(field string, val interface{}) (_users []User, err error) {
	sqlFmt := `SELECT COALESCE(users.name, '') AS name, COALESCE(users.email, '') AS email, COALESCE(users.password_digest, '') AS password_digest, COALESCE(users.remember_digest, '') AS remember_digest, COALESCE(users.admin, FALSE) AS admin, COALESCE(users.activation_digest, '') AS activation_digest, COALESCE(users.activated, FALSE) AS activated, COALESCE(users.activated_at, CONVERT_TZ('0001-01-01 00:00:00','+00:00','UTC')) AS activated_at, COALESCE(users.reset_digest, '') AS reset_digest, COALESCE(users.reset_sent_at, CONVERT_TZ('0001-01-01 00:00:00','+00:00','UTC')) AS reset_sent_at, COALESCE(users.weibo, '') AS weibo, COALESCE(users.weixin_openid, '') AS weixin_openid, COALESCE(users.avatar, '') AS avatar, users.id, users.created_at, users.updated_at FROM users WHERE %s = ?`
	sqlStr := fmt.Sprintf(sqlFmt, field)
	err = DB.Select(&_users, DB.Rebind(sqlStr), val)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return nil, err
	}
	return _users, nil
}

// AllUsers get all the User records.
func AllUsers() (users []User, err error) {
	err = DB.Select(&users, "SELECT COALESCE(users.name, '') AS name, COALESCE(users.email, '') AS email, COALESCE(users.password_digest, '') AS password_digest, COALESCE(users.remember_digest, '') AS remember_digest, COALESCE(users.admin, FALSE) AS admin, COALESCE(users.activation_digest, '') AS activation_digest, COALESCE(users.activated, FALSE) AS activated, COALESCE(users.activated_at, CONVERT_TZ('0001-01-01 00:00:00','+00:00','UTC')) AS activated_at, COALESCE(users.reset_digest, '') AS reset_digest, COALESCE(users.reset_sent_at, CONVERT_TZ('0001-01-01 00:00:00','+00:00','UTC')) AS reset_sent_at, COALESCE(users.weibo, '') AS weibo, COALESCE(users.weixin_openid, '') AS weixin_openid, COALESCE(users.avatar, '') AS avatar, users.id, users.created_at, users.updated_at FROM users")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return users, nil
}

// UserCount get the count of all the User records.
func UserCount() (c int64, err error) {
	err = DB.Get(&c, "SELECT count(*) FROM users")
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return c, nil
}

// UserCountWhere get the count of all the User records with a where clause.
func UserCountWhere(where string, args ...interface{}) (c int64, err error) {
	sql := "SELECT count(*) FROM users"
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

// UserIncludesWhere get the User associated models records, currently it's not same as the corresponding "includes" function but "preload" instead in Ruby on Rails. It means that the "sql" should be restricted on User model.
func UserIncludesWhere(assocs []string, sql string, args ...interface{}) (_users []User, err error) {
	_users, err = FindUsersWhere(sql, args...)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if len(assocs) == 0 {
		log.Println("No associated fields ard specified")
		return _users, err
	}
	if len(_users) <= 0 {
		return nil, errors.New("No results available")
	}
	ids := make([]interface{}, len(_users))
	for _, v := range _users {
		ids = append(ids, interface{}(v.Id))
	}
	idsHolder := strings.Repeat(",?", len(ids)-1)
	for _, assoc := range assocs {
		switch assoc {
				case "posts":
							where := fmt.Sprintf("user_id IN (?%s)", idsHolder)
						_posts, err := FindPostsWhere(where, ids...)
						if err != nil {
							log.Printf("Error when query associated objects: %v\n", assoc)
							continue
						}
						for _, vv := range _posts {
							for i, vvv := range  _users {
									if vv.UserId == vvv.Id {
										vvv.Posts = append(vvv.Posts, vv)
									}
								_users[i].Posts = vvv.Posts
						    }
					    }
		}
	}
	return _users, nil
}

// UserIds get all the IDs of User records.
func UserIds() (ids []int64, err error) {
	err = DB.Select(&ids, "SELECT id FROM users")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return ids, nil
}

// UserIdsWhere get all the IDs of User records by where restriction.
func UserIdsWhere(where string, args ...interface{}) ([]int64, error) {
	ids, err := UserIntCol("id", where, args...)
	return ids, err
}

// UserIntCol get some int64 typed column of User by where restriction.
func UserIntCol(col, where string, args ...interface{}) (intColRecs []int64, err error) {
	sql := "SELECT " + col + " FROM users"
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

// UserStrCol get some string typed column of User by where restriction.
func UserStrCol(col, where string, args ...interface{}) (strColRecs []string, err error) {
	sql := "SELECT " + col + " FROM users"
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

// FindUsersWhere query use a partial SQL clause that usually following after WHERE
// with placeholders, eg: FindUsersWhere("first_name = ? AND age > ?", "John", 18)
// will return those records in the table "users" whose first_name is "John" and age elder than 18.
func FindUsersWhere(where string, args ...interface{}) (users []User, err error) {
	sql := "SELECT COALESCE(users.name, '') AS name, COALESCE(users.email, '') AS email, COALESCE(users.password_digest, '') AS password_digest, COALESCE(users.remember_digest, '') AS remember_digest, COALESCE(users.admin, FALSE) AS admin, COALESCE(users.activation_digest, '') AS activation_digest, COALESCE(users.activated, FALSE) AS activated, COALESCE(users.activated_at, CONVERT_TZ('0001-01-01 00:00:00','+00:00','UTC')) AS activated_at, COALESCE(users.reset_digest, '') AS reset_digest, COALESCE(users.reset_sent_at, CONVERT_TZ('0001-01-01 00:00:00','+00:00','UTC')) AS reset_sent_at, COALESCE(users.weibo, '') AS weibo, COALESCE(users.weixin_openid, '') AS weixin_openid, COALESCE(users.avatar, '') AS avatar, users.id, users.created_at, users.updated_at FROM users"
	if len(where) > 0 {
		sql = sql + " WHERE " + where
	}
	stmt, err := DB.Preparex(DB.Rebind(sql))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	err = stmt.Select(&users, args...)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return users, nil
}

// FindUserBySql query use a complete SQL clause
// with placeholders, eg: FindUserBySql("SELECT * FROM users WHERE first_name = ? AND age > ? ORDER BY DESC LIMIT 1", "John", 18)
// will return only One record in the table "users" whose first_name is "John" and age elder than 18.
func FindUserBySql(sql string, args ...interface{}) (*User, error) {
	stmt, err := DB.Preparex(DB.Rebind(sql))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	_user := &User{}
	err = stmt.Get(_user, args...)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return _user, nil
}

// FindUsersBySql query use a complete SQL clause
// with placeholders, eg: FindUsersBySql("SELECT * FROM users WHERE first_name = ? AND age > ?", "John", 18)
// will return those records in the table "users" whose first_name is "John" and age elder than 18.
func FindUsersBySql(sql string, args ...interface{}) (users []User, err error) {
	stmt, err := DB.Preparex(DB.Rebind(sql))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	err = stmt.Select(&users, args...)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return users, nil
}

// CreateUser use a named params to create a single User record.
// A named params is key-value map like map[string]interface{}{"first_name": "John", "age": 23} .
func CreateUser(am map[string]interface{}) (int64, error) {
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
	sqlFmt := `INSERT INTO users (%s) VALUES (%s)`
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

// Create is a method for User to create a record.
func (_user *User) Create() (int64, error) {
	ok, err := govalidator.ValidateStruct(_user)
	if !ok {
		errMsg := "Validate User struct error: Unknown error"
		if err != nil {
			errMsg = "Validate User struct error: " + err.Error()
		}
		log.Println(errMsg)
		return 0, errors.New(errMsg)
	}
	t := time.Now()
	_user.CreatedAt = t
	_user.UpdatedAt = t
    sql := `INSERT INTO users (name,email,created_at,updated_at,password_digest,remember_digest,admin,activation_digest,activated,activated_at,reset_digest,reset_sent_at,weibo,weixin_openid,avatar) VALUES (:name,:email,:created_at,:updated_at,:password_digest,:remember_digest,:admin,:activation_digest,:activated,:activated_at,:reset_digest,:reset_sent_at,:weibo,:weixin_openid,:avatar)`
    result, err := DB.NamedExec(sql, _user)
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

// PostsCreate is used for User to create the associated objects Posts
func (_user *User) PostsCreate(am map[string]interface{}) error {
			am["user_id"] = _user.Id
		_, err := CreatePost(am)
	return err
}

// GetPosts is used for User to get associated objects Posts
// Say you have a User object named user, when you call user.GetPosts(),
// the object will get the associated Posts attributes evaluated in the struct.
func (_user *User) GetPosts() error {
	_posts, err := UserGetPosts(_user.Id)
	if err == nil {
		_user.Posts = _posts
    }
    return err
}

// UserGetPosts a helper fuction used to get associated objects for UserIncludesWhere().
func UserGetPosts(id int64) ([]Post, error) {
			_posts, err := FindPostsBy("user_id", id)
	return _posts, err
}




// Destroy is method used for a User object to be destroyed.
func (_user *User) Destroy() error {
	if _user.Id == 0 {
		return errors.New("Invalid Id field: it can't be a zero value")
	}
	err := DestroyUser(_user.Id)
	return err
}

// DestroyUser will destroy a User record specified by the id parameter.
func DestroyUser(id int64) error {
	// Destroy association objects at first
	// Not care if exec properly temporarily
	destroyUserAssociations(id)
	stmt, err := DB.Preparex(DB.Rebind(`DELETE FROM users WHERE id = ?`))
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

// DestroyUsers will destroy User records those specified by the ids parameters.
func DestroyUsers(ids ...int64) (int64, error) {
	if len(ids) == 0 {
		msg := "At least one or more ids needed"
		log.Println(msg)
		return 0, errors.New(msg)
	}
	// Destroy association objects at first
	// Not care if exec properly temporarily
	destroyUserAssociations(ids...)
	idsHolder := strings.Repeat(",?", len(ids)-1)
	sql := fmt.Sprintf(`DELETE FROM users WHERE id IN (?%s)`, idsHolder)
	idsT := []interface{}{}
	for _,id := range ids {
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

// DestroyUsersWhere delete records by a where clause restriction.
// e.g. DestroyUsersWhere("name = ?", "John")
// And this func will not call the association dependent action
func DestroyUsersWhere(where string, args ...interface{}) (int64, error) {
	sql := `DELETE FROM users WHERE `
	if len(where) > 0 {
		sql = sql + where
	} else {
		return 0, errors.New("No WHERE conditions provided")
	}
	ids, x_err := UserIdsWhere(where, args...)
	if x_err != nil {
		log.Printf("Delete associated objects error: %v\n", x_err)
	} else {
		destroyUserAssociations(ids...)
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

// destroyUserAssociations is a private function used to destroy a User record's associated objects.
// The func not return err temporarily.
func destroyUserAssociations(ids ...int64) {
	idsHolder := ""
	if len(ids) > 1 {
		idsHolder = strings.Repeat(",?", len(ids)-1)
	}
	idsT := []interface{}{}
	for _, id := range ids {
		idsT = append(idsT, interface{}(id))
	}
	var err error
	// make sure no declared-and-not-used exception
	_, _, _ = idsHolder, idsT, err
								where := fmt.Sprintf("user_id IN (?%s)", idsHolder)
							_, err = DestroyPostsWhere(where, idsT...)
							if err != nil {
								log.Printf("Destroy associated object %s error: %v\n", "Posts", err)
							}
}

// Save method is used for a User object to update an existed record mainly.
// If no id provided a new record will be created. FIXME: A UPSERT action will be implemented further.
func (_user *User) Save() error {
	ok, err := govalidator.ValidateStruct(_user)
	if !ok {
		errMsg := "Validate User struct error: Unknown error"
		if err != nil {
			errMsg = "Validate User struct error: " + err.Error()
		}
		log.Println(errMsg)
		return errors.New(errMsg)
	}
	if _user.Id == 0 {
		_, err = _user.Create()
		return err
	}
	_user.UpdatedAt = time.Now()
	sqlFmt := `UPDATE users SET %s WHERE id = %v`
	sqlStr := fmt.Sprintf(sqlFmt, "name = :name, email = :email, updated_at = :updated_at, password_digest = :password_digest, remember_digest = :remember_digest, admin = :admin, activation_digest = :activation_digest, activated = :activated, activated_at = :activated_at, reset_digest = :reset_digest, reset_sent_at = :reset_sent_at, weibo = :weibo, weixin_openid = :weixin_openid, avatar = :avatar", _user.Id)
    _, err = DB.NamedExec(sqlStr, _user)
    return err
}

// UpdateUser is used to update a record with a id and map[string]interface{} typed key-value parameters.
func UpdateUser(id int64, am map[string]interface{}) error {
	if len(am) == 0 {
		return errors.New("Zero key in the attributes map!")
	}
	am["updated_at"] = time.Now()
	keys := allKeys(am)
	sqlFmt := `UPDATE users SET %s WHERE id = %v`
	setKeysArr := []string{}
	for _,v := range keys {
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

// Update is a method used to update a User record with the map[string]interface{} typed key-value parameters.
func (_user *User) Update(am map[string]interface{}) error {
	if _user.Id == 0 {
		return errors.New("Invalid Id field: it can't be a zero value")
	}
	err := UpdateUser(_user.Id, am)
	return err
}

// UpdateAttributes method is supposed to be used to update User records as corresponding update_attributes in Ruby on Rails.
func (_user *User) UpdateAttributes(am map[string]interface{}) error {
	if _user.Id == 0 {
		return errors.New("Invalid Id field: it can't be a zero value")
	}
	err := UpdateUser(_user.Id, am)
	return err
}

// UpdateColumns method is supposed to be used to update User records as corresponding update_columns in Ruby on Rails.
func (_user *User) UpdateColumns(am map[string]interface{}) error {
	if _user.Id == 0 {
		return errors.New("Invalid Id field: it can't be a zero value")
	}
	err := UpdateUser(_user.Id, am)
	return err
}

// UpdateUsersBySql is used to update User records by a SQL clause
// using the '?' binding syntax.
func UpdateUsersBySql(sql string, args ...interface{}) (int64, error) {
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
