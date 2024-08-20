package main

import (
	"github.com/gin-gonic/gin"

	"example.com/task_manager/data"
	"example.com/task_manager/router"
)



func main() {

	data.StartMongoDB()
	r := gin.Default()
	r = router.Router(r)
	r.Run()

}
