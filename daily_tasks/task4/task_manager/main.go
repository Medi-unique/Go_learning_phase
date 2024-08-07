package main

import (
	"example.com/task_manager/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r = router.Router(r)
	r.Run()
}
