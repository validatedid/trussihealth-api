package issuer

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/validatedid/trussihealth-api/src/packages/config"
	"github.com/validatedid/trussihealth-api/src/packages/restClient"
)

type Eidas struct {
	httpClient restClient.HTTPClient
}

type EsealedVerifiableCredential struct {
	Content []byte
}

func NewEidas(client restClient.HTTPClient) (e *Eidas) {
	return &Eidas{httpClient: client}
}

func (e Eidas) EsealVc(payload VerifiableCredential) error {
	requestBody := fmt.Sprintf(`{
	"payload": %s,
	"password": "%s"
	}`, config.ISSUER_DID, payload.Content, config.CERTIFICATE_PASSWORD)
	request, _ := http.NewRequest("POST", config.ESEAL_API, bytes.NewBufferString(requestBody))
	request.Header.Set("Content-Type", "application/json")
	_, err := e.httpClient.Do(request)
	return err
}
