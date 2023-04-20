package vidchain_test

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/validatedid/trussihealth-api/src/packages/config"
	"github.com/validatedid/trussihealth-api/src/packages/vidchain"
	"io"
	"net/http"
	"testing"
)

type httpClientMock struct {
	mock.Mock
}

func newHttpClientMock() *httpClientMock {
	return &httpClientMock{}
}

func (m *httpClientMock) Do(req *http.Request) (*http.Response, error) {
	args := m.Called(req)
	return args.Get(0).(*http.Response), args.Error(1)
}

func TestGetAccessTokenWhenNotValid(t *testing.T) {
	expectedAccessToken := "eyJhbGciOiJFUzI1NksiLCJ0eXAiOiJKV1QifQ.eyJzdWIiOiJlbnRpdGF0U3dhZ2dlciIsImRpZCI6ImRpZDpldGhyOjB4RGZCQTdFN0Q2ZmQ5RDNCNUI5MDBjRTJhYTNkOUU2YUE0MzU3NEZDMCIsIm5vbmNlIjoiei0wNDI3ZGMyNTE1YjEiLCJpYXQiOjE2ODE5NzI4NzYsImV4cCI6MTY4MTk3Mzc3NiwiYXVkIjoidmlkY2hhaW4tYXBpIn0.jraYKsU3h7BenBvp71xNNHEX_537DLCCT9nNR3LxjmSMZsOKYdvLDIYPizQ3jySa4uuwtyA55uE93rtDFWRBDQ"
	mockHttpClient := newHttpClientMock()
	apiAuthenticator := vidchain.NewVidchainApiAuthenticator(mockHttpClient)
	mockedResponse := `{ "accessToken": "eyJhbGciOiJFUzI1NksiLCJ0eXAiOiJKV1QifQ.eyJzdWIiOiJlbnRpdGF0U3dhZ2dlciIsImRpZCI6ImRpZDpldGhyOjB4RGZCQTdFN0Q2ZmQ5RDNCNUI5MDBjRTJhYTNkOUU2YUE0MzU3NEZDMCIsIm5vbmNlIjoiei0wNDI3ZGMyNTE1YjEiLCJpYXQiOjE2ODE5NzI4NzYsImV4cCI6MTY4MTk3Mzc3NiwiYXVkIjoidmlkY2hhaW4tYXBpIn0.jraYKsU3h7BenBvp71xNNHEX_537DLCCT9nNR3LxjmSMZsOKYdvLDIYPizQ3jySa4uuwtyA55uE93rtDFWRBDQ", "tokenType": "Bearer", "expiresIn": 1681973776, "issuedAt": 1681972876 }`
	var resp http.Response
	resp.Body = io.NopCloser(bytes.NewBufferString(mockedResponse))
	mockHttpClient.On("Do", mock.Anything).Return(&resp, nil)

	accessToken := apiAuthenticator.GetAccessToken()

	assert.Equal(t, expectedAccessToken, accessToken)

	authenticationBody := vidchain.AuthenticationBody{GrantType: "urn:ietf:params:oauth:grant-type:jwt-bearer", Scope: "vidchain profile entity", ExpiresIn: 900, Assertion: config.TRUSSIHEALTH_ASSERTION}
	jsonBytes, _ := json.Marshal(authenticationBody)
	expectedRequest, _ := http.NewRequest("POST", "https://url", bytes.NewBuffer(jsonBytes))

	calledRequest := mockHttpClient.Calls[0].Arguments[0].(*http.Request)
	assert.Equal(t, calledRequest.Body, expectedRequest.Body)

}

func TestGetAccessTokenAlreadyValid(t *testing.T) {
	mockHttpClient := newHttpClientMock()
	apiAuthenticator := vidchain.NewVidchainApiAuthenticator(mockHttpClient)
	// token mocked expires in a hundred years
	mockedResponse := `{ "accessToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwiZXhwIjo2NjUzNjU1NzU1fQ.LNwFTs14NpZYIedYkP8Nmk2d93nXgRZvTzTkbKj3s3s", "tokenType": "Bearer", "expiresIn": 1681973776, "issuedAt": 1681972876 }`
	var resp http.Response
	resp.Body = io.NopCloser(bytes.NewBufferString(mockedResponse))
	mockHttpClient.On("Do", mock.Anything).Return(&resp, nil)

	accessToken := apiAuthenticator.GetAccessToken()
	alreadyValidAccessToken := apiAuthenticator.GetAccessToken()

	assert.Equal(t, accessToken, alreadyValidAccessToken)
	mockHttpClient.AssertNumberOfCalls(t, "Do", 1)
}
