package controllers

import (
	"fmt"
	"io/ioutil"
	"log"
	m "main/src/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	// you can import models
)

// FetchAllEvent 获取所有的event
func FetchAllEvent(c *gin.Context) {
	var events []m.Event
	page, _ := strconv.ParseInt(c.Query("page"), 10, 64)
	perPage, _ := strconv.ParseInt(c.Query("per_page"), 10, 64)
	forSelect, _ := strconv.ParseBool(c.Query("for_select"))
	if perPage == 0 {
		perPage = 8
	}
	query := m.DB.Order("updated_at desc")
	if !forSelect {
		query = query.Preload("Photos", func(db *gorm.DB) *gorm.DB {
			return db.Order("is_logo DESC")
		}).Limit(perPage).Offset(page * perPage)
	}

	query.Find(&events)

	if len(events) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No events found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": events})
}

// ShowEvent 获取单个event
func ShowEvent(c *gin.Context) {
	id, err := ToInt(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, BuildResp("400", "Parsing id error!", nil))
		return
	}
	var event m.Event
	m.DB.Preload("Photos").Find(&event, id)
	//m.DB.Model(&event).Related("Photos")
	if event.ID == 0 {
		msg := fmt.Sprintf("Get post error: %v", 404)
		c.JSON(http.StatusOK, BuildResp("404", msg, nil))
		return
	}
	resp := BuildResp("200", "Get post success", event)
	c.JSON(http.StatusOK, resp)
}

// CreateEvent 创建POST
func CreateEvent(c *gin.Context) {
	var event m.Event
	err := c.BindJSON(&event)
	if err != nil {
		data, _ := ioutil.ReadAll(c.Request.Body)
		log.Printf("c.Request.body: %v", string(data))
		log.Println(err)
		c.JSON(http.StatusOK, err)
		return
	}
	log.Println(event.Photos)

	m.DB.Create(&event)
	resp := BuildResp("200", "Create event success", event)
	c.JSON(http.StatusOK, resp)
}

// UpdateEvent PUT /events/1
func UpdateEvent(c *gin.Context) {
	id, err := ToInt(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, BuildResp("400", "Parsing id error!", nil))
		return
	}
	var event m.Event
	m.DB.Find(&event, id)
	if event.ID == 0 {
		msg := fmt.Sprintf("Update event error: %v", err)
		log.Println(msg)
		c.JSON(http.StatusOK, BuildResp("400", msg, nil))
	}
	var json m.Event
	c.BindJSON(&json)

	m.DB.Model(&event).Select([]string{"content", "title", "event_id"}).Updates(json)
	for _, photo := range json.Photos {
		if photo.ID == 0 {
			m.DB.Model(&event).Association("Photos").Append(&photo)
		} else {
			m.DB.Model(&photo).Updates(photo)
		}
	}

	resp := BuildResp("200", "Update event success", event)
	c.JSON(http.StatusOK, resp)
}
