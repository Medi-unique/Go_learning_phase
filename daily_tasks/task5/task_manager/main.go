package main

import (
	"example.com/task_manager/data"
	"example.com/task_manager/router"

	"github.com/gin-gonic/gin"
)

type Trainer struct {
	Name string
	Age  int
	City string
}

func main() {
	
	data.StartMongoDB()
	r := gin.Default()
	r = router.Router(r)
	r.Run()

}
