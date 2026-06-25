package user

import (
	"context"
	"ewallet-transaction/external/proto/token_validation/tokenvalidation_v1"
	"fmt"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"time"
)

type UserServiceClient struct {
	Client  tokenvalidation_v1.TokenValidationClient
	Timeout time.Duration
}

func NewUserServiceClient(conn *grpc.ClientConn) *UserServiceClient {
	return &UserServiceClient{
		Client:  tokenvalidation_v1.NewTokenValidationClient(conn),
		Timeout: 5 * time.Second,
	}
}

func (u *UserServiceClient) ValidateToken(ctx context.Context, token string) (Token, error) {
	var resp Token

	req := &tokenvalidation_v1.TokenRequest{
		Token: token,
	}

	rpcCtx, cancel := context.WithTimeout(ctx, u.Timeout)
	defer cancel()

	response, err := u.Client.ValidateToken(rpcCtx, req)
	if err != nil {
		return resp, errors.Wrap(err, "failed to validate external")
	}

	if response.Message != "Success" {
		return resp, fmt.Errorf("got response error from ums: %s", response.Message)
	}

	resp.UserID = response.Data.UserId
	resp.Username = response.Data.Username
	resp.FullName = response.Data.FullName

	return resp, nil
}
