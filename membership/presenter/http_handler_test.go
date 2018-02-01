package presenter

import (
	"crypto/rsa"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/wuriyanto48/ecommerce-grpc-microservice/membership/middleware"
	"github.com/wuriyanto48/ecommerce-grpc-microservice/membership/model"
	"github.com/wuriyanto48/ecommerce-grpc-microservice/membership/usecase"
	"github.com/wuriyanto48/ecommerce-grpc-microservice/membership/usecase/mocks"
)

func generateMeSuccessResult() <-chan usecase.UseCaseResult {
	output := make(chan usecase.UseCaseResult)
	go func() {
		exampleMember := model.NewMember()
		exampleMember.ID = "M1"
		exampleMember.FirstName = "Wuriyanto"
		exampleMember.LastName = "Musobar"
		exampleMember.Email = "wuriyanto48@yahoo.co.id"
		exampleMember.Password = "123456"
		exampleMember.PasswordSalt = "salt"
		exampleMember.BirthDate = time.Now()
		exampleMember.Version = 1
		exampleMember.CreatedAt = time.Now()
		exampleMember.UpdatedAt = time.Now()

		output <- usecase.UseCaseResult{Result: exampleMember}
	}()
	return output
}

func generateMeFailNotPointerResult() <-chan usecase.UseCaseResult {
	output := make(chan usecase.UseCaseResult)
	go func() {
		exampleMember := model.Member{}

		output <- usecase.UseCaseResult{Result: exampleMember}
	}()
	return output
}

func generateMeFailResult() <-chan usecase.UseCaseResult {
	output := make(chan usecase.UseCaseResult)
	go func() {

		output <- usecase.UseCaseResult{Error: errors.New("Error occured")}
	}()
	return output
}

func TestHttpMemberHandler(t *testing.T) {

	validAccessToken := `Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJQV1QiLCJpYXQiOjE1MTIzMTg0MjMsImlzcyI6InB1cndva2VydG9kZXYuZ2l0aHViLmlvIiwic3ViIjoiMDAxIn0.GeEDHv82F8xp_98QQLiWxZ5aVBzZej0e-Ios8M9l0tYdrrTbdP3zxutSi5H7rxrd43PmlFi0pMMGtbVw64kPkBspCE3Kebbeersa8isn1zBejZO62mIgpRIGRhAJ_rphsxXYqOKKQlgj2ecI39dRR7IRJZdNYoTXXeBktUeUcDU`
	//expiredToken := `Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJQV1QiLCJleHAiOjE1MTIzMjA1NzcsImlhdCI6MTUxMjMyMDU2NywiaXNzIjoicHVyd29rZXJ0b2Rldi5naXRodWIuaW8iLCJzdWIiOiIwMDEifQ.WLrCMOaQu0pb2XKdZOpZp81_9wJ7vAm3GKhA0h8lbioF5kqHVzmlNjnkJ1ayjYQMTNfYhhONX8IHOCCoMGsG5xibKG88_ICHQHbOtFOINaTIxpewmRbDlbrg8BA8aihXQD-L8a_Nm5ZYgpHAcCZ02m0MNSwHYwHBYYbHpnvqRIo`
	//noHeaderToken := `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJQV1QiLCJleHAiOjE1MTIzMjA1NzcsImlhdCI6MTUxMjMyMDU2NywiaXNzIjoicHVyd29rZXJ0b2Rldi5naXRodWIuaW8iLCJzdWIiOiIwMDEifQ.WLrCMOaQu0pb2XKdZOpZp81_9wJ7vAm3GKhA0h8lbioF5kqHVzmlNjnkJ1ayjYQMTNfYhhONX8IHOCCoMGsG5xibKG88_ICHQHbOtFOINaTIxpewmRbDlbrg8BA8aihXQD-L8a_Nm5ZYgpHAcCZ02m0MNSwHYwHBYYbHpnvqRIo`
	//invalidToken := `abcd3`
	t.Run("Test Me", func(t *testing.T) {

		mockUseCase := new(mocks.MemberUseCase)

		mockUseCase.On("FindByID", mock.AnythingOfType("string")).Return(generateMeSuccessResult())

		req := httptest.NewRequest("GET", "/me", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", validAccessToken)
		rec := httptest.NewRecorder()

		verifyKey, _ := getPublicKey(PublicKey)

		//member http handler
		httpMemberHandler := NewHttpMemberHandler(mockUseCase)

		handler := middleware.Bearer(verifyKey, httpMemberHandler.Me())
		handler.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("Test Me With Invalid HTTP Method", func(t *testing.T) {

		mockUseCase := new(mocks.MemberUseCase)

		mockUseCase.On("FindByID", mock.AnythingOfType("string")).Return(generateMeSuccessResult())

		req := httptest.NewRequest("POST", "/me", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", validAccessToken)
		rec := httptest.NewRecorder()

		verifyKey, _ := getPublicKey(PublicKey)

		//member http handler
		httpMemberHandler := NewHttpMemberHandler(mockUseCase)

		handler := middleware.Bearer(verifyKey, httpMemberHandler.Me())
		handler.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusMethodNotAllowed, rec.Code)
	})

	t.Run("Test Me With Use Case Error FindByID ", func(t *testing.T) {

		mockUseCase := new(mocks.MemberUseCase)

		mockUseCase.On("FindByID", mock.AnythingOfType("string")).Return(generateMeFailResult())

		req := httptest.NewRequest("GET", "/me", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", validAccessToken)
		rec := httptest.NewRecorder()

		verifyKey, _ := getPublicKey(PublicKey)

		//member http handler
		httpMemberHandler := NewHttpMemberHandler(mockUseCase)

		handler := middleware.Bearer(verifyKey, httpMemberHandler.Me())
		handler.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})

	t.Run("Test Me With Use Case Error FindByID Result Member Pointer ", func(t *testing.T) {

		mockUseCase := new(mocks.MemberUseCase)

		mockUseCase.On("FindByID", mock.AnythingOfType("string")).Return(generateMeFailNotPointerResult())

		req := httptest.NewRequest("GET", "/me", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", validAccessToken)
		rec := httptest.NewRecorder()

		verifyKey, _ := getPublicKey(PublicKey)

		//member http handler
		httpMemberHandler := NewHttpMemberHandler(mockUseCase)

		handler := middleware.Bearer(verifyKey, httpMemberHandler.Me())
		handler.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})

}

const PublicKey = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCoqzL5JrMzed4tb8uEoLKd42EO
sYmb0HpbicGt/OUeJxaHtt59Ew0BbpreBeiuugXweEa5xctQOxGYr27h4ZOnR0hW
Si+h5Y35CKzMEmZnzQwzQphgqww0U+e9/OAvVfCW1xWvVFr0WbhIRn+w/9DUvp+6
jKz3fIj3yQaHWVMMNQIDAQAB
-----END PUBLIC KEY-----
`

func getPublicKey(publicKey string) (*rsa.PublicKey, error) {
	r := strings.NewReader(publicKey)
	verifyBytes, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		return nil, err
	}
	return verifyKey, nil
}
