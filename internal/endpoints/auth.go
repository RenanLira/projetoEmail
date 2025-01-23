package endpoints

import (
	"context"
	"net/http"
	authcredentials "projetoEmail/internal/infra/auth_credentials"
	internalerrors "projetoEmail/internal/internal_errors"
	"projetoEmail/internal/utils"
	"strings"
)

type AuthHandler struct {
	NewProvider func(ctx context.Context) (authcredentials.AuthProvider, error)
}

func (ah *AuthHandler) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token == "" {
			utils.SendJSON(w, internalerrors.NewErrUnauthorized(), nil)
			return
		}

		authorization := strings.Replace(token, "Bearer ", "", 1)

		provider, err := ah.NewProvider(r.Context())
		if err != nil {
			utils.SendJSON(w, internalerrors.NewErrInternal("error to connect to the provider"), nil)
			return
		}

		err = provider.VerifyToken(authorization)
		if err != nil {
			utils.SendJSON(w, internalerrors.NewErrUnauthorized(), nil)
			return
		}

		email := provider.GetClaimsToken(authorization)["email"].(string)
		ctx := context.WithValue(r.Context(), "email", email)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func NewAuthHandler() AuthHandler {
	return AuthHandler{NewProvider: authcredentials.NewAuthProvider}
}
