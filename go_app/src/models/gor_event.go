// Package models includes the functions on the model Event.
package models

import (
	time "go_app/src/time"
	"log"
)

// set flags to output more detailed log
func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// Event 事件
type Event struct {
	ID        int64          `json:"id,omitempty" db:"id" valid:"-"`
	Title     string         `json:"title,omitempty" db:"title" valid:"-"`
	Content   string         `json:"content,omitempty" db:"content" valid:"-"`
	UserID    int64          `json:"user_id,omitempty" db:"user_id" valid:"-"`
	CreatedAt time.TimeStamp `json:"created_at,omitempty" db:"created_at" valid:"-"`
	UpdatedAt time.TimeStamp `json:"updated_at,omitempty" db:"updated_at" valid:"-"`
	Level     int64          `json:"level,omitempty" db:"level" valid:"-"`
	ParentID  int64          `json:"parent_id,omitempty" db:"parent_id" valid:"-"`
	Nickname  string         `json:"nickname,omitempty" db:"nickname" valid:"-"`
	Types     int64          `json:"types,omitempty" db:"types" valid:"-"`
	StartedAt time.TimeStamp `json:"started_at,omitempty" db:"started_at" valid:"-"`
	EndedAt   time.TimeStamp `json:"ended_at,omitempty" db:"ended_at" valid:"-"`
	Status    int64          `json:"status,omitempty" db:"status" valid:"-"`
	Place     string         `json:"place,omitempty" db:"place" valid:"-"`
	Posts     []Post         `json:"posts,omitempty" db:"posts" valid:"-"`
	User      *User          `json:"user,omitempty" db:"user" valid:"-"`
	Photos    []Photo        `gorm:"auto_preload;polymorphic:Photoable;polymorphicValue:Event" json:"photos,omitempty" db:"photos"  valid:"-"`
}
