package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/validatedid/trussihealth-api/src/contexts/importData"
)

func PostHealthDataController(router *gin.Engine) {
	router.POST("/health-data", func(c *gin.Context) {

		jsonRequest, _ := io.ReadAll(c.Request.Body)
		var healthData importData.HealthDataRequest
		json.Unmarshal(jsonRequest, &healthData)
		importData.NewImportData(http.DefaultClient).Execute(healthData)
		c.JSON(http.StatusOK, gin.H{})
	})
}
