package vidchain_test

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/validatedid/trussihealth-api/src/packages/vidchain"
	"io"
	"net/http"
	"testing"
)

type ApiAuthenticatorMock struct {
	mock.Mock
}

func newApiAuthenticatorMock() *ApiAuthenticatorMock {
	return &ApiAuthenticatorMock{}
}

func (a *ApiAuthenticatorMock) GetAccessToken() string {
	args := a.Called()
	return args.String(0)
}

func TestCreateCredential(t *testing.T) {
	mockHttpClient := newHttpClientMock()
	apiAuthenticatorMock := newApiAuthenticatorMock()

	apiAuthenticatorMock.On("GetAccessToken").Return("eyJhbGciOiJFUzI1NksiLCJ0eXAiOiJKV1QifQ.eyJzdWIiOiJlbnRpdGF0U3dhZ2dlciIsImRpZCI6ImRpZDpldGhyOjB4RGZCQTdFN0Q2ZmQ5RDNCNUI5MDBjRTJhYTNkOUU2YUE0MzU3NEZDMCIsIm5vbmNlIjoiei0wNDI3ZGMyNTE1YjEiLCJpYXQiOjE2ODE5NzI4NzYsImV4cCI6MTY4MTk3Mzc3NiwiYXVkIjoidmlkY2hhaW4tYXBpIn0.jraYKsU3h7BenBvp71xNNHEX_537DLCCT9nNR3LxjmSMZsOKYdvLDIYPizQ3jySa4uuwtyA55uE93rtDFWRBDQ")
	var resp http.Response
	mockedResponse := `
		{
		"@context": [
		  "https://www.w3.org/2018/credentials/v1"
		],
		"id": "https://example.com/credential/2390",
		"type": [
		  "VerifiableCredential",
		  "HealthDataCredential"
		],
		"issuer": {
		  "id": "did:ethr:0xDfBA7E7D6fd9D3B5B900cE2aa3d9E6aA43574FC0",
		  "name": "entitatSwagger"
		},
		"issuanceDate": "2023-04-19T14:53:20.000Z",
		"validUntil": "2030-01-01T21:19:10Z",
		"credentialSubject": {
		  "documentId": "fake-document-id",
		  "hash": "185f8db32271fe25f561a6fc938b2e264306ec304eda518007d1764826381969",
		  "id": "did:ethr:0x2Bb1629Dc1f992E00a9E170464BE3802ba259B3E"
		},
		"proof": {
		  "type": "EcdsaSecp256k1Signature2019",
		  "created": "2023-04-19T14:53:20.000Z",
		  "proofPurpose": "assertionMethod",
		  "verificationMethod": "did:ethr:0xDfBA7E7D6fd9D3B5B900cE2aa3d9E6aA43574FC0#keys-1",
		  "jws": "eyJhbGciOiJFUzI1NksiLCJraWQiOiJkaWQ6ZXRocjoweERmQkE3RTdENmZkOUQzQjVCOTAwY0UyYWEzZDlFNmFBNDM1NzRGQzAja2V5cy0xIiwidHlwIjoiSldUIn0.eyJpYXQiOjE2ODE5MTYwMDAsImlzcyI6ImRpZDpldGhyOjB4RGZCQTdFN0Q2ZmQ5RDNCNUI5MDBjRTJhYTNkOUU2YUE0MzU3NEZDMCIsInZjIjp7IkBjb250ZXh0IjpbImh0dHBzOi8vd3d3LnczLm9yZy8yMDE4L2NyZWRlbnRpYWxzL3YxIl0sImNyZWRlbnRpYWxTdWJqZWN0Ijp7ImRvY3VtZW50SWQiOiJmYWtlLWRvY3VtZW50LWlkIiwiaGFzaCI6IjE4NWY4ZGIzMjI3MWZlMjVmNTYxYTZmYzkzOGIyZTI2NDMwNmVjMzA0ZWRhNTE4MDA3ZDE3NjQ4MjYzODE5NjkiLCJpZCI6ImRpZDpldGhyOjB4MkJiMTYyOURjMWY5OTJFMDBhOUUxNzA0NjRCRTM4MDJiYTI1OUIzRSJ9LCJpZCI6Imh0dHBzOi8vZXhhbXBsZS5jb20vY3JlZGVudGlhbC8yMzkwIiwiaXNzdWVyIjp7ImlkIjoiZGlkOmV0aHI6MHhEZkJBN0U3RDZmZDlEM0I1QjkwMGNFMmFhM2Q5RTZhQTQzNTc0RkMwIiwibmFtZSI6ImVudGl0YXRTd2FnZ2VyIn0sInR5cGUiOlsiVmVyaWZpYWJsZUNyZWRlbnRpYWwiLCJIZWFsdGhEYXRhQ3JlZGVudGlhbCJdLCJ2YWxpZFVudGlsIjoiMjAzMC0wMS0wMVQyMToxOToxMFoifX0.-ufL4sstCbn0narLNQ7cPpHt8vIvWn43DTXD07lfo8fJMqwjeAMRTrWYv6F4m6QqCxmdVH5L7BLt4wpvdhchxg"
		}
	  }`
	resp.Body = io.NopCloser(bytes.NewBufferString(mockedResponse))
	mockHttpClient.On("Do", mock.Anything).Return(&resp, nil)

	credential := vidchain.NewCredential(mockHttpClient, apiAuthenticatorMock)
	vcPayload := vidchain.VcPayload{DocumentId: "123",
		Hash: "5bc234eb44fee4ea5d6004dfda23cf824d49a20fd90a88be6c21dccb1d4ad09e"}
	expectedVc := []byte(mockedResponse)
	vc := credential.CreateVc(vcPayload)
	assert.Equal(t, expectedVc, vc.Content)
}