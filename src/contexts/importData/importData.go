package importData

import (
	"net/http"

	"github.com/validatedid/trussihealth-api/src/packages/cryptography"
	"github.com/validatedid/trussihealth-api/src/packages/dataTransformer"
	"github.com/validatedid/trussihealth-api/src/packages/ipfs"
)

type ImportData struct{}

func (id ImportData) Execute(inJsonData string) {

	healthData := dataTransformer.DataTransformer{}.Extract(inJsonData)

	basicCryptographer := cryptography.BasicCryptographer{}
	basicCryptographer.Hash("healthData")
	encryptedData := basicCryptographer.Encrypt("healthData", []byte("thisis32bitlongpassphraseimusing"))

	ipfsStorageRepository := ipfs.NewStorageRepository(http.DefaultClient)
	ipfsIdentifier := ipfsStorageRepository.Save(encryptedData)
}
