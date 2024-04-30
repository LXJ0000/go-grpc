package service

import (
	"context"
	"time"

	"github.com/LXJ0000/go-grpc/user/domain"
	"github.com/LXJ0000/go-grpc/user/util/token"
)

type userService struct {
	userRepo       domain.UserRepository
	contextTimeout time.Duration
}

func NewUserService(userRepo domain.UserRepository, contextTimeout time.Duration) domain.UserService {
	return &userService{
		userRepo:       userRepo,
		contextTimeout: contextTimeout,
	}
}

func (svc *userService) Regis(c context.Context, user *domain.User) error {
	ctx, cancel := context.WithTimeout(c, svc.contextTimeout)
	defer cancel()
	return svc.userRepo.Create(ctx, user)
}

func (svc *userService) GetByEmail(c context.Context, email string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, svc.contextTimeout)
	defer cancel()
	return svc.userRepo.GetByEmail(ctx, email)
}

func (svc *userService) GetByID(c context.Context, id int64) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, svc.contextTimeout)
	defer cancel()
	return svc.userRepo.GetByID(ctx, id)
}

// Login、Register And Refresh Token
func (svc *userService) ExtractIDFromToken(requestToken string, secret string) (int64, error) {
	return token.ExtractIDFromToken(requestToken, secret)
}

func (svc *userService) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	return token.CreateAccessToken(user, secret, expiry)
}

func (svc *userService) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return token.CreateRefreshToken(user, secret, expiry)
}
