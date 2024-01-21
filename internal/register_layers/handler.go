package register_layers

import (
	"Test-Task2/internal/delivery/auth"
	"Test-Task2/internal/delivery/user"
	"Test-Task2/pkg/api/jwt"
	ssov1 "Test-Task2/proto"
	"google.golang.org/grpc"
)

type GDelivery struct {
	user *user.ServerUser
	auth *auth.ServerAuth
}

func NewGDelivery(uc *GUsecase, jwt *jwt.JWT) *GDelivery {
	return &GDelivery{
		user: user.Register(jwt, uc.user),
		auth: auth.Register(jwt, uc.auth),
	}
}

func (h *GDelivery) RegisterRoutes(gRPC *grpc.Server) {
	ssov1.RegisterUserServiceServer(gRPC, h.user)
	ssov1.RegisterAuthServer(gRPC, h.auth)
}
