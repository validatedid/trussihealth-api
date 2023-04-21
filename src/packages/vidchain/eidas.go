package vidchain

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/validatedid/trussihealth-api/src/packages/config"
	"github.com/validatedid/trussihealth-api/src/packages/restClient"
)

type Eidas struct {
	httpClient    restClient.HTTPClient
	authenticator Authenticator
}

type EsealedVerifiableCredential struct {
	Content []byte
}

func NewEidas(client restClient.HTTPClient, apiAuthenticator Authenticator) (e *Eidas) {
	return &Eidas{httpClient: client, authenticator: apiAuthenticator}
}

func (e Eidas) EsealVc(payload VerifiableCredential) (esealedVerifiableCredential EsealedVerifiableCredential) {
	accessToken := e.authenticator.GetAccessToken()
	requestBody := fmt.Sprintf(`{
	"issuer": "%s",
	"payload": %s,
	"password": "%s"
	}`, config.ISSUER_DID, payload.Content, config.CERTIFICATE_PASSWORD)
	request, _ := http.NewRequest("POST", config.EIDAS_PATH, bytes.NewBufferString(requestBody))
	request.Header.Set("Authorization", "Bearer "+accessToken)
	request.Header.Set("Content-Type", "application/json")
	response, _ := e.httpClient.Do(request)
	body, _ := io.ReadAll(response.Body)
	esealedVerifiableCredential.Content = body
	return esealedVerifiableCredential
}
