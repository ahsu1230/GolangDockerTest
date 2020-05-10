package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ahsu1230/GolangDockerTest/src/database"
	"github.com/ahsu1230/GolangDockerTest/src/math"
	"github.com/ahsu1230/GolangDockerTest/src/parsers"
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

	fmt.Println("Parsing configurations...")
	config := parsers.ParserYaml("configurations.yml")
	dbc := config.Database

	fmt.Println("Connecting to database...")
	db, err := database.OpenMySql(dbc.Host, dbc.Port, dbc.Username, dbc.Password, dbc.DbName)
	if err != nil || db == nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("Database connected!")

	dbIsHealthy := false
	for i := 0; i < 10; i++ {
		fmt.Println("Waiting to ping...")
		ctx, stop := context.WithCancel(context.Background())
		defer stop()

		fmt.Println("Pinging...")
		err := database.Ping(db, ctx)
		if err != nil {
			fmt.Println("Database is not ready yet. Sleeping...", i)
			time.Sleep(time.Second * 5)
		} else {
			fmt.Println("Database is healthy!")
			dbIsHealthy = true
			break
		}
	}

	if !dbIsHealthy {
		log.Fatal("Database is not healthy. Abort running webserver.")
	}

	engine := initGin()
	engine.Run()
}
