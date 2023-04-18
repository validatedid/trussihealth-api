package medicalData

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Listen(router *gin.Engine) {
	router.GET("/medical-data", func(c *gin.Context) {

		workflow2 := &Workflow2{}

		workflow1 := &Workflow1{}
		workflow1.setNext(workflow2)

		medicalData := &MedicalData{name: "abc", data: "patient data"}

		workflow1.execute(medicalData)

		c.JSON(http.StatusOK, gin.H{
			"result": "ok",
		})
	})
}
