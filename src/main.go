package main

import (
	"github.com/gin-gonic/gin"
	"github.com/validatedid/trussihealth-api/src/api"
	"github.com/validatedid/trussihealth-api/src/packages/config"
)

func main() {
	r := gin.Default()

	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"trussihealth": config.PASSWORD,
	}))

	api.PostHealthDataController(authorized)
	api.GetHealthDataController(authorized)
	r.Run()
}
