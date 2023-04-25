package main

import (
	"github.com/gin-gonic/gin"
	"github.com/validatedid/trussihealth-api/src/api"
)

func main() {
	r := gin.Default()
	api.PostHealthDataController(r)
	api.GetHealthDataController(r)
	r.Run()
}
