package endpoints

import (
	"net/http"
	internalerrors "projetoEmail/internal/internal_errors"
	"projetoEmail/internal/utils"
	"strings"

	"github.com/coreos/go-oidc/v3/oidc"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token == "" {
			utils.SendJSON(w, internalerrors.NewErrUnauthorized(), nil)
			return
		}

		authorization := strings.Replace(token, "Bearer ", "", 1)

		provider, err := oidc.NewProvider(r.Context(), "http://localhost:8081/realms/providerGoEmail")
		if err != nil {
			utils.SendJSON(w, internalerrors.NewErrInternal("error to connect to the provider"), nil)
			return
		}

		verifier := provider.Verifier(&oidc.Config{ClientID: "projetoEmail"})
		_, err = verifier.Verify(r.Context(), authorization)
		if err != nil {
			utils.SendJSON(w, internalerrors.NewErrUnauthorized(), nil)
			return
		}


		next.ServeHTTP(w, r)
	})
}
