package importData

import (
	"encoding/json"
	"fmt"
	"github.com/validatedid/trussihealth-api/src/packages/restClient"
	"github.com/validatedid/trussihealth-api/src/packages/vidchain"
	"net/http"

	"github.com/validatedid/trussihealth-api/src/packages/cryptography"
	"github.com/validatedid/trussihealth-api/src/packages/dataTransformer"
	"github.com/validatedid/trussihealth-api/src/packages/ipfs"
)

type ImportData struct {
	httpClient restClient.HTTPClient
}

func NewImportData(client restClient.HTTPClient) (a *ImportData) {
	return &ImportData{httpClient: client}
}

func (id *ImportData) Execute(inJsonData string) {
	healthData := dataTransformer.DataTransformer{}.Extract(inJsonData)
	json, _ := json.Marshal(healthData)
	basicCryptographer := cryptography.BasicCryptographer{}
	hash := basicCryptographer.Hash(string(json))
	encryptedData := basicCryptographer.Encrypt("healthData", []byte("thisis32bitlongpassphraseimusing"))

	ipfsStorageRepository := ipfs.NewStorageRepository(http.DefaultClient)
	ipfsIdentifier := ipfsStorageRepository.Save(encryptedData)

	vidchainApiAuthenticator := vidchain.NewVidchainApiAuthenticator(http.DefaultClient)

	credential := vidchain.NewCredential(http.DefaultClient, vidchainApiAuthenticator)
	vcPayload := vidchain.VcPayload{DocumentId: ipfsIdentifier, Hash: hash}
	vc := credential.CreateVc(vcPayload)

	eidas := vidchain.NewEidas(http.DefaultClient, vidchainApiAuthenticator)
	esealed := eidas.EsealVc(vc)
	fmt.Println(string(esealed.Content))
}
