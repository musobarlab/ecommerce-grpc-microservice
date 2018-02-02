package presenter

import (
	"errors"

	pb "github.com/wuriyanto48/ecommerce-grpc-microservice/membership/protogo/membership"
	"github.com/wuriyanto48/ecommerce-grpc-microservice/membership/src/model"
	"github.com/wuriyanto48/ecommerce-grpc-microservice/membership/src/usecase"

	"golang.org/x/net/context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//GrpcHandler
type GrpcHandler struct {
	membershipUseCase usecase.MemberUseCase
}

// NewGrpcHandler
func NewGrpcHandler(membershipUseCase usecase.MemberUseCase) *GrpcHandler {
	return &GrpcHandler{membershipUseCase}
}

// Save
func (h *GrpcHandler) Save(ctx context.Context, data *pb.MembershipRequest) (*pb.MembershipResponse, error) {
	return nil, nil
}

// FindByID
func (h *GrpcHandler) FindByID(arg *pb.QueryRequest, stream pb.MembershipService_FindByIDServer) error {
	id := arg.ID

	memberResult := <-h.membershipUseCase.FindByID(id)

	if memberResult.Error != nil {
		return status.Error(codes.InvalidArgument, memberResult.Error.Error())
	}

	member, ok := memberResult.Result.(*model.Member)

	if !ok {
		err := errors.New("Result is Not Member")
		return status.Error(codes.InvalidArgument, err.Error())
	}

	response := &pb.MembershipResponse{
		ID:           member.ID,
		FirstName:    member.FirstName,
		LastName:     member.LastName,
		Email:        member.Email,
		Password:     member.Password,
		PasswordSalt: member.PasswordSalt,
		BirthDate:    member.BirthDate.String(),
		Version:      int32(member.Version),
		CreatedAt:    member.CreatedAt.String(),
		UpdatedAt:    member.UpdatedAt.String(),
	}

	if err := stream.Send(response); err != nil {
		status.Error(codes.Internal, err.Error())
	}

	return nil
}

// FindByEmail
func (h *GrpcHandler) FindByEmail(arg *pb.QueryRequest, stream pb.MembershipService_FindByEmailServer) error {
	email := arg.Email

	memberResult := <-h.membershipUseCase.FindByEmail(email)

	if memberResult.Error != nil {
		return status.Error(codes.InvalidArgument, memberResult.Error.Error())
	}

	member, ok := memberResult.Result.(*model.Member)

	if !ok {
		err := errors.New("Result is Not Member")
		return status.Error(codes.InvalidArgument, err.Error())
	}

	response := &pb.MembershipResponse{
		ID:           member.ID,
		FirstName:    member.FirstName,
		LastName:     member.LastName,
		Email:        member.Email,
		Password:     member.Password,
		PasswordSalt: member.PasswordSalt,
		BirthDate:    member.BirthDate.String(),
		Version:      int32(member.Version),
		CreatedAt:    member.CreatedAt.String(),
		UpdatedAt:    member.UpdatedAt.String(),
	}

	if err := stream.Send(response); err != nil {
		status.Error(codes.Internal, err.Error())
	}

	return nil
}
