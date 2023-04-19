package storeData_test

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/validatedid/trussihealth-api/src/packages/ipfs"
)

type HttpClientTestDouble struct {
	fakedResponse string
}

func (td HttpClientTestDouble) Do(req *http.Request) (*http.Response, error) {
	var resp http.Response
	resp.Body = io.NopCloser(bytes.NewBufferString(td.fakedResponse))
	return &resp, nil
}

func TestStoreData(t *testing.T) {
	identifier := "Ipfs identifier is null"
	clientTestDouble := &HttpClientTestDouble{fakedResponse: identifier}

	data := "encrypted_data"

	ipfsStorageRepository := ipfs.NewStorageRepository(clientTestDouble)
	ipfsIdentifier := ipfsStorageRepository.Save(data)

	assert.Equal(t, ipfsIdentifier, identifier)
	assert.NotNil(t, ipfsIdentifier, "Ipfs identifier is null")
}
