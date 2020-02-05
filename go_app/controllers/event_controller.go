package controllers

import (
	"fmt"
	mm "main/src/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	// you can import models
)

// FetchAllEvent 获取所有的event
func FetchAllEvent(c *gin.Context) {
	var events []mm.Event
	lastID, _ := strconv.ParseInt(c.Query("last_id"), 10, 64)
	perPage, _ := strconv.ParseInt(c.Query("per_page"), 10, 64)
	forSelect, _ := strconv.ParseBool(c.Query("for_select"))
	if perPage == 0 {
		perPage = 8
	}
	query := mm.DB.Order("updated_at desc")
	if !forSelect {
		query = query.Preload("Photos").Limit(perPage)
	}
	if lastID > 0 {
		query = query.Where("id < ?", lastID)
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
	var event mm.Event
	mm.DB.Preload("Photos").Find(&event, id)
	//mm.DB.Model(&event).Related("Photos")
	if event.ID == 0 {
		msg := fmt.Sprintf("Get post error: %v", 404)
		c.JSON(http.StatusOK, BuildResp("404", msg, nil))
		return
	}
	resp := BuildResp("200", "Get post success", event)
	c.JSON(http.StatusOK, resp)
}
