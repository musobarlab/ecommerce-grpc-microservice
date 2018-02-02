package middleware

import (
	"crypto/rsa"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

func TestBearer(t *testing.T) {

	validAccessToken := `Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJQV1QiLCJpYXQiOjE1MTIzMTg0MjMsImlzcyI6InB1cndva2VydG9kZXYuZ2l0aHViLmlvIiwic3ViIjoiMDAxIn0.GeEDHv82F8xp_98QQLiWxZ5aVBzZej0e-Ios8M9l0tYdrrTbdP3zxutSi5H7rxrd43PmlFi0pMMGtbVw64kPkBspCE3Kebbeersa8isn1zBejZO62mIgpRIGRhAJ_rphsxXYqOKKQlgj2ecI39dRR7IRJZdNYoTXXeBktUeUcDU`
	expiredToken := `Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJQV1QiLCJleHAiOjE1MTIzMjA1NzcsImlhdCI6MTUxMjMyMDU2NywiaXNzIjoicHVyd29rZXJ0b2Rldi5naXRodWIuaW8iLCJzdWIiOiIwMDEifQ.WLrCMOaQu0pb2XKdZOpZp81_9wJ7vAm3GKhA0h8lbioF5kqHVzmlNjnkJ1ayjYQMTNfYhhONX8IHOCCoMGsG5xibKG88_ICHQHbOtFOINaTIxpewmRbDlbrg8BA8aihXQD-L8a_Nm5ZYgpHAcCZ02m0MNSwHYwHBYYbHpnvqRIo`
	noHeaderToken := `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJQV1QiLCJleHAiOjE1MTIzMjA1NzcsImlhdCI6MTUxMjMyMDU2NywiaXNzIjoicHVyd29rZXJ0b2Rldi5naXRodWIuaW8iLCJzdWIiOiIwMDEifQ.WLrCMOaQu0pb2XKdZOpZp81_9wJ7vAm3GKhA0h8lbioF5kqHVzmlNjnkJ1ayjYQMTNfYhhONX8IHOCCoMGsG5xibKG88_ICHQHbOtFOINaTIxpewmRbDlbrg8BA8aihXQD-L8a_Nm5ZYgpHAcCZ02m0MNSwHYwHBYYbHpnvqRIo`
	invalidHeaderToken := `Basic eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJQV1QiLCJleHAiOjE1MTIzMjA1NzcsImlhdCI6MTUxMjMyMDU2NywiaXNzIjoicHVyd29rZXJ0b2Rldi5naXRodWIuaW8iLCJzdWIiOiIwMDEifQ.WLrCMOaQu0pb2XKdZOpZp81_9wJ7vAm3GKhA0h8lbioF5kqHVzmlNjnkJ1ayjYQMTNfYhhONX8IHOCCoMGsG5xibKG88_ICHQHbOtFOINaTIxpewmRbDlbrg8BA8aihXQD-L8a_Nm5ZYgpHAcCZ02m0MNSwHYwHBYYbHpnvqRIo`
	invalidToken := `abcd3`
	t.Run("Test Valid Access Token", func(t *testing.T) {

		req := httptest.NewRequest("GET", "/", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", validAccessToken)
		rec := httptest.NewRecorder()

		verifyKey, _ := getPublicKey(PublicKey)

		testHandler := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

			res.WriteHeader(http.StatusOK)
			res.Header().Set("Content-Type", "application/json")

			io.WriteString(res, `{"alive": true}`)
		})

		handler := Bearer(verifyKey, testHandler)
		handler.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "001", req.Header.Get("MemberId"))
	})

	t.Run("Test Expired Access Token", func(t *testing.T) {

		req := httptest.NewRequest("GET", "/", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", expiredToken)
		rec := httptest.NewRecorder()

		verifyKey, _ := getPublicKey(PublicKey)

		testHandler := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

			res.WriteHeader(http.StatusOK)
			res.Header().Set("Content-Type", "application/json")

			io.WriteString(res, `{"alive": true}`)
		})

		handler := Bearer(verifyKey, testHandler)
		handler.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusUnauthorized, rec.Code)
	})

	t.Run("Test Access Token without Header", func(t *testing.T) {

		req := httptest.NewRequest("GET", "/", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", noHeaderToken)
		rec := httptest.NewRecorder()

		verifyKey, _ := getPublicKey(PublicKey)

		testHandler := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

			res.WriteHeader(http.StatusOK)
			res.Header().Set("Content-Type", "application/json")

			io.WriteString(res, `{"alive": true}`)
		})

		handler := Bearer(verifyKey, testHandler)
		handler.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusUnauthorized, rec.Code)
	})

	t.Run("Test Access Token with Invalid Header", func(t *testing.T) {

		req := httptest.NewRequest("GET", "/", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", invalidHeaderToken)
		rec := httptest.NewRecorder()

		verifyKey, _ := getPublicKey(PublicKey)

		testHandler := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

			res.WriteHeader(http.StatusOK)
			res.Header().Set("Content-Type", "application/json")

			io.WriteString(res, `{"alive": true}`)
		})

		handler := Bearer(verifyKey, testHandler)
		handler.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusUnauthorized, rec.Code)
	})

	t.Run("Test Empty Access Token", func(t *testing.T) {

		req := httptest.NewRequest("GET", "/", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "")
		rec := httptest.NewRecorder()

		verifyKey, _ := getPublicKey(PublicKey)

		testHandler := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

			res.WriteHeader(http.StatusOK)
			res.Header().Set("Content-Type", "application/json")

			io.WriteString(res, `{"alive": true}`)
		})

		handler := Bearer(verifyKey, testHandler)
		handler.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusUnauthorized, rec.Code)
	})

	t.Run("Test Invalid Access Token", func(t *testing.T) {

		req := httptest.NewRequest("GET", "/", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", invalidToken)
		rec := httptest.NewRecorder()

		verifyKey, _ := getPublicKey(PublicKey)

		testHandler := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

			res.WriteHeader(http.StatusOK)
			res.Header().Set("Content-Type", "application/json")

			io.WriteString(res, `{"alive": true}`)
		})

		handler := Bearer(verifyKey, testHandler)
		handler.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusUnauthorized, rec.Code)
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
