package main

import (
	"fmt"
	"os"
	"strconv"

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

	fmt.Println("Retrieving environment variables...")

	fmt.Println("Connecting to database...")
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	db, err := database.OpenMySql(os.Getenv("DB_HOST"), port, os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_DEFAULT"))
	if err != nil || db == nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("Database connected!")

	engine := initGin()
	engine.Run()
}
