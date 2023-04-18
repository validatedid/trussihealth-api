package main

import (
	"github.com/gin-gonic/gin"
	"github.com/validatedid/trussihealth-api/contexts/medicalData"
)

func main() {
	r := gin.Default()
	medicalData.Listen(r)
	r.Run() // listen and serve on 0.0.0.0:8080
}
