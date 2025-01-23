package endpoints_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"projetoEmail/internal/endpoints"
	authcredentials "projetoEmail/internal/infra/auth_credentials"
	internalerrors "projetoEmail/internal/internal_errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockHandler struct {
	called bool
}

func (m *mockHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	m.called = true
}

type AuthProviderMock struct {
	mock.Mock
}

func (ap *AuthProviderMock) VerifyToken(token string) error {
	args := ap.Called(token)

	return args.Error(0)
}

func (ap *AuthProviderMock) GetClaimsToken(token string) jwt.MapClaims {
	args := ap.Called(token)

	return args.Get(0).(jwt.MapClaims)
}

func Test_Auth_WhenAuthorizationIsMissing_ReturnError(t *testing.T) {

	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()
	authHandler := endpoints.AuthHandler{NewProvider: nil}
	handler := authHandler.Auth(&mockHandler{})

	handler.ServeHTTP(res, req)

	assert.Equal(t, http.StatusUnauthorized, res.Code)
	assert.Contains(t, res.Body.String(), internalerrors.NewErrUnauthorized().Error())

}

func Test_Auth_WhenProviderFailsToConnect_ReturnError(t *testing.T) {

	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "Bearer token")
	res := httptest.NewRecorder()

	authHandler := endpoints.AuthHandler{NewProvider: func(ctx context.Context) (authcredentials.AuthProvider, error) {
		return nil, internalerrors.NewErrInternal("error to connect to the provider")
	}}

	handler := authHandler.Auth(&mockHandler{})

	handler.ServeHTTP(res, req)

	assert.Equal(t, http.StatusInternalServerError, res.Code)

}

func Test_Auth_WhenTokenIsInvalid_ReturnError(t *testing.T) {

	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "Bearer token")
	res := httptest.NewRecorder()

	authProvider := new(AuthProviderMock)

	authHandler := endpoints.AuthHandler{NewProvider: func(ctx context.Context) (authcredentials.AuthProvider, error) {
		return authProvider, nil
	}}

	authProvider.On("VerifyToken", "token").Return(internalerrors.NewErrUnauthorized())

	handler := authHandler.Auth(&mockHandler{})
	handler.ServeHTTP(res, req)

	assert.Equal(t, http.StatusUnauthorized, res.Code)
}

func Test_Auth_WhenTokenIsValid_ReturnSuccess(t *testing.T) {

	emailExpected := "test@exemple.com"

	var emailActual string
	nextHttp := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		emailActual = r.Context().Value("email").(string)
	})

	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "Bearer token")

	res := httptest.NewRecorder()

	authProvider := new(AuthProviderMock)

	authHandler := endpoints.AuthHandler{
		NewProvider: func(ctx context.Context) (authcredentials.AuthProvider, error) {
			return authProvider, nil
		},
	}

	authProvider.On("VerifyToken", "token").Return(nil)
	authProvider.On("GetClaimsToken", "token").Return(jwt.MapClaims{"email": emailExpected})

	handler := authHandler.Auth(nextHttp)
	handler.ServeHTTP(res, req)

	assert.Equal(t, res.Code, http.StatusOK)
	assert.Equal(t, emailActual, emailExpected)
}
