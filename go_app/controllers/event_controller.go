package controllers

import (
	"fmt"
	"log"
	m "main/models"
	mm "main/src/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	// you can import models
)

func EventHandler(c *gin.Context) {
	// you can use model functions to do CRUD
	//
	// user, _ := m.FindUser(1)
	// u, err := json.Marshal(user)
	// if err != nil {
	// 	log.Printf("JSON encoding error: %v\n", err)
	// 	u = []byte("Get data error!")
	// }

	// GET /posts
	/* 		posts, err := m.LastPosts(10)
	   		if err != nil {
	   			msg := fmt.Sprintf("Get post index error: %v", err)
	   			c.JSON(http.StatusOK, BuildResp("400", msg, nil))
	   			return
	   		}
	   		resp := BuildResp("200", "Get post index success", posts)
	   		c.JSON(http.StatusOK, resp) */

	lastId, _ := strconv.ParseInt(c.Query("last_id"), 10, 64)
	fmt.Print(lastId)
	pp := &m.EventPage{
		Order:   map[string]string{"id": "desc"},
		LastId:  lastId,
		PerPage: 20,
	}

	direction := "current"
	if lastId > 0 {
		direction = "next"
	}

	ps, err := pp.GetPage(direction)
	fmt.Print(err)

	c.JSON(200, gin.H{
		"status": "success",
		"data":   ps,
	})
}

// FetchAllEvent 获取所有的event
func FetchAllEvent(c *gin.Context) {
	var events []mm.Event
	lastID, _ := strconv.ParseInt(c.Query("last_id"), 10, 64)
	perPage, _ := strconv.ParseInt(c.Query("per_page"), 10, 64)
	forSelect, _ := strconv.ParseBool(c.Query("for_select"))
	if perPage == 0 {
		perPage = 8
	}
	query := mm.DB
	if !forSelect {
		query = query.Preload("Photos").Order("created_at desc").Limit(perPage)
	} else {
		query = query.Order("created_at desc")
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

func CreateEvent(c *gin.Context) {
	var po m.Event
	err := c.BindJSON(&po)
	if err != nil {
		c.JSON(http.StatusOK, err)
		return
	}
	log.Println(po.Photos)
	id, err := po.Create()
	if err != nil {
		msg := fmt.Sprintf("Create post error: %v", err)
		log.Println(msg)
		c.JSON(http.StatusOK, BuildResp("400", msg, nil))
		return
	}
	c.JSON(http.StatusOK, BuildResp("200", "Create post success", map[string]int64{"id": id}))
}

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

// PUT /posts/1
func UpdateEvent(c *gin.Context) {
	id, err := ToInt(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, BuildResp("400", "Parsing id error!", nil))
		return
	}
	post, err := m.FindEvent(id)
	if err != nil {
		msg := fmt.Sprintf("Update post error: %v", err)
		log.Println(msg)
		c.JSON(http.StatusOK, BuildResp("400", msg, nil))
	}
	update_param := map[string]interface{}{}
	var json m.Event
	if c.BindJSON(&json) == nil {
		if json.Content != "" {
			update_param["content"] = json.Content
		}
	}
	err = post.Update(update_param)
	if err != nil {
		msg := fmt.Sprintf("Update post error: %v", err)
		log.Println(msg)
		c.JSON(http.StatusOK, BuildResp("400", msg, nil))
		return
	}
	resp := BuildResp("200", "Update post success", nil)
	c.JSON(http.StatusOK, resp)
}

// DELETE /posts/1
func DestroyEvent(c *gin.Context) {
	id, err := ToInt(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, BuildResp("400", "Params error!", nil))
		return
	}
	err = m.DestroyEvent(id)
	if err != nil {
		fmt.Println(err)
	}
	resp := BuildResp("200", "Article destroied", nil)
	c.JSON(http.StatusOK, resp)
}
