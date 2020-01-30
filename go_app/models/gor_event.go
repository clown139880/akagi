// Package models includes the functions on the model Event.
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

type Event struct {
	Id        int64     `json:"id,omitempty" db:"id" valid:"-"`
	Title     string    `json:"title,omitempty" db:"title" valid:"-"`
	Content   string    `json:"content,omitempty" db:"content" valid:"-"`
	UserId    int64     `json:"user_id,omitempty" db:"user_id" valid:"-"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at" valid:"-"`
	UpdatedAt time.Time `json:"updated_at,omitempty" db:"updated_at" valid:"-"`
	Level     int64     `json:"level,omitempty" db:"level" valid:"-"`
	ParentId  int64     `json:"parent_id,omitempty" db:"parent_id" valid:"-"`
	Nickname  string    `json:"nickname,omitempty" db:"nickname" valid:"-"`
	Types     int64     `json:"types,omitempty" db:"types" valid:"-"`
	StartedAt time.Time `json:"started_at,omitempty" db:"started_at" valid:"-"`
	EndedAt   time.Time `json:"ended_at,omitempty" db:"ended_at" valid:"-"`
	Status    int64     `json:"status,omitempty" db:"status" valid:"-"`
	Place     string    `json:"place,omitempty" db:"place" valid:"-"`
	Photos    []Photo   `json:"photos,omitempty" db:"photos" valid:"-"`
	Posts     []Post    `json:"posts,omitempty" db:"posts" valid:"-"`
	User      User      `json:"user,omitempty" db:"user" valid:"-"`
}

// DataStruct for the pagination
type EventPage struct {
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

// Current get the current page of EventPage object for pagination.
func (_p *EventPage) Current() ([]Event, error) {
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
	events, err := EventIncludesWhere(append([]string{}, "photos"), whereStr, whereParams...)
	if err != nil {
		return nil, err
	}
	if len(events) != 0 {
		_p.FirstId, _p.LastId = events[0].Id, events[len(events)-1].Id
	}
	return events, nil
}

// Previous get the previous page of EventPage object for pagination.
func (_p *EventPage) Previous() ([]Event, error) {
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
	events, err := EventIncludesWhere(append([]string{}, "photos"), whereStr, whereParams...)
	if err != nil {
		return nil, err
	}
	if len(events) != 0 {
		_p.FirstId, _p.LastId = events[0].Id, events[len(events)-1].Id
	}
	_p.PageNum--
	return events, nil
}

// Next get the next page of EventPage object for pagination.
func (_p *EventPage) Next() ([]Event, error) {
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
	events, err := EventIncludesWhere(append([]string{}, "photos"), whereStr, whereParams...)
	if err != nil {
		return nil, err
	}
	if len(events) != 0 {
		_p.FirstId, _p.LastId = events[0].Id, events[len(events)-1].Id
	}
	_p.PageNum++
	return events, nil
}

// GetPage is a helper function for the EventPage object to return a corresponding page due to
// the parameter passed in, i.e. one of "previous, current or next".
func (_p *EventPage) GetPage(direction string) (ps []Event, err error) {
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

// buildOrder is for EventPage object to build a SQL ORDER BY clause.
func (_p *EventPage) buildOrder() {
	tempList := []string{}
	for k, v := range _p.Order {
		tempList = append(tempList, fmt.Sprintf("%v %v", k, v))
	}
	_p.orderStr = " ORDER BY " + strings.Join(tempList, ", ")
}

// buildIdRestrict is for EventPage object to build a SQL clause for ID restriction,
// implementing a simple keyset style pagination.
func (_p *EventPage) buildIdRestrict(direction string) (idStr string, idParams []interface{}) {
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

// buildPageCount calculate the TotalItems/TotalPages for the EventPage object.
func (_p *EventPage) buildPageCount() error {
	count, err := EventCountWhere(_p.WhereString, _p.WhereParams...)
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

// FindEvent find a single event by an ID.
func FindEvent(id int64) (*Event, error) {
	if id == 0 {
		return nil, errors.New("Invalid ID: it can't be zero")
	}
	_event := Event{}
	err := DB.Get(&_event, DB.Rebind(`SELECT COALESCE(events.title, '') AS title, COALESCE(events.content, '') AS content, COALESCE(events.user_id, 0) AS user_id, COALESCE(events.level, 0) AS level, COALESCE(events.parent_id, 0) AS parent_id, COALESCE(events.nickname, '') AS nickname, COALESCE(events.types, 0) AS types, COALESCE(events.started_at, CONVERT_TZ('0001-01-01 00:00:00','+00:00','UTC')) AS started_at, COALESCE(events.ended_at, CONVERT_TZ('0001-01-01 00:00:00','+00:00','UTC')) AS ended_at, COALESCE(events.status, 0) AS status, COALESCE(events.place, '') AS place, events.id, events.created_at, events.updated_at FROM events WHERE events.id = ? LIMIT 1`), id)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return nil, err
	}
	return &_event, nil
}

// FirstEvent find the first one event by ID ASC order.
func FirstEvent() (*Event, error) {
	_event := Event{}
	err := DB.Get(&_event, DB.Rebind(`SELECT COALESCE(events.title, '') AS title, COALESCE(events.content, '') AS content, COALESCE(events.user_id, 0) AS user_id, COALESCE(events.level, 0) AS level, COALESCE(events.parent_id, 0) AS parent_id, COALESCE(events.nickname, '') AS nickname, COALESCE(events.types, 0) AS types, COALESCE(events.started_at, CONVERT_TZ('0001-01-01 00:00:00','+00:00','UTC')) AS started_at, COALESCE(events.ended_at, CONVERT_TZ('0001-01-01 00:00:00','+00:00','UTC')) AS ended_at, COALESCE(events.status, 0) AS status, COALESCE(events.place, '') AS place, events.id, events.created_at, events.updated_at FROM events ORDER BY events.id ASC LIMIT 1`))
	if err != nil {
		log.Printf("Error: %v\n", err)
		return nil, err
	}
	return &_event, nil
}

// FirstEvents find the first N events by ID ASC order.
func FirstEvents(n uint32) ([]Event, error) {
	_events := []Event{}
	sql := fmt.Sprintf("SELECT COALESCE(events.title, '') AS title, COALESCE(events.content, '') AS content, COALESCE(events.user_id, 0) AS user_id, COALESCE(events.level, 0) AS level, COALESCE(events.parent_id, 0) AS parent_id, COALESCE(events.nickname, '') AS nickname, COALESCE(events.types, 0) AS types, COALESCE(events.started_at, CONVERT_TZ('0001-01-01 00:00:00','+00:00','UTC')) AS started_at, COALESCE(events.ended_at, CONVERT_TZ('0001-01-01 00:00:00','+00:00','UTC')) AS ended_at, COALESCE(events.status, 0) AS status, COALESCE(events.place, '') AS place, events.id, events.created_at, events.updated_at FROM events ORDER BY events.id ASC LIMIT %v", n)
	err := DB.Select(&_events, DB.Rebind(sql))
	if err != nil {
		log.Printf("Error: %v\n", err)
		return nil, err
	}
	return _events, nil
}

// LastEvent find the last one event by ID DESC order.
func LastEvent() (*Event, error) {
	_event := Event{}
	err := DB.Get(&_event, DB.Rebind(`SELECT COALESCE(events.title, '') AS title, COALESCE(events.content, '') AS content, COALESCE(events.user_id, 0) AS user_id, COALESCE(events.level, 0) AS level, COALESCE(events.parent_id, 0) AS parent_id, COALESCE(events.nickname, '') AS nickname, COALESCE(events.types, 0) AS types, COALESCE(events.started_at, CONVERT_TZ('0001-01-01 00:00:00','+00:00','UTC')) AS started_at, COALESCE(events.ended_at, CONVERT_TZ('0001-01-01 00:00:00','+00:00','UTC')) AS ended_at, COALESCE(events.status, 0) AS status, COALESCE(events.place, '') AS place, events.id, events.created_at, events.updated_at FROM events ORDER BY events.id DESC LIMIT 1`))
	if err != nil {
		log.Printf("Error: %v\n", err)
		return nil, err
	}
	return &_event, nil
}

// LastEvents find the last N events by ID DESC order.
func LastEvents(n uint32) ([]Event, error) {
	_events := []Event{}
	sql := fmt.Sprintf("SELECT COALESCE(events.title, '') AS title, COALESCE(events.content, '') AS content, COALESCE(events.user_id, 0) AS user_id, COALESCE(events.level, 0) AS level, COALESCE(events.parent_id, 0) AS parent_id, COALESCE(events.nickname, '') AS nickname, COALESCE(events.types, 0) AS types, COALESCE(events.started_at, CONVERT_TZ('0001-01-01 00:00:00','+00:00','UTC')) AS started_at, COALESCE(events.ended_at, CONVERT_TZ('0001-01-01 00:00:00','+00:00','UTC')) AS ended_at, COALESCE(events.status, 0) AS status, COALESCE(events.place, '') AS place, events.id, events.created_at, events.updated_at FROM events ORDER BY events.id DESC LIMIT %v", n)
	err := DB.Select(&_events, DB.Rebind(sql))
	if err != nil {
		log.Printf("Error: %v\n", err)
		return nil, err
	}
	return _events, nil
}

// FindEvents find one or more events by the given ID(s).
func FindEvents(ids ...int64) ([]Event, error) {
	if len(ids) == 0 {
		msg := "At least one or more ids needed"
		log.Println(msg)
		return nil, errors.New(msg)
	}
	_events := []Event{}
	idsHolder := strings.Repeat(",?", len(ids)-1)
	sql := DB.Rebind(fmt.Sprintf(`SELECT COALESCE(events.title, '') AS title, COALESCE(events.content, '') AS content, COALESCE(events.user_id, 0) AS user_id, COALESCE(events.level, 0) AS level, COALESCE(events.parent_id, 0) AS parent_id, COALESCE(events.nickname, '') AS nickname, COALESCE(events.types, 0) AS types, COALESCE(events.started_at, CONVERT_TZ('0001-01-01 00:00:00','+00:00','UTC')) AS started_at, COALESCE(events.ended_at, CONVERT_TZ('0001-01-01 00:00:00','+00:00','UTC')) AS ended_at, COALESCE(events.status, 0) AS status, COALESCE(events.place, '') AS place, events.id, events.created_at, events.updated_at FROM events WHERE events.id IN (?%s)`, idsHolder))
	idsT := []interface{}{}
	for _, id := range ids {
		idsT = append(idsT, interface{}(id))
	}
	err := DB.Select(&_events, sql, idsT...)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return nil, err
	}
	return _events, nil
}

// FindEventBy find a single event by a field name and a value.
func FindEventBy(field string, val interface{}) (*Event, error) {
	_event := Event{}
	sqlFmt := `SELECT COALESCE(events.title, '') AS title, COALESCE(events.content, '') AS content, COALESCE(events.user_id, 0) AS user_id, COALESCE(events.level, 0) AS level, COALESCE(events.parent_id, 0) AS parent_id, COALESCE(events.nickname, '') AS nickname, COALESCE(events.types, 0) AS types, COALESCE(events.started_at, CONVERT_TZ('0001-01-01 00:00:00','+00:00','UTC')) AS started_at, COALESCE(events.ended_at, CONVERT_TZ('0001-01-01 00:00:00','+00:00','UTC')) AS ended_at, COALESCE(events.status, 0) AS status, COALESCE(events.place, '') AS place, events.id, events.created_at, events.updated_at FROM events WHERE %s = ? LIMIT 1`
	sqlStr := fmt.Sprintf(sqlFmt, field)
	err := DB.Get(&_event, DB.Rebind(sqlStr), val)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return nil, err
	}
	return &_event, nil
}

// FindEventsBy find all events by a field name and a value.
func FindEventsBy(field string, val interface{}) (_events []Event, err error) {
	sqlFmt := `SELECT COALESCE(events.title, '') AS title, COALESCE(events.content, '') AS content, COALESCE(events.user_id, 0) AS user_id, COALESCE(events.level, 0) AS level, COALESCE(events.parent_id, 0) AS parent_id, COALESCE(events.nickname, '') AS nickname, COALESCE(events.types, 0) AS types, COALESCE(events.started_at, CONVERT_TZ('0001-01-01 00:00:00','+00:00','UTC')) AS started_at, COALESCE(events.ended_at, CONVERT_TZ('0001-01-01 00:00:00','+00:00','UTC')) AS ended_at, COALESCE(events.status, 0) AS status, COALESCE(events.place, '') AS place, events.id, events.created_at, events.updated_at FROM events WHERE %s = ?`
	sqlStr := fmt.Sprintf(sqlFmt, field)
	err = DB.Select(&_events, DB.Rebind(sqlStr), val)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return nil, err
	}
	return _events, nil
}

// AllEvents get all the Event records.
func AllEvents() (events []Event, err error) {
	err = DB.Select(&events, "SELECT COALESCE(events.title, '') AS title, events.id, events.created_at, events.updated_at FROM events")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return events, nil
}

// EventCount get the count of all the Event records.
func EventCount() (c int64, err error) {
	err = DB.Get(&c, "SELECT count(*) FROM events")
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return c, nil
}

// EventCountWhere get the count of all the Event records with a where clause.
func EventCountWhere(where string, args ...interface{}) (c int64, err error) {
	sql := "SELECT count(*) FROM events"
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

// EventIncludesWhere get the Event associated models records, currently it's not same as the corresponding "includes" function but "preload" instead in Ruby on Rails. It means that the "sql" should be restricted on Event model.
func EventIncludesWhere(assocs []string, sql string, args ...interface{}) (_events []Event, err error) {
	_events, err = FindEventsWhere(sql, args...)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if len(assocs) == 0 {
		log.Println("No associated fields ard specified")
		return _events, err
	}
	if len(_events) <= 0 {
		return nil, errors.New("No results available")
	}
	ids := make([]interface{}, len(_events))
	for _, v := range _events {
		ids = append(ids, interface{}(v.Id))
	}
	idsHolder := strings.Repeat(",?", len(ids)-1)
	for _, assoc := range assocs {
		switch assoc {
		case "photos":
			// FIXME: optimize the query
			for i, vvv := range _events {
				_photos, err := EventGetPhotos(vvv.Id)
				if err != nil {
					continue
				}
				vvv.Photos = _photos
				_events[i] = vvv
			}
		case "posts":
			where := fmt.Sprintf("event_id IN (?%s)", idsHolder)
			_posts, err := FindPostsWhere(where, ids...)
			if err != nil {
				log.Printf("Error when query associated objects: %v\n", assoc)
				continue
			}
			for _, vv := range _posts {
				for i, vvv := range _events {
					if vv.EventId == vvv.Id {
						vvv.Posts = append(vvv.Posts, vv)
					}
					_events[i].Posts = vvv.Posts
				}
			}
		}
	}
	return _events, nil
}

// EventIds get all the IDs of Event records.
func EventIds() (ids []int64, err error) {
	err = DB.Select(&ids, "SELECT id FROM events")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return ids, nil
}

// EventIdsWhere get all the IDs of Event records by where restriction.
func EventIdsWhere(where string, args ...interface{}) ([]int64, error) {
	ids, err := EventIntCol("id", where, args...)
	return ids, err
}

// EventIntCol get some int64 typed column of Event by where restriction.
func EventIntCol(col, where string, args ...interface{}) (intColRecs []int64, err error) {
	sql := "SELECT " + col + " FROM events"
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

// EventStrCol get some string typed column of Event by where restriction.
func EventStrCol(col, where string, args ...interface{}) (strColRecs []string, err error) {
	sql := "SELECT " + col + " FROM events"
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

// FindEventsWhere query use a partial SQL clause that usually following after WHERE
// with placeholders, eg: FindUsersWhere("first_name = ? AND age > ?", "John", 18)
// will return those records in the table "users" whose first_name is "John" and age elder than 18.
func FindEventsWhere(where string, args ...interface{}) (events []Event, err error) {
	sql := "SELECT COALESCE(events.title, '') AS title, COALESCE(events.content, '') AS content, COALESCE(events.user_id, 0) AS user_id, COALESCE(events.level, 0) AS level, COALESCE(events.parent_id, 0) AS parent_id, COALESCE(events.nickname, '') AS nickname, COALESCE(events.types, 0) AS types, COALESCE(events.started_at, CONVERT_TZ('0001-01-01 00:00:00','+00:00','UTC')) AS started_at, COALESCE(events.ended_at, CONVERT_TZ('0001-01-01 00:00:00','+00:00','UTC')) AS ended_at, COALESCE(events.status, 0) AS status, COALESCE(events.place, '') AS place, events.id, events.created_at, events.updated_at FROM events"
	if len(where) > 0 {
		sql = sql + " WHERE " + where
	}
	stmt, err := DB.Preparex(DB.Rebind(sql))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	err = stmt.Select(&events, args...)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return events, nil
}

// FindEventBySql query use a complete SQL clause
// with placeholders, eg: FindUserBySql("SELECT * FROM users WHERE first_name = ? AND age > ? ORDER BY DESC LIMIT 1", "John", 18)
// will return only One record in the table "users" whose first_name is "John" and age elder than 18.
func FindEventBySql(sql string, args ...interface{}) (*Event, error) {
	stmt, err := DB.Preparex(DB.Rebind(sql))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	_event := &Event{}
	err = stmt.Get(_event, args...)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return _event, nil
}

// FindEventsBySql query use a complete SQL clause
// with placeholders, eg: FindUsersBySql("SELECT * FROM users WHERE first_name = ? AND age > ?", "John", 18)
// will return those records in the table "users" whose first_name is "John" and age elder than 18.
func FindEventsBySql(sql string, args ...interface{}) (events []Event, err error) {
	stmt, err := DB.Preparex(DB.Rebind(sql))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	err = stmt.Select(&events, args...)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return events, nil
}

// CreateEvent use a named params to create a single Event record.
// A named params is key-value map like map[string]interface{}{"first_name": "John", "age": 23} .
func CreateEvent(am map[string]interface{}) (int64, error) {
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
	sqlFmt := `INSERT INTO events (%s) VALUES (%s)`
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

// Create is a method for Event to create a record.
func (_event *Event) Create() (int64, error) {
	ok, err := govalidator.ValidateStruct(_event)
	if !ok {
		errMsg := "Validate Event struct error: Unknown error"
		if err != nil {
			errMsg = "Validate Event struct error: " + err.Error()
		}
		log.Println(errMsg)
		return 0, errors.New(errMsg)
	}
	t := time.Now()
	_event.CreatedAt = t
	_event.UpdatedAt = t
	sql := `INSERT INTO events (title,content,user_id,created_at,updated_at,level,parent_id,nickname,types,started_at,ended_at,status,place) VALUES (:title,:content,:user_id,:created_at,:updated_at,:level,:parent_id,:nickname,:types,:started_at,:ended_at,:status,:place)`
	result, err := DB.NamedExec(sql, _event)
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

// PhotosCreate is used for Event to create the associated objects Photos
func (_event *Event) PhotosCreate(am map[string]interface{}) error {
	am["photoable_id"] = _event.Id
	am["photoable_type"] = "Event"
	_, err := CreatePhoto(am)
	return err
}

// GetPhotos is used for Event to get associated objects Photos
// Say you have a Event object named event, when you call event.GetPhotos(),
// the object will get the associated Photos attributes evaluated in the struct.
func (_event *Event) GetPhotos() error {
	_photos, err := EventGetPhotos(_event.Id)
	if err == nil {
		_event.Photos = _photos
	}
	return err
}

// EventGetPhotos a helper fuction used to get associated objects for EventIncludesWhere().
func EventGetPhotos(id int64) ([]Photo, error) {
	where := `photoable_type = "Event" AND photoable_id = ?`
	_photos, err := FindPhotosWhere(where, id)
	return _photos, err
}

// PostsCreate is used for Event to create the associated objects Posts
func (_event *Event) PostsCreate(am map[string]interface{}) error {
	am["event_id"] = _event.Id
	_, err := CreatePost(am)
	return err
}

// GetPosts is used for Event to get associated objects Posts
// Say you have a Event object named event, when you call event.GetPosts(),
// the object will get the associated Posts attributes evaluated in the struct.
func (_event *Event) GetPosts() error {
	_posts, err := EventGetPosts(_event.Id)
	if err == nil {
		_event.Posts = _posts
	}
	return err
}

// EventGetPosts a helper fuction used to get associated objects for EventIncludesWhere().
func EventGetPosts(id int64) ([]Post, error) {
	_posts, err := FindPostsBy("event_id", id)
	return _posts, err
}

// CreateUser is a method for a Event object to create an associated User record.
func (_event *Event) CreateUser(am map[string]interface{}) error {
	am["event_id"] = _event.Id
	_, err := CreateUser(am)
	return err
}

// Destroy is method used for a Event object to be destroyed.
func (_event *Event) Destroy() error {
	if _event.Id == 0 {
		return errors.New("Invalid Id field: it can't be a zero value")
	}
	err := DestroyEvent(_event.Id)
	return err
}

// DestroyEvent will destroy a Event record specified by the id parameter.
func DestroyEvent(id int64) error {
	// Destroy association objects at first
	// Not care if exec properly temporarily
	destroyEventAssociations(id)
	stmt, err := DB.Preparex(DB.Rebind(`DELETE FROM events WHERE id = ?`))
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

// DestroyEvents will destroy Event records those specified by the ids parameters.
func DestroyEvents(ids ...int64) (int64, error) {
	if len(ids) == 0 {
		msg := "At least one or more ids needed"
		log.Println(msg)
		return 0, errors.New(msg)
	}
	// Destroy association objects at first
	// Not care if exec properly temporarily
	destroyEventAssociations(ids...)
	idsHolder := strings.Repeat(",?", len(ids)-1)
	sql := fmt.Sprintf(`DELETE FROM events WHERE id IN (?%s)`, idsHolder)
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

// DestroyEventsWhere delete records by a where clause restriction.
// e.g. DestroyEventsWhere("name = ?", "John")
// And this func will not call the association dependent action
func DestroyEventsWhere(where string, args ...interface{}) (int64, error) {
	sql := `DELETE FROM events WHERE `
	if len(where) > 0 {
		sql = sql + where
	} else {
		return 0, errors.New("No WHERE conditions provided")
	}
	ids, x_err := EventIdsWhere(where, args...)
	if x_err != nil {
		log.Printf("Delete associated objects error: %v\n", x_err)
	} else {
		destroyEventAssociations(ids...)
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

// destroyEventAssociations is a private function used to destroy a Event record's associated objects.
// The func not return err temporarily.
func destroyEventAssociations(ids ...int64) {
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
	where := fmt.Sprintf("event_id IN (?%s)", idsHolder)
	_, err = DestroyPostsWhere(where, idsT...)
	if err != nil {
		log.Printf("Destroy associated object %s error: %v\n", "Posts", err)
	}
}

// Save method is used for a Event object to update an existed record mainly.
// If no id provided a new record will be created. FIXME: A UPSERT action will be implemented further.
func (_event *Event) Save() error {
	ok, err := govalidator.ValidateStruct(_event)
	if !ok {
		errMsg := "Validate Event struct error: Unknown error"
		if err != nil {
			errMsg = "Validate Event struct error: " + err.Error()
		}
		log.Println(errMsg)
		return errors.New(errMsg)
	}
	if _event.Id == 0 {
		_, err = _event.Create()
		return err
	}
	_event.UpdatedAt = time.Now()
	sqlFmt := `UPDATE events SET %s WHERE id = %v`
	sqlStr := fmt.Sprintf(sqlFmt, "title = :title, content = :content, user_id = :user_id, updated_at = :updated_at, level = :level, parent_id = :parent_id, nickname = :nickname, types = :types, started_at = :started_at, ended_at = :ended_at, status = :status, place = :place", _event.Id)
	_, err = DB.NamedExec(sqlStr, _event)
	return err
}

// UpdateEvent is used to update a record with a id and map[string]interface{} typed key-value parameters.
func UpdateEvent(id int64, am map[string]interface{}) error {
	if len(am) == 0 {
		return errors.New("Zero key in the attributes map!")
	}
	am["updated_at"] = time.Now()
	keys := allKeys(am)
	sqlFmt := `UPDATE events SET %s WHERE id = %v`
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

// Update is a method used to update a Event record with the map[string]interface{} typed key-value parameters.
func (_event *Event) Update(am map[string]interface{}) error {
	if _event.Id == 0 {
		return errors.New("Invalid Id field: it can't be a zero value")
	}
	err := UpdateEvent(_event.Id, am)
	return err
}

// UpdateAttributes method is supposed to be used to update Event records as corresponding update_attributes in Ruby on Rails.
func (_event *Event) UpdateAttributes(am map[string]interface{}) error {
	if _event.Id == 0 {
		return errors.New("Invalid Id field: it can't be a zero value")
	}
	err := UpdateEvent(_event.Id, am)
	return err
}

// UpdateColumns method is supposed to be used to update Event records as corresponding update_columns in Ruby on Rails.
func (_event *Event) UpdateColumns(am map[string]interface{}) error {
	if _event.Id == 0 {
		return errors.New("Invalid Id field: it can't be a zero value")
	}
	err := UpdateEvent(_event.Id, am)
	return err
}

// UpdateEventsBySql is used to update Event records by a SQL clause
// using the '?' binding syntax.
func UpdateEventsBySql(sql string, args ...interface{}) (int64, error) {
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
