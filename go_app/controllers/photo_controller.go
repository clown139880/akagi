package controllers

import (
	m "go_app/src/models"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

// FetchAllPhoto 获取所有体重
func FetchAllPhoto(c *gin.Context) {
	var photoss []m.Photo
	m.DB.Order("created_at").Find(&photoss)

	if len(photoss) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No photoss found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": photoss})
}

// FetchSinglePhoto 方法返回一条 photos 数据
func FetchSinglePhoto(c *gin.Context) {
	var photos m.Photo
	photosID := c.Param("id")
	log.Print(reflect.TypeOf(photosID))

	if photosID == "0" {
		m.DB.Last(&photos)
	} else {
		m.DB.First(&photos, photosID)
	}
	if photos.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": photos})
}

// CreatePhoto 记录体重
func CreatePhoto(c *gin.Context) {

	/* 	photos, _ := strconv.ParseFloat(c.PostForm("photos"), 64)
	   	photosLog := m.Photo{Photo: photos} */
	var photosLog m.Photo
	err := c.BindJSON(&photosLog)
	if err != nil {
		data, _ := ioutil.ReadAll(c.Request.Body)
		log.Printf("c.Request.body: %v", string(data))
		log.Println(err)
		c.JSON(http.StatusOK, err)
		return
	}

	m.DB.Save(&photosLog)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Photo Create successfully"})
}
