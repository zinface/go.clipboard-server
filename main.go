package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ClipBoard struct {
	Data     string    `json:"data,omitempty"`
	Mime     string    `json:"mime,omitempty"`
	CreateAt time.Time `json:"create_at"`
}

//	去除 Data 字段以降低数据量，返回其余基本信息
func stripData(c ClipBoard) ClipBoard {
	c.Data = ""
	return c
}

func main() {
	var currentBoard ClipBoard

	r := gin.Default()

	//	GET /clipboard - 获取当前剪贴板
	r.GET("/clipboard", func(c *gin.Context) {
		c.JSON(http.StatusOK, currentBoard)
	})

	//	POST /clipboard - 设置当前剪贴板
	r.POST("/clipboard", func(c *gin.Context) {
		var board ClipBoard
		if err := c.BindJSON(&board); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		currentBoard = board
		currentBoard.CreateAt = time.Now()
		//	返回剪贴板基本信息
		c.JSON(http.StatusOK, stripData(currentBoard))
	})

	//	GET /clipboard/info - 获取剪贴板基本信息
	r.GET("/clipboard/info", func(c *gin.Context) {
		c.JSON(http.StatusOK, stripData(currentBoard))
	})

	r.Run("0.0.0.0:9090")
}
