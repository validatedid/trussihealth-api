package main

import (
	"github.com/gin-gonic/gin"
	"github.com/validatedid/trussihealth-api/src/api"
)

func main() {
	r := gin.Default()
	api.PostHealthData(r)
	r.Run() // listen and serve on 0.0.0.0:8080
}
