package storeData_test

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/validatedid/trussihealth-api/src/packages/ipfs"
	"io"
	"net/http"
	"testing"
)

type HttpClientTestDouble struct {
	fakedResponse string
	Request       *http.Request
}

func (td HttpClientTestDouble) Do(req *http.Request) (*http.Response, error) {
	td.Request = *req
	var resp http.Response
	resp.Body = io.NopCloser(bytes.NewBufferString(td.fakedResponse))
	return &resp, nil
}

func TestStoreData(t *testing.T) {
	identifier := "Ipfs identifier is null"
	clientTestDouble := &HttpClientTestDouble{fakedResponse: identifier}
	ipfsStorageRepository := ipfs.NewStorageRepository(clientTestDouble)
	ipfsIdentifier := ipfsStorageRepository.Save("encrypted_data")

	encryptedContentBytes, _ := io.ReadAll(clientTestDouble.Request.Body)
	encryptedContent := string(encryptedContentBytes)
	assert.Equal(t, encryptedContent, "encrypted_data")
	assert.Equal(t, ipfsIdentifier, identifier)
	assert.NotNil(t, ipfsIdentifier, "Ipfs identifier is null")
}
