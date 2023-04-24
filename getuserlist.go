package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skytsai2/database"
)

func GetUserList(c *gin.Context) {
	res := database.GetUserList()
	c.AsciiJSON(http.StatusOK, gin.H{
		"status":    "000",
		"errorCode": "00",
		"data":      res,
	})
}
