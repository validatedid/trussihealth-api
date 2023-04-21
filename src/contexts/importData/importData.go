package importData

import (
	"encoding/json"

	"github.com/validatedid/trussihealth-api/src/packages/config"
	"github.com/validatedid/trussihealth-api/src/packages/restClient"
	"github.com/validatedid/trussihealth-api/src/packages/vidchain"

	"github.com/validatedid/trussihealth-api/src/packages/cryptography"
	"github.com/validatedid/trussihealth-api/src/packages/dataTransformer"
	"github.com/validatedid/trussihealth-api/src/packages/ipfs"
)

type HealthDataRequest struct {
	Data interface{}
	Did  string
}

type ImportData struct {
	httpClient restClient.HTTPClient
}

func NewImportData(client restClient.HTTPClient) (a *ImportData) {
	return &ImportData{httpClient: client}
}

func (id *ImportData) Execute(healthDataRequest HealthDataRequest) {
	healthDataDetails := dataTransformer.DataTransformer{}.Extract(string(healthDataRequest.Data.([]byte)))
	json, _ := json.Marshal(healthDataDetails)
	basicCryptographer := cryptography.BasicCryptographer{}
	hash := basicCryptographer.Hash(string(json))
	encryptedData := basicCryptographer.Encrypt(string(json), []byte(config.ENCRYPTION_KEY))

	ipfsStorageRepository := ipfs.NewStorageRepository(id.httpClient)
	ipfsIdentifier := ipfsStorageRepository.Save(encryptedData)

	vidchainApiAuthenticator := vidchain.NewVidchainApiAuthenticator(id.httpClient)

	credential := vidchain.NewCredential(id.httpClient, vidchainApiAuthenticator)
	vcPayload := vidchain.VcPayload{DocumentId: ipfsIdentifier, Hash: hash}
	vc := credential.CreateVc(vcPayload)

	eidas := vidchain.NewEidas(id.httpClient, vidchainApiAuthenticator)
	eidas.EsealVc(vc)
}
