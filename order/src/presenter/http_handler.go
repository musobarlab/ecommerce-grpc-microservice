package presenter

import (
	"fmt"
	"net/http"

	serviceModel "github.com/wuriyanto48/ecommerce-grpc-microservice/order/src/services/model"
	"github.com/wuriyanto48/ecommerce-grpc-microservice/order/src/usecase"
	"github.com/wuriyanto48/ecommerce-grpc-microservice/order/utils"
)

// HttpOrderHandler model
type HttpOrderHandler struct {
	orderUseCase usecase.OrderUseCase
}

// NewHttpOrderHandler for initialise HttpOrderHandler model
func NewHttpOrderHandler(orderUseCase usecase.OrderUseCase) *HttpOrderHandler {
	return &HttpOrderHandler{orderUseCase: orderUseCase}
}

// Me http handler function, for get Member by its ID from Authorization
func (h *HttpOrderHandler) Me() http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

		if req.Method != "GET" {
			utils.JsonResponse(res, "Invalid Method", http.StatusMethodNotAllowed)
			return
		}

		memberID := req.Header.Get("MemberId")

		memberResult := <-h.orderUseCase.FindMemberByID(memberID)

		if memberResult.Error != nil {
			fmt.Println(memberResult.Error)
			utils.JsonResponse(res, "Member not found", http.StatusInternalServerError)
			return
		}

		member, ok := memberResult.Result.(serviceModel.Member)

		if !ok {
			utils.JsonResponse(res, "Result is not member", http.StatusInternalServerError)
			return
		}

		utils.JsonResponse(res, member, http.StatusOK)

	})
}
