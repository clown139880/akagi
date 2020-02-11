// Package models includes the functions on the model Photo.
package models

import (
	"log"
	time "main/src/time"
	"os"
)

// set flags to output more detailed log
func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// Photo 图片
type Photo struct {
	ID            int64          `json:"id,omitempty" db:"id" valid:"-"`
	Key           string         `json:"key" db:"key" valid:"-"`
	IsLogo        bool           `json:"is_logo,omitempty" db:"is_logo" valid:"-"`
	URL           string         `gorm:"-" json:"url,omitempty" valid:"-"`
	OriginURL     string         `gorm:"column:url" json:"-"  valid:"-"`
	PhotoableID   int64          `json:"photoable_id,omitempty" db:"photoable_id" valid:"-"`
	PhotoableType string         `json:"photoable_type,omitempty" db:"photoable_type" valid:"-"`
	CreatedAt     time.TimeStamp `json:"created_at,omitempty" db:"created_at" valid:"-"`
	UpdatedAt     time.TimeStamp `json:"updated_at,omitempty" db:"updated_at" valid:"-"`
}

// AfterFind 为图片提供缩略图
func (p *Photo) AfterFind() (err error) {
	if p.Key != "" {
		p.URL = os.Getenv("END_POINT") + p.Key + "!small"
	} else {
		p.URL = p.OriginURL
	}
	return
}

// BeforeSave 为图片提供缩略图
func (p *Photo) BeforeSave() (err error) {
	if p.Key != "" {
		p.OriginURL = os.Getenv("END_POINT") + p.Key
	} else {
		p.OriginURL = p.URL
	}
	return
}
