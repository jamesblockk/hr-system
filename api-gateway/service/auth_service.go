package service

import (
	"context"
	"hr-system/api-gateway/client"
)

type AuthService struct {
	authClient *client.AuthServiceClient
}

func NewAuthService(authClient *client.AuthServiceClient) *AuthService {
	return &AuthService{authClient: authClient}
}

func (s *AuthService) Login(ctx context.Context, email, password string) (string, error) {
	return s.authClient.Login(ctx, email, password)
}

func (s *AuthService) Register(ctx context.Context, email, password string) error {
	return s.authClient.Register(ctx, email, password)
}
