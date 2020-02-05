package controllers

import (
	"fmt"
	"io/ioutil"
	"log"
	m "main/src/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	// you can import models
)

// FetchAllPost 获取所有的post
func FetchAllPost(c *gin.Context) {
	var posts []m.Post
	lastID, _ := strconv.ParseInt(c.Query("last_id"), 10, 64)
	perPage, _ := strconv.ParseInt(c.Query("per_page"), 10, 64)
	eventID, _ := strconv.ParseInt(c.Query("event_id"), 10, 64)
	if perPage == 0 {
		perPage = 8
	}
	query := m.DB.Preload("Photos").Order("created_at desc").Limit(perPage)
	log.Printf("lastID:%v", lastID)
	log.Printf("eventID:%v", eventID)
	if lastID > 0 {
		query = query.Where("id < ?", lastID)
	}
	if eventID > 0 {
		query = query.Where("event_id = ?", eventID)
	}
	query.Find(&posts)

	if len(posts) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No posts found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": posts})
}

// CreatePost 创建POST
func CreatePost(c *gin.Context) {
	var post m.Post
	err := c.BindJSON(&post)
	if err != nil {
		data, _ := ioutil.ReadAll(c.Request.Body)
		log.Printf("c.Request.body: %v", string(data))
		log.Println(err)
		c.JSON(http.StatusOK, err)
		return
	}
	log.Println(post.Photos)

	m.DB.Save(&post)
	c.JSON(http.StatusOK, BuildResp("200", "Create post success", map[string]int64{"id": post.ID}))
}

// ShowPost 获取单个POST
func ShowPost(c *gin.Context) {
	id, err := ToInt(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, BuildResp("400", "Parsing id error!", nil))
		return
	}
	var post m.Post
	m.DB.Preload("Photos").Find(&post, id)
	if post.ID == 0 {
		msg := fmt.Sprintf("Get post error: %v", err)
		c.JSON(http.StatusOK, BuildResp("400", msg, nil))
		return
	}
	resp := BuildResp("200", "Get post success", post)
	c.JSON(http.StatusOK, resp)
}

// UpdatePost PUT /posts/1
func UpdatePost(c *gin.Context) {
	id, err := ToInt(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, BuildResp("400", "Parsing id error!", nil))
		return
	}
	var post m.Post
	m.DB.Find(&post, id)
	if post.ID == 0 {
		msg := fmt.Sprintf("Update post error: %v", err)
		log.Println(msg)
		c.JSON(http.StatusOK, BuildResp("400", msg, nil))
	}
	var json m.Post
	c.BindJSON(&json)

	m.DB.Model(&post).Select([]string{"content", "title", "event_id"}).Updates(json)
	resp := BuildResp("200", "Update post success", nil)
	c.JSON(http.StatusOK, resp)
}

// DestroyPost DELETE /posts/1
func DestroyPost(c *gin.Context) {
	id, err := ToInt(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, BuildResp("400", "Parsing id error!", nil))
		return
	}
	var post m.Post
	m.DB.Find(&post, id)
	m.DB.Delete(&post)
	resp := BuildResp("200", "Post destroied", nil)
	c.JSON(http.StatusOK, resp)
}
