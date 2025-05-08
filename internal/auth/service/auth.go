package service

import (
	"context"
	"net/http"
	"time"

	"github.com/adopabianko/dbo-service/internal/auth/dto"
	"github.com/adopabianko/dbo-service/internal/auth/entity"
	"github.com/adopabianko/dbo-service/pkg/stacktrace"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	Email  string `json:"email"`
	UserID int    `json:"user_id"`
	jwt.StandardClaims
}

func (s *service) Login(ctx context.Context, params dto.LoginRequest) (string, error) {
	user, err := s.repository.Login(ctx, params.Email)
	if err != nil {
		return "", err
	}

	// verify password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password))
	if err != nil {
		return "", stacktrace.WrapWithCode(err, http.StatusUnauthorized, "error compare password")
	}

	// Set expiration time
	expirationTime := time.Now().Add(24 * time.Hour)

	// Buat claims
	claims := &Claims{
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Generate token using signing method HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token with secret key
	var jwtKey = []byte(s.config.JWT.JWTSecret)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", stacktrace.WrapWithCode(err, http.StatusUnauthorized, "error signing token")
	}

	return tokenString, nil
}

func (s *service) Register(ctx context.Context, params dto.RegisterRequest) error {
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		return stacktrace.WrapWithCode(err, http.StatusInternalServerError, "error hashing password")
	}

	var opt entity.Auth
	opt.Email = params.Email
	opt.Password = string(hashedPassword)
	opt.CreatedBy = params.CreatedBy

	_, err = s.repository.Register(ctx, opt)
	if err != nil {
		return err
	}

	return nil
}
