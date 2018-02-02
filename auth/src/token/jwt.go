package token

import (
	"crypto/rsa"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Claim model
type Claim struct {
	Issuer   string
	Audience string
	Subject  string
}

// AccessToken model
type AccessToken struct {
	Token string `json:"accessToken"`
}

// AccessTokenResponse response model
type AccessTokenResponse struct {
	Error       error
	AccessToken *AccessToken
}

// AccessTokenGenerator interface
type AccessTokenGenerator interface {
	GenerateAccessToken(cl Claim) <-chan AccessTokenResponse
}

// jwtGenerator private model
type jwtGenerator struct {
	signKey  *rsa.PrivateKey
	tokenAge time.Duration
}

// NewJwtGenerator function for initialise jwtGenerator object
func NewJwtGenerator(signKey *rsa.PrivateKey, tokenAge time.Duration) AccessTokenGenerator {
	return &jwtGenerator{signKey: signKey, tokenAge: tokenAge}
}

// GenerateAccessToken function will return accessToken
func (j *jwtGenerator) GenerateAccessToken(cl Claim) <-chan AccessTokenResponse {
	result := make(chan AccessTokenResponse)
	go func() {
		defer close(result)
		token := jwt.New(jwt.SigningMethodRS256)
		claims := make(jwt.MapClaims)
		claims["iss"] = cl.Issuer
		claims["aud"] = cl.Audience
		claims["exp"] = time.Now().Add(j.tokenAge).Unix()
		claims["iat"] = time.Now().Unix()
		claims["sub"] = cl.Subject
		token.Claims = claims

		tokenString, err := token.SignedString(j.signKey)
		if err != nil {
			result <- AccessTokenResponse{Error: err, AccessToken: nil}
			return
		}
		result <- AccessTokenResponse{Error: nil, AccessToken: &AccessToken{Token: fmt.Sprintf("Bearer %v", tokenString)}}
	}()
	return result
}
