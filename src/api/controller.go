package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/validatedid/trussihealth-api/src/contexts/importData"
)

func PostHealthData(router *gin.Engine) {
	router.POST("/health-data", func(c *gin.Context) {

		jsonRequest, _ := json.Marshal(c.Request.Body)
		var healthData importData.HealthDataRequest
		json.Unmarshal(jsonRequest, &healthData)
		importData.NewImportData(http.DefaultClient).Execute(healthData)

		c.JSON(http.StatusOK, gin.H{})
	})
}
