package main

import (
	"log"

	"example.com/task_manager/data"
	"example.com/task_manager/router"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Trainer struct {
	Name string
	Age  int
	City string
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	data.StartMongoDB()
	r := gin.Default()
	r = router.Router(r)
	r.Run()

}
