package auth

import (
	"context"
	grpcv1 "grpc/internal/protos/grpc.v1"

	"google.golang.org/grpc"
)

type Auth interface {
	RegisterNewUser(ctx context.Context,  email string, password string) (userID int64, err error)
	Login(ctx context.Context,  email string, password string) (token string, err error)
}

type serverAPI struct {
	grpcv1.UnimplementedAuthServer
	auth Auth
}

func Register(grpcServer *grpc.Server, auth Auth) {
	grpcv1.RegisterAuthServer(grpcServer, serverAPI{auth: auth})
}

func (s serverAPI) Register(ctx context.Context, req *grpcv1.RegisterRequest) (*grpcv1.RegisterResponse, error) {
	uid := 777
	return &grpcv1.RegisterResponse{UserID: int64(uid)}, nil
}

func (s serverAPI) Login(ctx context.Context, req *grpcv1.LoginRequest) (*grpcv1.LoginResponse, error) {
	token := "hello"
	return &grpcv1.LoginResponse{Token: token}, nil
}