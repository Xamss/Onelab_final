package service

import (
	"context"
	"registration/domain"
)

type Service interface {
	CreateAccount(ctx context.Context, u *domain.User) error
	Login(ctx context.Context, username, password string) (string, error)
	VerifyToken(token string) (int64, error)
}
