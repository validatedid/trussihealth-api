package issuer

import (
	"net/http"

	"github.com/stretchr/testify/mock"
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
