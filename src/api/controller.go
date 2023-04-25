package api

import (
	"encoding/json"
	shell "github.com/ipfs/go-ipfs-api"
	"github.com/validatedid/trussihealth-api/src/packages/config"
	"github.com/validatedid/trussihealth-api/src/packages/ipfs"
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
		sh := shell.NewShell(config.IPFS_URL)
		ipfsWrapper := ipfs.NewIPFSClientWrapper(sh)
		importData.NewImportData(http.DefaultClient, ipfsWrapper).Execute(healthData)
		c.JSON(http.StatusOK, gin.H{})
	})
}
