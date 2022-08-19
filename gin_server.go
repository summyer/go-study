package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//https://blog.csdn.net/qq_43756091/article/details/119855735
var logger = log.Default()

func main() {
	engine := gin.Default()
	engine.GET("/get", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	engine.POST("/post", func(context *gin.Context) {
		body, _ := ioutil.ReadAll(context.Request.Body)
		log.Printf("body:%s", body)
		context.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})

	})

	engine.POST("/postForm", func(context *gin.Context) {
		name := context.PostForm("name")
		logger.Printf("name:%s", name)
		context.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	engine.HEAD("/head", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	engine.POST("/complex", func(context *gin.Context) {
		// 从url获取拼接参数
		env := context.Query("env")
		logger.Printf("env:%s", env)

		// 获取请求头参数
		header := context.Request.Header
		logger.Printf("Cookie:%s", header.Get("Cookie"))
		logger.Printf("Content-Type:%s", header.Get("Content-Type"))
		logger.Printf("Token:%s", header.Get("Token"))

		// 获取请求体参数
		body, _ := ioutil.ReadAll(context.Request.Body)
		log.Printf("body:%s", body)
		context.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	engine.Run(":8080")
}
