package vidchain_test

import (
	"bytes"
	"fmt"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/validatedid/trussihealth-api/src/packages/config"
	"github.com/validatedid/trussihealth-api/src/packages/issuer"

	"io"
	"net/http"
	"testing"
)

func TestEsealVc(t *testing.T) {
	mockHttpClient := newHttpClientMock()

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
		  "proof": [
			{
			  "type": "EcdsaSecp256k1Signature2019",
			  "created": "2023-04-19T14:53:20.000Z",
			  "proofPurpose": "assertionMethod",
			  "verificationMethod": "did:ethr:0xDfBA7E7D6fd9D3B5B900cE2aa3d9E6aA43574FC0#keys-1",
			  "jws": "eyJhbGciOiJFUzI1NksiLCJraWQiOiJkaWQ6ZXRocjoweERmQkE3RTdENmZkOUQzQjVCOTAwY0UyYWEzZDlFNmFBNDM1NzRGQzAja2V5cy0xIiwidHlwIjoiSldUIn0.eyJpYXQiOjE2ODE5MTYwMDAsImlzcyI6ImRpZDpldGhyOjB4RGZCQTdFN0Q2ZmQ5RDNCNUI5MDBjRTJhYTNkOUU2YUE0MzU3NEZDMCIsInZjIjp7IkBjb250ZXh0IjpbImh0dHBzOi8vd3d3LnczLm9yZy8yMDE4L2NyZWRlbnRpYWxzL3YxIl0sImNyZWRlbnRpYWxTdWJqZWN0Ijp7ImRvY3VtZW50SWQiOiJmYWtlLWRvY3VtZW50LWlkIiwiaGFzaCI6IjE4NWY4ZGIzMjI3MWZlMjVmNTYxYTZmYzkzOGIyZTI2NDMwNmVjMzA0ZWRhNTE4MDA3ZDE3NjQ4MjYzODE5NjkiLCJpZCI6ImRpZDpldGhyOjB4MkJiMTYyOURjMWY5OTJFMDBhOUUxNzA0NjRCRTM4MDJiYTI1OUIzRSJ9LCJpZCI6Imh0dHBzOi8vZXhhbXBsZS5jb20vY3JlZGVudGlhbC8yMzkwIiwiaXNzdWVyIjp7ImlkIjoiZGlkOmV0aHI6MHhEZkJBN0U3RDZmZDlEM0I1QjkwMGNFMmFhM2Q5RTZhQTQzNTc0RkMwIiwibmFtZSI6ImVudGl0YXRTd2FnZ2VyIn0sInR5cGUiOlsiVmVyaWZpYWJsZUNyZWRlbnRpYWwiLCJIZWFsdGhEYXRhQ3JlZGVudGlhbCJdLCJ2YWxpZFVudGlsIjoiMjAzMC0wMS0wMVQyMToxOToxMFoifX0.-ufL4sstCbn0narLNQ7cPpHt8vIvWn43DTXD07lfo8fJMqwjeAMRTrWYv6F4m6QqCxmdVH5L7BLt4wpvdhchxg"
			},
			{
			  "type": "CAdESRSASignature2020",
			  "created": "2023-04-20T13:37:07Z",
			  "proofPurpose": "assertionMethod",
			  "verificationMethod": "did:ethr:0xDfBA7E7D6fd9D3B5B900cE2aa3d9E6aA43574FC0#MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA5GX6o+OfaZHbeRi3pWpkskkD96MJMmDl1tTS0y693Qv0cxPMd0kaR0MGkRK5tFWL1d+SVb7mxmYGrZuUjbE1B8O2MK8MvoRhQl6NiTmmy9JY9xKZKUT6VUuxlwJTMU732YocPCIJgN2RJ4ZwmG8nyoggDplRi3P3VYhU1LpNjureoJAHS9DIc1OpwgQWq4qpcFkMV3+8/pxbmwawYqNpfxi+j3+qK7wf+u1dj0BGJ2JX9uFbX3Vd+P0HFN0/W0IzdkYUs0/R2K2GbRFxJkuAEWjeJkqAp5CEzbKxlM9y6V9KM8ZNpEAlHvPb/E1HfrmMwP7B1yV2u6B6kwNWSVa7QwIDAQAB",
			  "cades": "-----BEGIN PKCS7-----MIIH2QYJKoZIhvcNAQcCoIIHyjCCB8YCAQExDzANBglghkgBZQMEAgEFADBPBgkqhkiG9w0BBwGgQgRA47DEQpj8HBSa+/TImW+5JCeuQeRkm5NMpJWZG3hSuFVftFPPnGtq3FQ27NIeCdv049V/x+AmzifMlzSA3QJZXKCCBBcwggQTMIIC+6ADAgECAhQnSFMmzAXkEGt5CBsRAvwVMtfAcDANBgkqhkiG9w0BAQsFADCBmDELMAkGA1UEBhMCTkwxFjAUBgNVBAgMDU5vb3JkLUhvbGxhbmQxEjAQBgNVBAcMCUFtc3RlcmRhbTEUMBIGA1UECgwLVGVzdENvbXBhbnkxCzAJBgNVBAsMAklUMRUwEwYDVQQDDAxTY290dCBNYWxsZXkxIzAhBgkqhkiG9w0BCQEWFHNtYWxsZXlAc3BoZXJlb24uY29tMB4XDTIwMTIwMjE0MzEwNVoXDTMwMTEzMDE0MzEwNVowgZgxCzAJBgNVBAYTAk5MMRYwFAYDVQQIDA1Ob29yZC1Ib2xsYW5kMRIwEAYDVQQHDAlBbXN0ZXJkYW0xFDASBgNVBAoMC1Rlc3RDb21wYW55MQswCQYDVQQLDAJJVDEVMBMGA1UEAwwMU2NvdHQgTWFsbGV5MSMwIQYJKoZIhvcNAQkBFhRzbWFsbGV5QHNwaGVyZW9uLmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAORl+qPjn2mR23kYt6VqZLJJA/ejCTJg5dbU0tMuvd0L9HMTzHdJGkdDBpESubRVi9XfklW+5sZmBq2blI2xNQfDtjCvDL6EYUJejYk5psvSWPcSmSlE+lVLsZcCUzFO99mKHDwiCYDdkSeGcJhvJ8qIIA6ZUYtz91WIVNS6TY7q3qCQB0vQyHNTqcIEFquKqXBZDFd/vP6cW5sGsGKjaX8Yvo9/qiu8H/rtXY9ARidiV/bhW191Xfj9BxTdP1tCM3ZGFLNP0dithm0RcSZLgBFo3iZKgKeQhM2ysZTPculfSjPGTaRAJR7z2/xNR365jMD+wdcldrugepMDVklWu0MCAwEAAaNTMFEwHQYDVR0OBBYEFPZvjfNUC6pVbAUhktJNI50jefkgMB8GA1UdIwQYMBaAFPZvjfNUC6pVbAUhktJNI50jefkgMA8GA1UdEwEB/wQFMAMBAf8wDQYJKoZIhvcNAQELBQADggEBAGcrKcTfDTgXIQelgmz3h5iu8oCAfqW9q8eUNuJnGrViA/BA3uVsUojJnGx9zHVAa7VGT974HsP6l4hga0vT6VDFsLzyiCHTcJo+qskSGGZLEsQ34vcdCX6ki1VRMofyRrEd4LzrKbak4QlbbLuZW09iSuAyJOOY4mPrSEkRATztVP9MH+ygBc6byhVfLUdrcfAdni+i4b9LBiFdZqqc5QM6J9Hnv0LLnewyS4bSbtK5fE5mnbSOdftNpHwSUfyZmJ9NIGORPZE2PVCEUb5OXmdH9DrePjEjD1/ZKP6NVyWy4hwRCHTRujyl90yED6wxfXMKw68bewOex+youYvuC5UxggNCMIIDPgIBATCBsTCBmDELMAkGA1UEBhMCTkwxFjAUBgNVBAgMDU5vb3JkLUhvbGxhbmQxEjAQBgNVBAcMCUFtc3RlcmRhbTEUMBIGA1UECgwLVGVzdENvbXBhbnkxCzAJBgNVBAsMAklUMRUwEwYDVQQDDAxTY290dCBNYWxsZXkxIzAhBgkqhkiG9w0BCQEWFHNtYWxsZXlAc3BoZXJlb24uY29tAhQnSFMmzAXkEGt5CBsRAvwVMtfAcDANBglghkgBZQMEAgEFAKCCAWEwGAYJKoZIhvcNAQkDMQsGCSqGSIb3DQEHATAcBgkqhkiG9w0BCQUxDxcNMjMwNDIwMTMzNzA4WjAvBgkqhkiG9w0BCQQxIgQg2vswwh6l/CQEEaaP3RHu3cCKmUXOmd2ehC0+qAlSHw0wgfUGCyqGSIb3DQEJEAIvMYHlMIHiMIHfMIHcBCD3SVjat5f3AQRhEKblwdKspJ3ZS87w3K1LUXSqPg5RPzCBtzCBnqSBmzCBmDELMAkGA1UEBhMCTkwxFjAUBgNVBAgMDU5vb3JkLUhvbGxhbmQxEjAQBgNVBAcMCUFtc3RlcmRhbTEUMBIGA1UECgwLVGVzdENvbXBhbnkxCzAJBgNVBAsMAklUMRUwEwYDVQQDDAxTY290dCBNYWxsZXkxIzAhBgkqhkiG9w0BCQEWFHNtYWxsZXlAc3BoZXJlb24uY29tAhQnSFMmzAXkEGt5CBsRAvwVMtfAcDANBgkqhkiG9w0BAQsFAASCAQAOW30jpCZaJRD/Pm9ET0quDskIFdt/1/amqdhUuXDjnYWlHESTgJF/NIfdAfQvAFLGdqmc++0ccEgEYAnDyI2iLM5wTrQJhp60MAD6dauQ1cPHC+ZwC5xMqO9fklVQJi65Kt6bF1zLcZjIxlPq4tKBB/Ennx+0tJGkA0C68BuqX1b3qlaa0pFn8XdeCKLFiK5dv1A8cQE2fPB2DrmK8j2xd9jqNra9tYQEGT4vqDRPemQhiSiNiTqs+eNsgCjJsdMMmAOmdWkefMBnAW+vSA6CpPajUdz9ziaT8x53+lZ4l9L15U4W33x7kCFcTrHFR0dStq4jgwcwzzHLD1XnD/wI-----END PKCS7-----"
			}
		  ]
		}`
	resp.Body = io.NopCloser(bytes.NewBufferString(mockedResponse))
	mockHttpClient.On("Do", mock.Anything).Return(&resp, nil)

	verifiableCredential := `
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
	eidas := issuer.NewEidas(mockHttpClient)
	vcPayload := issuer.VerifiableCredential{Content: []byte(verifiableCredential)}

	err := eidas.EsealVc(vcPayload)

	assert.Nil(t, err)

	data := fmt.Sprintf(`{
	"issuer": "%s",
	"payload": %s,
	"password": "%s"
	}`, config.ISSUER_DID, vcPayload.Content, config.CERTIFICATE_PASSWORD)
	expectedRequestEseal, _ := http.NewRequest("POST", "https://url", bytes.NewBufferString(data))

	calledRequest := mockHttpClient.Calls[0].Arguments[0].(*http.Request)
	assert.Equal(t, calledRequest.Body, expectedRequestEseal.Body)
}
