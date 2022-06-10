package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"gitee.com/zinface/go.qclipboard-server/models"
)

var r *gin.Engine

var currentBoard = models.ClipBoard{}

func init() {
	r = gin.Default()
}

func main() {

	r.GET("/check", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"createAt": currentBoard.CreateAt,
		})
	})

	r.POST("/set", func(ctx *gin.Context) {

		var cb = models.ClipBoard{}
		var data = ctx.PostForm("data")
		if strings.Compare(cb.BaseData, data) != 0 {
			cb.BaseData = ctx.PostForm("data")
			cb.MimeType = ctx.PostForm("mime")
			cb.CreateAt = time.Now()
			currentBoard = cb
		}
		ctx.JSON(http.StatusOK, gin.H{})
	})

	r.GET("/get", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, currentBoard)
	})

	r.Run("0.0.0.0:9090")
}
