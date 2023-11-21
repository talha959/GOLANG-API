package main

import (
	"GIN/api"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/upload", api.HandleFileUpload)
	err := r.Run(":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
