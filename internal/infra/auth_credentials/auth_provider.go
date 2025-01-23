package authcredentials

import (
	"context"
	"os"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/golang-jwt/jwt/v5"
)

// import "os"

type AuthProvider interface {
	VerifyToken(token string) error
	GetClaimsToken(token string) jwt.MapClaims
}

type AuthProviderImp struct {
	Provider *oidc.Provider
	Context  context.Context
}

func (ap AuthProviderImp) VerifyToken(token string) error {
	verifier := ap.Provider.Verifier(&oidc.Config{ClientID: "projetoEmail"})
	_, err := verifier.Verify(ap.Context, token)

	return err
}

func (ap AuthProviderImp) GetClaimsToken(token string) jwt.MapClaims {
	tokenData, _ := jwt.Parse(token, nil)
	claims := tokenData.Claims.(jwt.MapClaims)

	return claims
}

func NewAuthProvider(ctx context.Context) (AuthProvider, error) {

	provider, err := oidc.NewProvider(ctx, os.Getenv("OIDC_PROVIDER"))

	if err != nil {
		return AuthProviderImp{}, err
	}

	return AuthProviderImp{
		Provider: provider,
		Context:  ctx,
	}, nil

}
