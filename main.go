package main

import (
	"context"
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

	fmt.Println("Connecting to database...")
	db, err := database.OpenMySql()
	if err != nil || db == nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("Database connected!")

	fmt.Println("Waiting to ping...")
	ctx, stop := context.WithCancel(context.Background())
	defer stop()

	fmt.Println("Pinging...")
	database.Ping(db, ctx)

	engine := initGin()
	engine.Run()
}
