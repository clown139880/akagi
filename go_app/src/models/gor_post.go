// Package models includes the functions on the model Post.
package models

import (
	"log"
	time "main/src/time"
	utils "main/src/utils"
	"os"
	"regexp"
	"strings"

	"github.com/jinzhu/gorm"
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

// AfterFind 为图片提供缩略图
func (p *Post) AfterFind() (err error) {
	for _, photo := range p.Photos {
		p.Content = strings.ReplaceAll(p.Content, photo.OriginURL, photo.URL)
	}
	return
}

// AfterCreate 在创建之后处理图片
func (p *Post) AfterCreate(scope *gorm.Scope) (err error) {
	scope.DB().Model(p).Update("content", p.Content)
	return
}

// BeforeSave 在保存前处理图片
func (p *Post) BeforeSave() (err error) {
	p.Content = strings.ReplaceAll(p.Content, "!small", "")
	if p.ID > 0 {
		handlePhotos(p)
	}
	return
}

// AfterSave 在保存前处理图片
func (p *Post) AfterSave(tx *gorm.DB) (err error) {
	tx.Model(&p.Event).Where("id = ?", p.EventID).UpdateColumn("updated_at", time.Now())
	return
}

func handlePhotos(p *Post) {
	re, _ := regexp.Compile(`!\[([^\]])*\]\(([^\)])+\)`)
	reURL, _ := regexp.Compile(`http[^\)]+`)
	for _, photoMarkDown := range re.FindAllString(p.Content, -1) {
		photoURL := reURL.FindString(photoMarkDown)
		//print(photoURL)
		key := strings.Replace(photoURL, "https://akagi.oss-cn-hangzhou.aliyuncs.com/", "", 1)
		//print(key)
		var photo Photo
		DB.Where("`key` = ?", key).Find(&photo)
		if photo.ID == 0 {
			//没有找到photos里面的记录
			if strings.Contains(photoURL, "akagi.oss-cn-hangzhou.aliyuncs.com") {
				photo.Key = key
				photo.OriginURL = photoURL
				photo.PhotoableType = "Post"
				photo.PhotoableID = p.ID
				DB.Save(&photo)
			} else {
				photo.Key = utils.UploadFromURL(photoURL)
				photo.PhotoableType = "Post"
				photo.OriginURL = os.Getenv("END_POINT") + photo.Key
				photo.PhotoableID = p.ID
				DB.Save(&photo)
				p.Content = strings.ReplaceAll(p.Content, photoURL, photo.OriginURL)
			}
		}
	}
}
