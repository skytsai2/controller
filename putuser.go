package user

import (
	"github.com/gin-gonic/gin"
	"github.com/skytsai2/database"
	"net/http"
)

type EditUserPost struct {
	ID  int    `json:"id" binding:"required"`
	Tel string `json:"tel" binding:"required"`
}

func PutUser(c *gin.Context) {

	var editUserPost EditUserPost
	err := c.ShouldBindJSON(&editUserPost)
	if err != nil {
		c.JSON(500, gin.H{
			"errorCode": "01",
			"status":    "503",
			"msg":       err.Error(),
		})
		return
	}

	database.EditUser(editUserPost.ID, editUserPost.Tel)
	c.AsciiJSON(http.StatusOK, gin.H{
		"status":    "000",
		"errorCode": "00",
	})
}
