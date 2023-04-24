package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skytsai2/database"
)

type AddUserPost struct {
	Name string `json:"name" binding:"required"`
	Tel  string `json:"tel" binding:"required"`
}

func PostUser(c *gin.Context) {

	var addUserPost AddUserPost
	err := c.ShouldBindJSON(&addUserPost)
	if err != nil {
		c.JSON(500, gin.H{
			"errorCode": "01",
			"status":    "503",
			"msg":       err.Error(),
		})
		return
	}

	database.AddUser(addUserPost.Name, addUserPost.Tel)
	// will output : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
	// c.AsciiJSON(http.StatusOK, data)
	c.JSON(http.StatusOK, gin.H{
		"Name": addUserPost.Name,
		"Tel":  addUserPost.Tel,
	})
}
