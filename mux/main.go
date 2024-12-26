package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	// 定义路由和处理程序
	router.HandleFunc("/api", yourHandler).Methods("GET", "POST")

	// 配置CORS
	credentials := handlers.AllowCredentials()
	methods := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})
	headers := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	origins := handlers.AllowedOrigins([]string{"https://example.com"}) // 指定特定的域名
	maxAge := handlers.MaxAge(int((24 * time.Hour).Seconds()))          // 设置预检请求的缓存时间

	corsHandler := handlers.CORS(credentials, methods, headers, origins, maxAge)

	// 将CORS中间件应用于整个路由器
	log.Fatal(http.ListenAndServe(":5000", corsHandler(router)))
}

func yourHandler(w http.ResponseWriter, r *http.Request) {
	// 处理请求
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Hello, CORS!"}`))
}
