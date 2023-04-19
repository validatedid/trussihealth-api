package vidchain

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/validatedid/trussihealth-api/src/packages/restClient"
)

type VidchainApi struct {
	httpClient  restClient.HTTPClient
	accessToken string
}

type VcPayload struct {
	DocumentId string
	Hash       string
}

type VerifiableCredentialRequestBody {

}

type VerifiableCredential struct {
}

type AccessTokenResponse struct {
	accessToken string
	tokenType   string
	expiresIn   int
	issuedAt    int
}

func NewVidchainApiConnector(client restClient.HTTPClient) (i VidchainApi) {
	return VidchainApi{httpClient: client}
}

func (v VidchainApi) Authenticate() {
	data := `{
	  "grantType": "urn:ietf:params:oauth:grant-type:jwt-bearer",
	  "assertion": "ewogICAgImlzcyI6ImVudGl0YXRTd2FnZ2VyIiwKICAgImF1ZCI6InZpZGNoYWluLWFwaSIsCiAgICJub25jZSI6InotMDQyN2RjMjUxNWIxIiwKICAgImNhbGxiYWNrVXJsIjoiaHR0cDovLzEyNy4wLjAuMTo4MDgwL2RlbW8vZW50aXRhdC1leGVtcGxlL2NhbGxiYWNrIiwKICAgImFwaUtleSI6ICI2MDAxMGMwZi05MmQ2LTQyMDYtYmFjYi1hMDRhYzA4MGVjNjMiCn0=",
	  "scope": "vidchain profile entity",
	  "expiresIn": 900
	}`
	request, _ := http.NewRequest("POST", "https://dev.vidchain.net/api/v1/sessions", bytes.NewBufferString(data))
	response, _ := v.httpClient.Do(request)
	body, _ := io.ReadAll(response.Body)
	var accessTokenResponse AccessTokenResponse
	json.Unmarshal(body, &accessTokenResponse)
	v.accessToken = accessTokenResponse.accessToken
}

func (v VidchainApi) CreateVc(payload VcPayload) (verifableCredential VerifiableCredential) {
	if v.isAccessTokenValid() {
		v.Authenticate()
	}
	data := `{
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
		  "id": "did:ethr:0x2Bb1629Dc1f992E00a9E170464BE3802ba259B3E",
		  "documentId"
		}
	  },
	  "options": {
		"revocable": false
	  }
	}`
	request, _ := http.NewRequest("POST", "https://dev.vidchain.net/api/v1/verifiable-credentials", bytes.NewBufferString(data))
	response, _ := v.httpClient.Do(request)
}

func (v VidchainApi) isAccessTokenValid() bool {
	if v.accessToken == "" {
		return false
	}
	token, _ := jwt.Parse(v.accessToken)

	// Check if the token is valid
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("Access token is valid")
		return true
	}
	return false
}
