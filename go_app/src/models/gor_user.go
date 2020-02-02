// Package models includes the functions on the model User.
package models

import (
	"log"
	time "main/src/time"
)

// set flags to output more detailed log
func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// User 用户
type User struct {
	ID               int64          `json:"id,omitempty" db:"id" valid:"-"`
	Name             string         `json:"name,omitempty" db:"name" valid:"-"`
	Email            string         `json:"email,omitempty" db:"email" valid:"-"`
	CreatedAt        time.TimeStamp `json:"created_at,omitempty" db:"created_at" valid:"-"`
	UpdatedAt        time.TimeStamp `json:"updated_at,omitempty" db:"updated_at" valid:"-"`
	PasswordDigest   string         `json:"password_digest,omitempty" db:"password_digest" valid:"-"`
	RememberDigest   string         `json:"remember_digest,omitempty" db:"remember_digest" valid:"-"`
	Admin            bool           `json:"admin,omitempty" db:"admin" valid:"-"`
	ActivationDigest string         `json:"activation_digest,omitempty" db:"activation_digest" valid:"-"`
	Activated        bool           `json:"activated,omitempty" db:"activated" valid:"-"`
	ActivatedAt      time.TimeStamp `json:"activated_at,omitempty" db:"activated_at" valid:"-"`
	ResetDigest      string         `json:"reset_digest,omitempty" db:"reset_digest" valid:"-"`
	ResetSentAt      time.TimeStamp `json:"reset_sent_at,omitempty" db:"reset_sent_at" valid:"-"`
	Weibo            string         `json:"weibo,omitempty" db:"weibo" valid:"-"`
	WeixinOpenid     string         `json:"weixin_openid,omitempty" db:"weixin_openid" valid:"-"`
	Avatar           string         `json:"avatar,omitempty" db:"avatar" valid:"-"`
	Posts            []Post         `json:"posts,omitempty" db:"posts" valid:"-"`
}
