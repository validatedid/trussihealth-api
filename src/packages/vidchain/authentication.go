package vidchain

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/validatedid/trussihealth-api/src/packages/config"
	"github.com/validatedid/trussihealth-api/src/packages/restClient"
)

type Authenticator interface {
	GetAccessToken() string
}

type ApiAuthenticator struct {
	httpClient  restClient.HTTPClient
	accessToken string
}

func NewVidchainApiAuthenticator(client restClient.HTTPClient) (a *ApiAuthenticator) {
	return &ApiAuthenticator{httpClient: client}
}

func (a *ApiAuthenticator) GetAccessToken() string {
	valid := a.isAccessTokenValid()
	if !valid {
		a.authenticate()
	}
	return a.accessToken
}

type accessTokenResponse struct {
	AccessToken string `json:"accessToken"`
	TokenType   string `json:"tokenType"`
	ExpiresIn   int    `json:"expiresIn"`
	IssuedAt    int    `json:"issuedAt"`
}

type AuthenticationBody struct {
	GrantType string `json:"grantType"`
	Assertion string `json:"assertion"`
	Scope     string `json:"scope"`
	ExpiresIn int64  `json:"expiresIn"`
}

func (a *ApiAuthenticator) authenticate() {
	authenticationBody := AuthenticationBody{GrantType: "urn:ietf:params:oauth:grant-type:jwt-bearer", Scope: "vidchain profile entity", ExpiresIn: 900, Assertion: config.TRUSSIHEALTH_ASSERTION}
	jsonBytes, _ := json.Marshal(authenticationBody)
	request, _ := http.NewRequest("POST", config.SESSIONS_PATH, bytes.NewBuffer(jsonBytes))
	response, _ := a.httpClient.Do(request)
	body, _ := io.ReadAll(response.Body)
	var accessTokenResponse accessTokenResponse
	json.Unmarshal(body, &accessTokenResponse)
	a.accessToken = accessTokenResponse.AccessToken
}

func (a *ApiAuthenticator) isAccessTokenValid() bool {
	if a.accessToken == "" {
		return false
	}
	expired := checkJWTExpiration(a.accessToken)
	if expired {
		return false
	}
	return true
}

type Claims struct {
	Sub   string `json:"sub"`
	Did   string `json:"did"`
	Nonce string `json:"nonce"`
	Iat   int64  `json:"iat"`
	Exp   int64  `json:"exp"`
	Aud   string `json:"aud"`
}

func checkJWTExpiration(tokenString string) bool {
	parts := strings.Split(tokenString, ".")
	payloadBytes, _ := base64.RawURLEncoding.DecodeString(parts[1])
	var claims Claims
	json.Unmarshal(payloadBytes, &claims)
	if time.Now().Unix() > claims.Exp {
		fmt.Println("Token is expired")
		return true
	}
	return false
}
