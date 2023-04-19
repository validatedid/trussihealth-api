package ipfs

import (
	"bytes"
	"io"
	"net/http"

	"github.com/validatedid/trussihealth-api/src/packages/restClient"
)

type StorageRepository interface {
	Save(data string) (id string)
	GetById(id string) (data string)
}

type IpfsStorageRepository struct {
	httpClient restClient.HTTPClient
}

func NewStorageRepository(client restClient.HTTPClient) (i IpfsStorageRepository) {
	return IpfsStorageRepository{httpClient: client}
}

func (i IpfsStorageRepository) Save(data string) (id string) {
	request, _ := http.NewRequest("POST", "https://url", bytes.NewBufferString(data))
	response, _ := i.httpClient.Do(request)
	body, _ := io.ReadAll(response.Body)
	return string(body)
}
