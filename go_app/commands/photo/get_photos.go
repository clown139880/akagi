// Package command includes the functions on the model Event.
package main

import (
	"fmt"
	m "main/src/models"
	"os"
	"regexp"
	"strings"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	var photos []m.Photo
	m.DB.Where("photoable_type = ?", "Event").Find(&photos)
	for _, photo := range photos {
		if !strings.Contains(photo.OriginURL, "akagi") && photo.Key != "" {
			photo.OriginURL = os.Getenv("END_POINT") + photo.Key
			m.DB.Save(&photo)
		}
	}

	//uploadFromURL("https://ooo.0o0.ooo/2016/10/19/5807ba71104d9.jpg")
}

func handlePosts() {
	var posts []m.Post
	m.DB.Find(&posts)
	re, _ := regexp.Compile(`!\[([^\]])+\]\(([^\)])+\)`)
	reURL, _ := regexp.Compile(`http[^\)]+`)
	//fileObj, _ := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	for _, post := range posts {
		for _, photoMarkDown := range re.FindAllString(post.Content, -1) {
			photoURL := reURL.FindString(photoMarkDown)
			//print(photoURL)
			key := strings.Replace(photoURL, "https://akagi.oss-cn-hangzhou.aliyuncs.com/", "", 1)
			//print(key)
			var photo m.Photo
			m.DB.Where("`key` = ?", key).Find(&photo)
			if photo.ID == 0 {
				//没有找到photos里面的记录
				if strings.Contains(photoURL, "akagi.oss-cn-hangzhou.aliyuncs.com") {
					photo.Key = key
					photo.URL = photoURL
					photo.PhotoableType = "Post"
					photo.PhotoableID = post.ID
					m.DB.Save(&photo)
				} else {
					fmt.Printf("%v\r\n", photoURL)
					photo.Key = uploadFromURL(photoURL)
					photo.PhotoableType = "Post"
					photo.PhotoableID = post.ID
					m.DB.Save(&photo)
					post.Content = strings.ReplaceAll(post.Content, photoURL, os.Getenv("END_POINT")+photo.Key)
					m.DB.Save(&post)
				}
			}
		}
		m.DB.Save(&post)
	}

	//uploadFromURL("https://ooo.0o0.ooo/2016/10/19/5807ba71104d9.jpg")
}
