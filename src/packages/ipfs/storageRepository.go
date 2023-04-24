package ipfs

import (
	"bytes"
	"fmt"
	shell "github.com/ipfs/go-ipfs-api"
	"github.com/validatedid/trussihealth-api/src/packages/restClient"
)

type StorageRepository interface {
	Save(data string) (id string)
	GetById(id string) (data string)
}

type IpfsStorageRepository struct {
	httpClient restClient.HTTPClient
}

func NewStorageRepository(client restClient.HTTPClient) (i *IpfsStorageRepository) {
	return &IpfsStorageRepository{httpClient: client}
}

func (i *IpfsStorageRepository) Save(data string) (id string) {
	// create a new IPFS API client
	sh := shell.NewShell("http://52.157.145.27:5001")

	// read file contents into memory
	fileContents := []byte(data)

	// add file to IPFS
	hash, err := sh.Add(bytes.NewReader(fileContents))
	if err != nil {
		panic(err)
	}

	// print IPFS hash of the file
	fmt.Println("File uploaded to IPFS with hash:", hash)

	return hash
}
