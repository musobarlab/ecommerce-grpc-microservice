package presenter

import (
	"encoding/json"
	"net/http"

	"github.com/wuriyanto48/ecommerce-grpc-microservice/auth/src/model"
	"github.com/wuriyanto48/ecommerce-grpc-microservice/auth/src/token"
	"github.com/wuriyanto48/ecommerce-grpc-microservice/auth/src/usecase"
	"github.com/wuriyanto48/ecommerce-grpc-microservice/auth/utils"
)

// HttpAuthHandler model
type HttpAuthHandler struct {
	AuthUseCase usecase.AuthUseCase
}

// NewHttpHandler function for initialise *HttpAuthHandler
func NewHttpHandler(authUseCase usecase.AuthUseCase) *HttpAuthHandler {
	return &HttpAuthHandler{AuthUseCase: authUseCase}
}

// Auth function, this http handler for user login and getting accessToken
func (h *HttpAuthHandler) Auth() http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

		if req.Method != "POST" {
			utils.JsonResponse(res, "Invalid Method", http.StatusMethodNotAllowed)
			return
		}

		grantTypes, ok := req.URL.Query()["grant_type"]

		if !ok || len(grantTypes) < 1 {
			utils.JsonResponse(res, "Missing Grant Type", http.StatusBadRequest)
			return
		}

		// Query()["grant_type"] will return an array of items,
		// we only want the single item.
		grantType := grantTypes[0]

		switch grantType {
		case "password":
			var identityLogin model.Identity

			decoder := json.NewDecoder(req.Body)
			err := decoder.Decode(&identityLogin)

			if err != nil {
				utils.JsonResponse(res, "Error occured", http.StatusInternalServerError)
				return
			}

			accessTokenResult := <-h.AuthUseCase.GetAccessToken(identityLogin.Email, identityLogin.Password)

			if accessTokenResult.Error != nil {
				utils.JsonResponse(res, "Invalid username or password", http.StatusUnauthorized)
				return
			}

			accessToken, ok := accessTokenResult.Result.(*token.AccessToken)

			if !ok {
				utils.JsonResponse(res, "Invalid username or password", http.StatusUnauthorized)
				return
			}

			utils.JsonResponse(res, accessToken, http.StatusOK)
			return
		default:
			utils.JsonResponse(res, "Invalid Grant Type", http.StatusBadRequest)
			return
		}
	})
}
