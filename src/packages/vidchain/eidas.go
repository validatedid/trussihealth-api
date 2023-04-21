package vidchain

import (
	"bytes"
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
	request, _ := http.NewRequest("POST", config.EIDAS_PATH, bytes.NewBuffer(payload.Content))
	request.Header.Set("Authorization", "Bearer "+accessToken)
	response, _ := e.httpClient.Do(request)
	body, _ := io.ReadAll(response.Body)
	esealedVerifiableCredential.Content = body
	return esealedVerifiableCredential
}
