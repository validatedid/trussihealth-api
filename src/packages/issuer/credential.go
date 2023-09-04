package issuer

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/validatedid/trussihealth-api/src/packages/config"
	"github.com/validatedid/trussihealth-api/src/packages/restClient"
)

type Credential struct {
	httpClient restClient.HTTPClient
}

type VcPayload struct {
	DocumentId string
	Hash       string
}

type VerifiableCredential struct {
	Content []byte
}

func NewCredential(client restClient.HTTPClient) (c *Credential) {
	return &Credential{httpClient: client}
}

func (c Credential) CreateVc(payload VcPayload, did string) (verifiableCredential VerifiableCredential) {
	issuer := config.ISSUER_DID
	data := fmt.Sprintf(`{
	  "credential": {
		"id": "https://example.com/credential/2390",
		"issuer": {
		  "id": "%s",
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
	}`, issuer, did, payload.DocumentId, payload.Hash)
	request, _ := http.NewRequest("POST", config.VC_API, bytes.NewBufferString(data))
	request.Header.Set("Content-Type", "application/json")
	response, _ := c.httpClient.Do(request)
	body, _ := io.ReadAll(response.Body)
	verifiableCredential.Content = body
	return verifiableCredential
}
