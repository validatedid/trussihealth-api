package exportData

import (
	"fmt"
	"github.com/validatedid/trussihealth-api/src/packages/config"
	"github.com/validatedid/trussihealth-api/src/packages/cryptography"
	"github.com/validatedid/trussihealth-api/src/packages/ipfs"
	"io"
)

type ExportHealthDataRequest struct {
	Hash string
}

type ExportHealthDataResponse struct {
	Content string
}

type ExportData struct {
	ipfsClient ipfs.IPFSClient
}

func NewExportData(ipfsClient ipfs.IPFSClient) *ExportData {
	return &ExportData{ipfsClient: ipfsClient}
}

func (e *ExportData) Execute(request ExportHealthDataRequest) ExportHealthDataResponse {
	encryptedData, ipfsErr := e.ipfsClient.Cat(request.Hash)
	if ipfsErr != nil {
		fmt.Println("Error retrieving document from IPFS")
		panic(ipfsErr)
	}
	encryptedBytes, _ := io.ReadAll(encryptedData)
	plaintext, decryptError := cryptography.BasicCryptographer{}.Decrypt(string(encryptedBytes), []byte(config.ENCRYPTION_KEY))
	if decryptError != nil {
		fmt.Println("Error decrypting data")
	}
	return ExportHealthDataResponse{Content: plaintext}
}
