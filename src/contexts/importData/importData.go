package importData

import (
	"encoding/json"

	"github.com/validatedid/trussihealth-api/src/packages/config"
	"github.com/validatedid/trussihealth-api/src/packages/issuer"
	"github.com/validatedid/trussihealth-api/src/packages/restClient"

	"github.com/validatedid/trussihealth-api/src/packages/cryptography"
	"github.com/validatedid/trussihealth-api/src/packages/ipfs"
)

type HealthDataRequest struct {
	Data interface{} `json:"data"`
	Did  string      `json:"did"`
}

type ImportData struct {
	ipfsStorageRepository *ipfs.IpfsStorageRepository
	credential            *issuer.Credential
	eidas                 *issuer.Eidas
}

func NewImportData(client restClient.HTTPClient, ipfsClient ipfs.IPFSClient) (i *ImportData) {
	ipfs := ipfs.NewStorageRepository(ipfsClient)
	credentialCreator := issuer.NewCredential(client)
	eidasSealer := issuer.NewEidas(client)
	return &ImportData{ipfsStorageRepository: ipfs, credential: credentialCreator, eidas: eidasSealer}
}

func (i *ImportData) Execute(healthDataRequest HealthDataRequest) {
	hash, encryptedData := encryptData(healthDataRequest)

	ipfsIdentifier := i.ipfsStorageRepository.Save(encryptedData)

	vc := i.credential.CreateVc(issuer.VcPayload{DocumentId: ipfsIdentifier, Hash: hash}, healthDataRequest.Did)

	i.eidas.EsealVc(vc)
}

func encryptData(healthDataRequest HealthDataRequest) (string, string) {
	data, _ := json.Marshal(healthDataRequest.Data)
	basicCryptographer := cryptography.BasicCryptographer{}
	hash := basicCryptographer.Hash(string(data))
	encryptedData := basicCryptographer.Encrypt(string(data), []byte(config.ENCRYPTION_KEY))
	return hash, encryptedData
}
