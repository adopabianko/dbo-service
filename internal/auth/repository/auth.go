package repository

import (
	"context"
	"net/http"
	"strings"

	"github.com/adopabianko/dbo-service/internal/auth/entity"
	"github.com/adopabianko/dbo-service/pkg/stacktrace"
)

func (r *repository) Login(ctx context.Context, email string) (entity.Auth, error) {
	var opt entity.Auth
	err := r.db.QueryRow(ctx, `
	SELECT 
		email,
		password
	FROM auth WHERE email = $1`, email).Scan(
		&opt.Email,
		&opt.Password,
	)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return entity.Auth{}, stacktrace.WrapWithCode(err, http.StatusNotFound, "user not found")
		}
		return entity.Auth{}, stacktrace.WrapWithCode(err, http.StatusInternalServerError, "error when user login")
	}
	return opt, nil
}

func (r *repository) Register(ctx context.Context, auth entity.Auth) (entity.Auth, error) {
	query := `
	INSERT INTO auth (email, password, created_by)
	VALUES ($1, $2, $3)
	RETURNING id, email, created_at, created_by`

	err := r.db.QueryRow(ctx, query,
		auth.Email, auth.Password, auth.CreatedBy,
	).Scan(
		&auth.ID, &auth.Email,
		&auth.CreatedAt, &auth.CreatedBy,
	)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return entity.Auth{}, stacktrace.WrapWithCode(err, http.StatusConflict, "auth code already exists")
		}

		return entity.Auth{}, stacktrace.WrapWithCode(err, http.StatusInternalServerError, "error when create auth")
	}
	return auth, nil
}
