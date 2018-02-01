package middleware

import (
	"crypto/rsa"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

// Bearer this middleware function for verifying accessToken from Authorization Header
func Bearer(verifyKey *rsa.PublicKey, next http.Handler) http.Handler {

	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		accessToken := req.Header.Get("Authorization")
		if accessToken == "" {
			http.Error(res, "No Token Provided", http.StatusUnauthorized)
			return
		}
		tokenSlice := strings.Split(accessToken, " ")
		if len(tokenSlice) < 2 {
			http.Error(res, "Token is not valid", http.StatusUnauthorized)
			return
		}

		if tokenSlice[0] != "Bearer" {
			http.Error(res, "Token is not valid", http.StatusUnauthorized)
			return
		}
		tokenString := tokenSlice[1]
		token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
			return verifyKey, nil
		})

		if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
			memberID := claims.Subject
			req.Header.Add("MemberId", memberID)
			next.ServeHTTP(res, req)
		} else if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				http.Error(res, "Token is not valid", http.StatusUnauthorized)
				return
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				http.Error(res, "Token expired", http.StatusUnauthorized)
				return
			} else {
				http.Error(res, "Token is not valid", http.StatusUnauthorized)
				return
			}
		} else {
			http.Error(res, "Token is not valid", http.StatusUnauthorized)
			return
		}
	})
}
