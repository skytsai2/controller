package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skytsai2/database"
)

func GetUser(c *gin.Context) {
	userid := c.Param("userid")
	res := database.GetUser(userid)
	c.AsciiJSON(http.StatusOK, gin.H{
		"id":   res.ID,
		"name": res.Name,
		"tel":  res.Tel,
	})
}
