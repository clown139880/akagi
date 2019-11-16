package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeHandler(c *gin.Context) {
	// you can use model functions to do CRUD
	//
	// user, _ := m.FindUser(1)
	// u, err := json.Marshal(user)
	// if err != nil {
	// 	log.Printf("JSON encoding error: %v\n", err)
	// 	u = []byte("Get data error!")
	// }

	type Envs struct {
		GoOnRailsVer string
		GolangVer    string
	}

	gorVer := "we changed the title"
	golangVer := "Failed to get"

	envs := Envs{GoOnRailsVer: gorVer, GolangVer: golangVer}
	c.HTML(http.StatusOK, "index.tmpl", envs)
}
