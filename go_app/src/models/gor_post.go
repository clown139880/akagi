// Package models includes the functions on the model Post.
package models

import (
	"log"
	time "main/src/time"
)

// set flags to output more detailed log
func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// Post 定义原始的数据库字段
type Post struct {
	ID         int64          `json:"id,omitempty" db:"id" valid:"-"`
	VisitCount int64          `json:"visit_count,omitempty" db:"visit_count" valid:"-"`
	Content    string         `json:"content,omitempty" db:"content" valid:"required"`
	UserID     int64          `json:"user_id,omitempty" db:"user_id" valid:"-"`
	CreatedAt  time.TimeStamp `json:"created_at,omitempty" db:"created_at" valid:"-"`
	UpdatedAt  time.TimeStamp `json:"updated_at,omitempty" db:"updated_at" valid:"-"`
	EventID    int64          `json:"event_id,omitempty" db:"event_id" valid:"-"`
	Nickname   string         `json:"nickname,omitempty" db:"nickname" valid:"-"`
	Types      int64          `json:"types,omitempty" db:"types" valid:"-"`
	Title      string         `json:"title,omitempty" db:"title" valid:"-"`
	Status     int64          `json:"status,omitempty" db:"status" valid:"-"`
	IsSynced   bool           `json:"is_synced,omitempty" db:"is_synced" valid:"-"`
	User       *User          `json:"user,omitempty" db:"user" valid:"-"`
	Event      *Event         `json:"event,omitempty" db:"event" valid:"-"`
	Photos     []Photo        `gorm:"auto_preload;polymorphic:Photoable;polymorphic_value:Post" json:"photos,omitempty" db:"photos" valid:"-"`
}
