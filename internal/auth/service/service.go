package service

import (
	"context"

	"github.com/adopabianko/dbo-service/config"
	"github.com/adopabianko/dbo-service/internal/auth/dto"
	"github.com/adopabianko/dbo-service/internal/auth/repository"
)

type (
	IAuthService interface {
		Login(ctx context.Context, params dto.LoginRequest) (string, error)
		Register(ctx context.Context, params dto.RegisterRequest) error
	}

	service struct {
		repository repository.IAuthRepository
		config     *config.Provider
	}
)

func NewService(authRepo repository.IAuthRepository, config *config.Provider) IAuthService {
	return &service{
		repository: authRepo,
		config:     config,
	}
}
