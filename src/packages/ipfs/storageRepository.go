package ipfs

import (
	"bytes"
	"fmt"
	"io"
)

type StorageRepository interface {
	Save(data string) (id string)
	GetById(id string) (data string)
}

type IPFSClient interface {
	Add(data *bytes.Reader) (hash string, error error)
	Cat(path string) (io.ReadCloser, error)
}

type IpfsStorageRepository struct {
	ipfsClient IPFSClient
}

func NewStorageRepository(client IPFSClient) (i *IpfsStorageRepository) {
	// create a new IPFS API client
	// sh := shell.NewShell("http://52.157.145.27:5001")
	return &IpfsStorageRepository{ipfsClient: client}
}

func (i *IpfsStorageRepository) Save(data string) (id string) {

	// read file contents into memory
	fileContents := []byte(data)

	// add file to IPFS
	hash, err := i.ipfsClient.Add(bytes.NewReader(fileContents))
	if err != nil {
		panic(err)
	}

	// print IPFS hash of the file
	fmt.Println("File uploaded to IPFS with hash:", hash)

	return hash
}

func (i *IpfsStorageRepository) GetById(id string) string {
	data, err := i.ipfsClient.Cat(id)
	if err != nil {
		panic(err)
	}
	stringData, _ := io.ReadAll(data)
	// print IPFS hash of the file
	fmt.Println("File retrieved hash:", id)
	fmt.Println("File retrieved:", string(stringData))
	return string(stringData)
}
