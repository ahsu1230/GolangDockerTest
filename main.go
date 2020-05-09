package main

import (
	"fmt"

	"github.com/ahsu1230/GolangDockerTest/src/math"
	"github.com/gin-gonic/gin"
)

func initGin() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return r
}

func main() {
	fmt.Println("Hello!")
	fmt.Println("Math", math.Add(3, 5))

	engine := initGin()
	engine.Run()
}
