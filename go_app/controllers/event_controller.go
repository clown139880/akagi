package controllers

import (
	"fmt"
	"log"
	m "main/models"
	"net/http"

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

	events, err := m.AllEvents()
	if err != nil {
		msg := fmt.Sprintf("Get post index error: %v", err)
		c.JSON(http.StatusOK, BuildResp("400", msg, nil))
		return
	}

	c.JSON(200, gin.H{
		"status": "success",
		"data":   events,
	})
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
	post, err := m.FindEvent(id)
	if err != nil {
		msg := fmt.Sprintf("Get post error: %v", err)
		c.JSON(http.StatusOK, BuildResp("400", msg, nil))
		return
	}
	post.GetPhotos()
	resp := BuildResp("200", "Get post success", post)
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
