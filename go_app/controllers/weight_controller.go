package controllers

import (
	m "go_app/src/models"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

// FetchAllWeight 获取所有体重
func FetchAllWeight(c *gin.Context) {
	var weights []m.WeightLog
	m.DB.Order("created_at").Find(&weights)

	if len(weights) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No weights found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": weights})
}

// FetchSingleWeight 方法返回一条 weight 数据
func FetchSingleWeight(c *gin.Context) {
	var weight m.WeightLog
	weightID := c.Param("id")
	log.Print(reflect.TypeOf(weightID))

	if weightID == "0" {
		m.DB.Last(&weight)
	} else {
		m.DB.First(&weight, weightID)
	}
	if weight.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": weight})
}

// CreateWeight 记录体重
func CreateWeight(c *gin.Context) {

	/* 	weight, _ := strconv.ParseFloat(c.PostForm("weight"), 64)
	   	weightLog := m.WeightLog{Weight: weight} */
	var weightLog m.WeightLog
	err := c.BindJSON(&weightLog)
	if err != nil {
		data, _ := ioutil.ReadAll(c.Request.Body)
		log.Printf("c.Request.body: %v", string(data))
		log.Println(err)
		c.JSON(http.StatusOK, err)
		return
	}

	m.DB.Save(&weightLog)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Weight Create successfully"})
}
