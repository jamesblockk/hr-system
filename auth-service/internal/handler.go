package internal

import (
	"context"
	"hr-system/common/jwt"
	"hr-system/common/proto"
	"time"
)

type Server struct {
	proto.UnimplementedAuthServiceServer
}

func (s *Server) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	// token := "mock-jwt-token"
	token, err := jwt.Generate(map[string]interface{}{"user_id": 7}, 24*time.Hour)
	if err != nil {
		return nil, err
	}

	return &proto.LoginResponse{Token: token}, nil
}

func (s *Server) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	return &proto.RegisterResponse{Status: "User registered successfully"}, nil
}
