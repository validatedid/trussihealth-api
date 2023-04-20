package vidchain

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/validatedid/trussihealth-api/src/packages/restClient"
	"io"
	"net/http"
	"time"
)

type ApiAuthenticator struct {
	httpClient  restClient.HTTPClient
	accessToken string
}

func NewVidchainApiAuthenticator(client restClient.HTTPClient) (a ApiAuthenticator) {
	return ApiAuthenticator{httpClient: client}
}

func (a ApiAuthenticator) GetAccessToken() string {
	valid := a.isAccessTokenValid()
	if !valid {
		a.authenticate()
	}
	return a.accessToken
}

type AccessTokenResponse struct {
	accessToken string
	tokenType   string
	expiresIn   int
	issuedAt    int
}

func (a ApiAuthenticator) authenticate() {
	data := `{
	  "grantType": "urn:ietf:params:oauth:grant-type:jwt-bearer",
	  "assertion": "ewogICAgImlzcyI6ImVudGl0YXRTd2FnZ2VyIiwKICAgImF1ZCI6InZpZGNoYWluLWFwaSIsCiAgICJub25jZSI6InotMDQyN2RjMjUxNWIxIiwKICAgImNhbGxiYWNrVXJsIjoiaHR0cDovLzEyNy4wLjAuMTo4MDgwL2RlbW8vZW50aXRhdC1leGVtcGxlL2NhbGxiYWNrIiwKICAgImFwaUtleSI6ICI2MDAxMGMwZi05MmQ2LTQyMDYtYmFjYi1hMDRhYzA4MGVjNjMiCn0=",
	  "scope": "vidchain profile entity",
	  "expiresIn": 900
	}`
	request, _ := http.NewRequest("POST", "https://dev.vidchain.net/api/v1/sessions", bytes.NewBufferString(data))
	response, _ := a.httpClient.Do(request)
	body, _ := io.ReadAll(response.Body)
	var accessTokenResponse AccessTokenResponse
	json.Unmarshal(body, &accessTokenResponse)
	a.accessToken = accessTokenResponse.accessToken
}

func (a ApiAuthenticator) isAccessTokenValid() bool {
	if a.accessToken == "" {
		return false
	}
	err := checkJWTExpiration(a.accessToken)
	if err != nil {
		return false
	}
	return true
}

func checkJWTExpiration(tokenString string) error {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return fmt.Errorf("Failed to parse JWT token: %v", err)
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if exp, ok := claims["exp"].(float64); ok {
			if time.Now().Unix() > int64(exp) {
				return fmt.Errorf("Expired JWT token")
			}
			// Token is not expired
			return nil
		}
		return fmt.Errorf("Missing or invalid 'exp' claim in JWT token")
	}
	return fmt.Errorf("Invalid JWT token claims")
}
