package repository

import (
	"context"

	"github.com/adopabianko/dbo-service/internal/auth/entity"

	"github.com/jackc/pgx/v5/pgxpool"
)

type (
	IAuthRepository interface {
		Login(ctx context.Context, email string) (entity.Auth, error)
		Register(ctx context.Context, params entity.Auth) (entity.Auth, error)
	}

	repository struct {
		db *pgxpool.Pool
	}
)

func NewRepository(db *pgxpool.Pool) IAuthRepository {
	return &repository{
		db: db,
	}
}
