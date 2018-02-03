package usecase

import (
	"errors"

	"github.com/wuriyanto48/ecommerce-grpc-microservice/auth/src/model"
	"github.com/wuriyanto48/ecommerce-grpc-microservice/auth/src/query"
	"github.com/wuriyanto48/ecommerce-grpc-microservice/auth/src/token"
)

// authUseCaseImpl model
type authUseCaseImpl struct {
	identityQuery        query.IdentityQuery
	accessTokenGenerator token.AccessTokenGenerator
}

// NewAuthUseCase function for initialise authUseCaseImpl model
func NewAuthUseCase(identityQuery query.IdentityQuery, accessTokenGenerator token.AccessTokenGenerator) AuthUseCase {
	return &authUseCaseImpl{
		identityQuery:        identityQuery,
		accessTokenGenerator: accessTokenGenerator,
	}
}

// GetAccessToken function will return accessToken that created from Identity ID
func (a *authUseCaseImpl) GetAccessToken(email, password string) <-chan UseCaseResult {
	output := make(chan UseCaseResult)

	go func() {
		defer close(output)

		identityResult := <-a.identityQuery.FindByEmail(email)

		if identityResult.Error != nil {
			output <- UseCaseResult{Error: identityResult.Error}
			return
		}

		identity, ok := identityResult.Result.(*model.Identity)

		if !ok {
			output <- UseCaseResult{Error: errors.New("Invalid Identity")}
			return
		}

		err := identity.IsValidPassword(password)

		if err != nil {
			output <- UseCaseResult{Error: err}
			return
		}

		var claims token.Claim
		claims.Issuer = "purwokertodev.github.io"
		claims.Subject = identity.ID
		claims.Audience = "PWT"

		accessTokenResult := <-a.accessTokenGenerator.GenerateAccessToken(claims)

		if accessTokenResult.Error != nil {
			output <- UseCaseResult{Error: errors.New("Cannot generate access token")}
			return
		}

		accessToken := accessTokenResult.AccessToken

		output <- UseCaseResult{Result: accessToken}

	}()

	return output
}
