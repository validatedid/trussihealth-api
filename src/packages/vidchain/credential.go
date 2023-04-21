package vidchain

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/validatedid/trussihealth-api/src/packages/config"
	"github.com/validatedid/trussihealth-api/src/packages/restClient"
)

type Credential struct {
	httpClient    restClient.HTTPClient
	authenticator Authenticator
}

type VcPayload struct {
	DocumentId string
	Hash       string
}

type VerifiableCredential struct {
	Content []byte
}

func NewCredential(client restClient.HTTPClient, apiAuthenticator Authenticator) (c *Credential) {
	return &Credential{httpClient: client, authenticator: apiAuthenticator}
}

func (c Credential) CreateVc(payload VcPayload, did string) (verifiableCredential VerifiableCredential) {
	accessToken := c.authenticator.GetAccessToken()
	data := fmt.Sprintf(`{
	  "credential": {
		"id": "https://example.com/credential/2390",
		"issuer": {
		  "id": "did:ethr:0xDfBA7E7D6fd9D3B5B900cE2aa3d9E6aA43574FC0",
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
 		  "documentHash": "%s",
		}
	  },
	  "options": {
		"revocable": false
	  }
	}`, did, payload.DocumentId, payload.Hash)
	request, _ := http.NewRequest("POST", config.VERIFIABLE_CREDENTIAL_PATH, bytes.NewBufferString(data))
	request.Header.Set("Authorization", "Bearer "+accessToken)
	response, _ := c.httpClient.Do(request)
	body, _ := io.ReadAll(response.Body)
	verifiableCredential.Content = body
	return verifiableCredential
}
