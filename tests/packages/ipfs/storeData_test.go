package storeData_test

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/validatedid/trussihealth-api/src/packages/ipfs"
	"io"
	"strings"
	"testing"
)

type IpfsClientMock struct {
	mock.Mock
}

func newIpfsClientMock() *IpfsClientMock {
	return &IpfsClientMock{}
}

func (m *IpfsClientMock) Add(data *bytes.Reader) (string, error) {
	args := m.Called(data)
	return args.String(0), args.Error(1)
}

func (m *IpfsClientMock) Cat(path string) (io.ReadCloser, error) {
	args := m.Called(path)
	return args.Get(0).(io.ReadCloser), args.Error(1)
}

func TestStoreData(t *testing.T) {
	encryptedData := "encrypted_data"
	responseHash := "QmbK8c7LhqauRbQdoNiX4aEHtp64cs9ypeQnCPBCJqCx3a"

	clientTestDouble := newIpfsClientMock()
	clientTestDouble.On("Add", mock.Anything).Return(responseHash, nil)
	ipfsStorageRepository := ipfs.NewStorageRepository(clientTestDouble)

	expectedRequest := bytes.NewReader([]byte(encryptedData))
	ipfsIdentifier := ipfsStorageRepository.Save(encryptedData)

	assert.Equal(t, ipfsIdentifier, responseHash)
	assert.NotNil(t, ipfsIdentifier, "Ipfs identifier is null")
	calledRequest := clientTestDouble.Calls[0].Arguments[0].(*bytes.Reader)
	assert.Equal(t, calledRequest, expectedRequest)
}

func TestGetById(t *testing.T) {
	clientTestDouble := newIpfsClientMock()
	ipfsStorageRepository := ipfs.NewStorageRepository(clientTestDouble)
	documentMock := "fakeData"
	readCloser := io.NopCloser(strings.NewReader(documentMock))

	clientTestDouble.On("Cat", mock.Anything).Return(readCloser, nil)
	data := ipfsStorageRepository.GetById("QmVHKK8MwmB6FfffTywF7giespBej7eW7i4x7y8683ZbAENhj")
	assert.Equal(t, documentMock, data)
}

// To run this test it is required adding your IP to the whitelist here:
// https://portal.azure.com/#@TestViDSocial.onmicrosoft.com/resource/subscriptions/d3585c03-8653-4ad7-972a-9b3ca65874b9/resourceGroups/rg-vidchain-dev/providers/Microsoft.Compute/virtualMachines/vidchain-ipfs-node/networking
// Remember to remove your IP after running this integration test
/*
func TestGetByE2EId(t *testing.T) {
	sh := shell.NewShell(config.IPFS_URL)
	ipfsWrapper := ipfs.NewIPFSClientWrapper(sh)
	ipfsStorageRepository := ipfs.NewStorageRepository(ipfsWrapper)
	data := ipfsStorageRepository.GetById("QmVHKK8MwmB6FTywF7giespBej7eW7i4x7y8683ZbAENhj")
	assert.NotNil(t, data)
}*/
