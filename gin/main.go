package main

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// 配置CORS中间件
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://example.com"},                     // 允许的源
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},                  // 允许的HTTP方法
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // 允许的HTTP头
		ExposeHeaders:    []string{"Content-Length"},                          // 暴露的HTTP头
		AllowCredentials: true,                                                // 是否允许携带cookie
		MaxAge:           12 * time.Hour,                                      // 缓存预检请求的时间
	}))

	router.GET("/api", yourHandler)
	router.Run(":5000")
}

func yourHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello, CORS!"})
}
