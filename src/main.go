package main

import (
	"github.com/gin-gonic/gin"
	"github.com/validatedid/trussihealth-api/src/api"
)

func main() {
	r := gin.Default()
	api.PostHealthDataController(r)
	r.Run(":3000") // listen and serve on 0.0.0.0:8080
}
