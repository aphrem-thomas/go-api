package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

const (
	hello = "hello world..."
	howdi = "howdi partner..."
)

func main() {
	r := gin.Default()
	fmt.Println(hello)
	r.GET("/howdi", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": howdi,
		})
	})
	r.Run()
}
