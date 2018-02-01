package usecase

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/wuriyanto48/ecommerce-grpc-microservice/auth/model"
	"github.com/wuriyanto48/ecommerce-grpc-microservice/auth/query"
	queryMocks "github.com/wuriyanto48/ecommerce-grpc-microservice/auth/query/mocks"
	"github.com/wuriyanto48/ecommerce-grpc-microservice/auth/token"
	atMocks "github.com/wuriyanto48/ecommerce-grpc-microservice/auth/token/mocks"
)

func generateAccessTokenResult() <-chan token.AccessTokenResponse {
	output := make(chan token.AccessTokenResponse)
	go func() {
		var at token.AccessToken
		at.Token = `Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJQV1QiLCJpYXQiOjE1MTIzMTg0MjMsImlzcyI6InB1cndva2VydG9kZXYuZ2l0aHViLmlvIiwic3ViIjoiMDAxIn0.GeEDHv82F8xp_98QQLiWxZ5aVBzZej0e-Ios8M9l0tYdrrTbdP3zxutSi5H7rxrd43PmlFi0pMMGtbVw64kPkBspCE3Kebbeersa8isn1zBejZO62mIgpRIGRhAJ_rphsxXYqOKKQlgj2ecI39dRR7IRJZdNYoTXXeBktUeUcDU`

		output <- token.AccessTokenResponse{AccessToken: &at}
	}()
	return output
}

func generateAccessTokenResultFail() <-chan token.AccessTokenResponse {
	output := make(chan token.AccessTokenResponse)
	go func() {
		output <- token.AccessTokenResponse{Error: errors.New("Failed to Generate Access Token")}
	}()
	return output
}

func generateQueryFindByEmailSuccessIdentityResult() <-chan query.QueryResult {
	output := make(chan query.QueryResult)
	go func() {
		var i model.Identity
		i.ID = "M1"
		i.Email = "wuriyanto48@yahoo.co.id"
		i.Password = "12345"
		output <- query.QueryResult{Result: &i}
	}()
	return output
}

func generateQueryFindByEmailFailIdentityResult() <-chan query.QueryResult {
	output := make(chan query.QueryResult)
	go func() {
		var i model.Identity
		i.ID = "M1"
		i.Email = "wuriyanto48@yahoo.co.id"
		i.Password = "12345"
		output <- query.QueryResult{Result: nil}
	}()
	return output
}

func TestAuthUseCase(t *testing.T) {

	t.Run("Test Generate Access Token", func(t *testing.T) {
		mockAccessToken := new(atMocks.AccessTokenGenerator)
		mockIdentityQuery := new(queryMocks.IdentityQuery)

		mockIdentityQuery.On("FindByEmail", mock.AnythingOfType("string")).Return(generateQueryFindByEmailSuccessIdentityResult())
		mockAccessToken.On("GenerateAccessToken", mock.AnythingOfType("token.Claim")).Return(generateAccessTokenResult())

		uc := NewAuthUseCase(mockIdentityQuery, mockAccessToken)

		accessTokenResult := <-uc.GetAccessToken("wuriyanto48@yahoo.co.id", "12345")

		assert.NoError(t, accessTokenResult.Error)
	})

	t.Run("Test Generate Access Token With Failed to Generate AccessToken", func(t *testing.T) {
		mockAccessToken := new(atMocks.AccessTokenGenerator)
		mockIdentityQuery := new(queryMocks.IdentityQuery)

		mockIdentityQuery.On("FindByEmail", mock.AnythingOfType("string")).Return(generateQueryFindByEmailSuccessIdentityResult())
		mockAccessToken.On("GenerateAccessToken", mock.AnythingOfType("token.Claim")).Return(generateAccessTokenResultFail())

		uc := NewAuthUseCase(mockIdentityQuery, mockAccessToken)

		accessTokenResult := <-uc.GetAccessToken("wuriyanto48@yahoo.co.id", "12345")

		assert.Error(t, accessTokenResult.Error)
	})

	t.Run("Test Generate Access Token With Invalid Email", func(t *testing.T) {
		mockAccessToken := new(atMocks.AccessTokenGenerator)
		mockIdentityQuery := new(queryMocks.IdentityQuery)

		mockIdentityQuery.On("FindByEmail", mock.AnythingOfType("string")).Return(generateQueryFindByEmailFailIdentityResult())
		mockAccessToken.On("GenerateAccessToken", mock.AnythingOfType("token.Claim")).Return(generateAccessTokenResult())

		uc := NewAuthUseCase(mockIdentityQuery, mockAccessToken)

		accessTokenResult := <-uc.GetAccessToken("wuriyanto48@yahoo.co.id", "12345")

		assert.Error(t, accessTokenResult.Error)
	})

	t.Run("Test Generate Access Token With Invalid Password", func(t *testing.T) {
		mockAccessToken := new(atMocks.AccessTokenGenerator)
		mockIdentityQuery := new(queryMocks.IdentityQuery)

		mockIdentityQuery.On("FindByEmail", mock.AnythingOfType("string")).Return(generateQueryFindByEmailSuccessIdentityResult())
		mockAccessToken.On("GenerateAccessToken", mock.AnythingOfType("token.Claim")).Return(generateAccessTokenResult())

		uc := NewAuthUseCase(mockIdentityQuery, mockAccessToken)

		accessTokenResult := <-uc.GetAccessToken("wuriyanto48@yahoo.co.id", "1234")

		assert.Error(t, accessTokenResult.Error)
	})
}
