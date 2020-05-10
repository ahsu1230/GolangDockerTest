package main

import (
	"fmt"

	"github.com/ahsu1230/GolangDockerTest/src/database"
	"github.com/ahsu1230/GolangDockerTest/src/math"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func initGin() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return r
}

func main() {
	fmt.Println("Hello!")
	fmt.Println("Math", math.Add(3, 5))

	db, err := database.OpenMySql()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// ctx, stop := context.WithCancel(context.Background())
	// defer stop()
	// database.Ping(db, ctx)

	engine := initGin()
	engine.Run()
}
