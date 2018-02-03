package grpc

import (
	"strconv"
	"time"

	"golang.org/x/net/context"

	pb "github.com/wuriyanto48/ecommerce-grpc-microservice/order/protogo/membership"
	"github.com/wuriyanto48/ecommerce-grpc-microservice/order/src/services/model"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

//membershipGrpcClientImpl struct
type membershipGrpcClientImpl struct {
	grpcAuthKey string
	client      pb.MembershipServiceClient
}

//NewMembershipGrpcClient
func NewMembershipGrpcClient(host, grpcAuthKey string) (*membershipGrpcClientImpl, error) {
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := pb.NewMembershipServiceClient(conn)

	return &membershipGrpcClientImpl{
		grpcAuthKey: grpcAuthKey,
		client:      client,
	}, nil
}

//FindByID
func (c *membershipGrpcClientImpl) FindByID(id string) <-chan ServiceResult {
	output := make(chan ServiceResult)
	go func() {
		md := metadata.Pairs("authorization", c.grpcAuthKey)
		ctx := metadata.NewOutgoingContext(context.Background(), md)
		arg := &pb.QueryRequest{ID: id}
		resStream, err := c.client.FindByID(ctx, arg)

		if err != nil {
			output <- ServiceResult{Error: err}
			return
		}

		res, err := resStream.Recv()

		if err != nil {
			output <- ServiceResult{Error: err}
			return
		}

		//birthDate from string to time
		bodInt, _ := strconv.Atoi(res.BirthDate)
		birthDate := time.Unix(int64(bodInt), int64(bodInt))

		if err != nil {
			output <- ServiceResult{Error: err}
			return
		}

		//createdAt from string to time
		createdAtInt, _ := strconv.Atoi(res.CreatedAt)
		createdAt := time.Unix(int64(bodInt), int64(createdAtInt))

		if err != nil {
			output <- ServiceResult{Error: err}
			return
		}

		//updatedAt from string to time
		updatedAtInt, _ := strconv.Atoi(res.UpdatedAt)
		updatedAt := time.Unix(int64(bodInt), int64(updatedAtInt))

		if err != nil {
			output <- ServiceResult{Error: err}
			return
		}

		member := model.Member{
			ID:           res.ID,
			FirstName:    res.FirstName,
			LastName:     res.LastName,
			Email:        res.Email,
			Password:     res.Password,
			PasswordSalt: res.PasswordSalt,
			BirthDate:    birthDate,
			Version:      int(res.Version),
			CreatedAt:    createdAt,
			UpdatedAt:    updatedAt,
		}

		output <- ServiceResult{Result: member}

	}()
	return output
}

//FindByEmail
func (c *membershipGrpcClientImpl) FindByEmail(email string) <-chan ServiceResult {
	output := make(chan ServiceResult)
	go func() {
		md := metadata.Pairs("authorization", c.grpcAuthKey)
		ctx := metadata.NewOutgoingContext(context.Background(), md)
		arg := &pb.QueryRequest{Email: email}
		resStream, err := c.client.FindByEmail(ctx, arg)

		if err != nil {
			output <- ServiceResult{Error: err}
			return
		}

		res, err := resStream.Recv()

		if err != nil {
			output <- ServiceResult{Error: err}
			return
		}

		birthDate, err := time.Parse(time.RFC3339, res.BirthDate)

		if err != nil {
			output <- ServiceResult{Error: err}
			return
		}

		createdAt, err := time.Parse(time.RFC3339, res.CreatedAt)

		if err != nil {
			output <- ServiceResult{Error: err}
			return
		}

		updatedAt, err := time.Parse(time.RFC3339, res.UpdatedAt)

		if err != nil {
			output <- ServiceResult{Error: err}
			return
		}

		member := model.Member{
			ID:           res.ID,
			FirstName:    res.FirstName,
			LastName:     res.LastName,
			Email:        res.Email,
			Password:     res.Password,
			PasswordSalt: res.PasswordSalt,
			BirthDate:    birthDate,
			Version:      int(res.Version),
			CreatedAt:    createdAt,
			UpdatedAt:    updatedAt,
		}

		output <- ServiceResult{Result: member}

	}()
	return output
}
