package storeData_test

import (
	"bytes"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/validatedid/trussihealth-api/src/packages/ipfs"
)

type httpClientMock struct {
	mock.Mock
}

func newHttpClientMock() *httpClientMock {
	return &httpClientMock{}
}

func (m *httpClientMock) Do(req *http.Request) (*http.Response, error) {
	args := m.Called(req)
	return args.Get(0).(*http.Response), args.Error(1)
}

func TestStoreData(t *testing.T) {
	identifier := "Ipfs identifier is null"
	encryptedData := "encrypted_data"
	var resp http.Response
	resp.Body = io.NopCloser(bytes.NewBufferString(identifier))

	clientTestDouble := newHttpClientMock()
	clientTestDouble.On("Do", mock.Anything).Return(&resp, nil)
	ipfsStorageRepository := ipfs.NewStorageRepository(clientTestDouble)
	expectedRequest, _ := http.NewRequest("POST", "https://url", bytes.NewBufferString(encryptedData))
	ipfsIdentifier := ipfsStorageRepository.Save(encryptedData)

	assert.Equal(t, ipfsIdentifier, identifier)
	assert.NotNil(t, ipfsIdentifier, "Ipfs identifier is null")
	calledRequest := clientTestDouble.Calls[0].Arguments[0].(*http.Request)
	assert.Equal(t, calledRequest.Body, expectedRequest.Body)
}
