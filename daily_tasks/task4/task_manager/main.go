package main

import (
	

	"example.com/task_manager/router"
	"github.com/gin-gonic/gin"
)



func main() {
	r := gin.Default()
	r = router.SetRoute(r)
	r.Run()

}
