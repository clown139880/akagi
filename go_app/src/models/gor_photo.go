// Package models includes the functions on the model Photo.
package models

import (
	"log"
	time "main/src/time"
	"strings"
)

// set flags to output more detailed log
func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// Photo 图片
type Photo struct {
	ID            int64          `json:"id,omitempty" db:"id" valid:"-"`
	Key           string         `json:"key,omitempty" db:"key" valid:"-"`
	IsLogo        bool           `json:"is_logo,omitempty" db:"is_logo" valid:"-"`
	URL           string         `json:"url,omitempty" db:"url" valid:"-"`
	PhotoableID   int64          `json:"photoable_id,omitempty" db:"photoable_id" valid:"-"`
	PhotoableType string         `json:"photoable_type,omitempty" db:"photoable_type" valid:"-"`
	CreatedAt     time.TimeStamp `json:"created_at,omitempty" db:"created_at" valid:"-"`
	UpdatedAt     time.TimeStamp `json:"updated_at,omitempty" db:"updated_at" valid:"-"`
}

// AfterFind 为图片提供缩略图
func (p *Photo) AfterFind() (err error) {
	if strings.Contains(p.URL, "oss") {
		p.URL += "!small"
	}
	return
}
