package vidchain_test

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/validatedid/trussihealth-api/src/packages/issuer"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

	holderDid := "did:ethr:0x2Bb1629Dc1f992E00a9E170464BE3802ba259B3E"
	documentHash := "5bc234eb44fee4ea5d6004dfda23cf824d49a20fd90a88be6c21dccb1d4ad09e"
	documentId := "123"
	var resp http.Response
	mockedResponse := fmt.Sprintf(`
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
          "id":  "%s",
		  "documentId": "%s",
		  "documentHash": "%s",
		},
		"proof": {
		  "type": "EcdsaSecp256k1Signature2019",
		  "created": "2023-04-19T14:53:20.000Z",
		  "proofPurpose": "assertionMethod",
		  "verificationMethod": "did:ethr:0xDfBA7E7D6fd9D3B5B900cE2aa3d9E6aA43574FC0#keys-1",
		  "jws": "eyJhbGciOiJFUzI1NksiLCJraWQiOiJkaWQ6ZXRocjoweERmQkE3RTdENmZkOUQzQjVCOTAwY0UyYWEzZDlFNmFBNDM1NzRGQzAja2V5cy0xIiwidHlwIjoiSldUIn0.eyJpYXQiOjE2ODE5MTYwMDAsImlzcyI6ImRpZDpldGhyOjB4RGZCQTdFN0Q2ZmQ5RDNCNUI5MDBjRTJhYTNkOUU2YUE0MzU3NEZDMCIsInZjIjp7IkBjb250ZXh0IjpbImh0dHBzOi8vd3d3LnczLm9yZy8yMDE4L2NyZWRlbnRpYWxzL3YxIl0sImNyZWRlbnRpYWxTdWJqZWN0Ijp7ImRvY3VtZW50SWQiOiJmYWtlLWRvY3VtZW50LWlkIiwiaGFzaCI6IjE4NWY4ZGIzMjI3MWZlMjVmNTYxYTZmYzkzOGIyZTI2NDMwNmVjMzA0ZWRhNTE4MDA3ZDE3NjQ4MjYzODE5NjkiLCJpZCI6ImRpZDpldGhyOjB4MkJiMTYyOURjMWY5OTJFMDBhOUUxNzA0NjRCRTM4MDJiYTI1OUIzRSJ9LCJpZCI6Imh0dHBzOi8vZXhhbXBsZS5jb20vY3JlZGVudGlhbC8yMzkwIiwiaXNzdWVyIjp7ImlkIjoiZGlkOmV0aHI6MHhEZkJBN0U3RDZmZDlEM0I1QjkwMGNFMmFhM2Q5RTZhQTQzNTc0RkMwIiwibmFtZSI6ImVudGl0YXRTd2FnZ2VyIn0sInR5cGUiOlsiVmVyaWZpYWJsZUNyZWRlbnRpYWwiLCJIZWFsdGhEYXRhQ3JlZGVudGlhbCJdLCJ2YWxpZFVudGlsIjoiMjAzMC0wMS0wMVQyMToxOToxMFoifX0.-ufL4sstCbn0narLNQ7cPpHt8vIvWn43DTXD07lfo8fJMqwjeAMRTrWYv6F4m6QqCxmdVH5L7BLt4wpvdhchxg"
		}
	  }`, holderDid, documentId, documentHash)
	resp.Body = io.NopCloser(bytes.NewBufferString(mockedResponse))
	mockHttpClient.On("Do", mock.Anything).Return(&resp, nil)

	credential := issuer.NewCredential(mockHttpClient)
	vcPayload := issuer.VcPayload{DocumentId: documentId,
		Hash: documentHash}
	credential.CreateVc(vcPayload, holderDid)

	data := fmt.Sprintf(`{
	  "credential": {
		"id": "https://example.com/credential/2390",
		"issuer": {
		  "id": "did:ethr:0x9A668DBe392230c407beC406CC4bb965C16CAbee",
		  "name": "entitatSwagger"
		},
		"type": [
		  "VerifiableCredential",
		  "HealthDataCredential"
		],
		"validUntil": "2030-01-01T21:19:10Z",
		"credentialSubject": {
		  "id": "%s",
		  "documentId": "%s",
 		  "documentHash": "%s"
		}
	  },
	  "options": {
		"revocable": false
	  }
	}`, holderDid, documentId, documentHash)
	expectedRequest, _ := http.NewRequest("POST", "https://url", bytes.NewBufferString(data))

	calledRequest := mockHttpClient.Calls[0].Arguments[0].(*http.Request)
	assert.Equal(t, calledRequest.Body, expectedRequest.Body)
}
