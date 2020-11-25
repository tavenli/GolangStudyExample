package main

import (
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("------------")

	// 禁用控制台颜色
	gin.DisableConsoleColor()

	// 创建记录日志的文件
	f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)

	// 如果需要将日志同时写入文件和控制台，请使用以下代码
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := gin.Default()
	router.Static("/static", "./static")

	//router.Delims("{.{", "}.}")
	//router.LoadHTMLGlob("./templates/**/*")
	//router.Use(ShowRequestInfo())

	router.GET("/index", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "ok",
		})
	})

	router.Run(":7070")

}
