package user

import (
	"github.com/gin-gonic/gin"
	"github.com/skytsai2/database"
	"net/http"
)

func DelUser(c *gin.Context) {
	userid := c.Param("userid")
	database.DelUser(userid)
	c.AsciiJSON(http.StatusOK, gin.H{
		"status":    "000",
		"errorCode": "00",
	})
}
