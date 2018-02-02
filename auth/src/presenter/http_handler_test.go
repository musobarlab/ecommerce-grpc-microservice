package presenter

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/wuriyanto48/ecommerce-grpc-microservice/auth/src/model"
	"github.com/wuriyanto48/ecommerce-grpc-microservice/auth/src/query"
	"github.com/wuriyanto48/ecommerce-grpc-microservice/auth/src/token"
	"github.com/wuriyanto48/ecommerce-grpc-microservice/auth/src/usecase"
	"github.com/wuriyanto48/ecommerce-grpc-microservice/auth/src/usecase/mocks"
)

func generateIdentityAccessTokenResult() <-chan usecase.UseCaseResult {
	output := make(chan usecase.UseCaseResult)
	go func() {
		var accessToken token.AccessToken
		accessToken.Token = `Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJQV1QiLCJpYXQiOjE1MTIzMTg0MjMsImlzcyI6InB1cndva2VydG9kZXYuZ2l0aHViLmlvIiwic3ViIjoiMDAxIn0.GeEDHv82F8xp_98QQLiWxZ5aVBzZej0e-Ios8M9l0tYdrrTbdP3zxutSi5H7rxrd43PmlFi0pMMGtbVw64kPkBspCE3Kebbeersa8isn1zBejZO62mIgpRIGRhAJ_rphsxXYqOKKQlgj2ecI39dRR7IRJZdNYoTXXeBktUeUcDU`

		output <- usecase.UseCaseResult{Result: &accessToken}
	}()
	return output
}

func generateIdentityAccessTokenResultError() <-chan usecase.UseCaseResult {
	output := make(chan usecase.UseCaseResult)
	go func() {
		output <- usecase.UseCaseResult{Error: errors.New("Error")}
	}()
	return output
}

func generateIdentityAccessTokenResultErrorNoResult() <-chan usecase.UseCaseResult {
	output := make(chan usecase.UseCaseResult)
	go func() {
		output <- usecase.UseCaseResult{Result: "no result"}
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
		output <- query.QueryResult{Result: i}
	}()
	return output
}

func TestAuthHandler(t *testing.T) {

	var i model.Identity
	i.ID = "M1"
	i.Email = "wuriyanto48@yahoo.co.id"
	i.Password = "12345"

	t.Run("Test Password Credentials", func(t *testing.T) {

		mockUseCase := new(mocks.AuthUseCase)

		body, _ := json.Marshal(i)

		mockUseCase.On("GetAccessToken", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(generateIdentityAccessTokenResult())
		req := httptest.NewRequest("POST", "/auth?grant_type=password", strings.NewReader(string(body)))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		//signKey, _ := getPrivateKey(PrivateKey)

		httpAuthHandler := NewHttpHandler(mockUseCase)

		handler := http.Handler(httpAuthHandler.Auth())

		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.

		handler.ServeHTTP(rec, req)

		// Check the status code is what we expect.
		assert.Equal(t, http.StatusOK, rec.Code)

	})

	t.Run("Test Password Credentials Invalid Method", func(t *testing.T) {

		mockUseCase := new(mocks.AuthUseCase)

		body, _ := json.Marshal(i)

		mockUseCase.On("GetAccessToken", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(generateIdentityAccessTokenResult())
		req := httptest.NewRequest("GET", "/auth?grant_type=password", strings.NewReader(string(body)))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		//signKey, _ := getPrivateKey(PrivateKey)

		httpAuthHandler := NewHttpHandler(mockUseCase)

		handler := http.Handler(httpAuthHandler.Auth())

		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.

		handler.ServeHTTP(rec, req)

		// Check the status code is what we expect.
		assert.Equal(t, http.StatusMethodNotAllowed, rec.Code)

	})

	t.Run("Test Password Credentials Missing Grant Type", func(t *testing.T) {

		mockUseCase := new(mocks.AuthUseCase)

		body, _ := json.Marshal(i)

		mockUseCase.On("GetAccessToken", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(generateIdentityAccessTokenResult())
		req := httptest.NewRequest("POST", "/auth", strings.NewReader(string(body)))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		//signKey, _ := getPrivateKey(PrivateKey)

		httpAuthHandler := NewHttpHandler(mockUseCase)

		handler := http.Handler(httpAuthHandler.Auth())

		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.

		handler.ServeHTTP(rec, req)

		// Check the status code is what we expect.
		assert.Equal(t, http.StatusBadRequest, rec.Code)

	})

	t.Run("Test Password Credentials Invalid Grant Type", func(t *testing.T) {

		mockUseCase := new(mocks.AuthUseCase)

		body, _ := json.Marshal(i)

		mockUseCase.On("GetAccessToken", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(generateIdentityAccessTokenResult())
		req := httptest.NewRequest("POST", "/auth?grant_type=no", strings.NewReader(string(body)))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		//signKey, _ := getPrivateKey(PrivateKey)

		httpAuthHandler := NewHttpHandler(mockUseCase)

		handler := http.Handler(httpAuthHandler.Auth())

		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.

		handler.ServeHTTP(rec, req)

		// Check the status code is what we expect.
		assert.Equal(t, http.StatusBadRequest, rec.Code)

	})

	t.Run("Test Password Credentials Invalid Body", func(t *testing.T) {

		mockUseCase := new(mocks.AuthUseCase)

		mockUseCase.On("GetAccessToken", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(generateIdentityAccessTokenResult())
		req := httptest.NewRequest("POST", "/auth?grant_type=password", strings.NewReader(""))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		//signKey, _ := getPrivateKey(PrivateKey)

		httpAuthHandler := NewHttpHandler(mockUseCase)

		handler := http.Handler(httpAuthHandler.Auth())

		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.

		handler.ServeHTTP(rec, req)

		// Check the status code is what we expect.
		assert.Equal(t, http.StatusInternalServerError, rec.Code)

	})

	t.Run("Test Password Credentials Fail Generate Token", func(t *testing.T) {

		mockUseCase := new(mocks.AuthUseCase)

		body, _ := json.Marshal(i)

		mockUseCase.On("GetAccessToken", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(generateIdentityAccessTokenResultError())
		req := httptest.NewRequest("POST", "/auth?grant_type=password", strings.NewReader(string(body)))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		//signKey, _ := getPrivateKey(PrivateKey)

		httpAuthHandler := NewHttpHandler(mockUseCase)

		handler := http.Handler(httpAuthHandler.Auth())

		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.

		handler.ServeHTTP(rec, req)

		// Check the status code is what we expect.
		assert.Equal(t, http.StatusUnauthorized, rec.Code)

	})

	t.Run("Test Password Credentials Fail Generate Access Token With No Result", func(t *testing.T) {

		mockUseCase := new(mocks.AuthUseCase)

		body, _ := json.Marshal(i)

		mockUseCase.On("GetAccessToken", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(generateIdentityAccessTokenResultErrorNoResult())
		req := httptest.NewRequest("POST", "/auth?grant_type=password", strings.NewReader(string(body)))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		//signKey, _ := getPrivateKey(PrivateKey)

		httpAuthHandler := NewHttpHandler(mockUseCase)

		handler := http.Handler(httpAuthHandler.Auth())

		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.

		handler.ServeHTTP(rec, req)

		// Check the status code is what we expect.
		assert.Equal(t, http.StatusUnauthorized, rec.Code)

	})

}
