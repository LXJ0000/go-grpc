package service

import (
	"context"

	"github.com/LXJ0000/go-grpc/user/domain"
)

type userService struct {
	userRepo domain.UserRepository
}

func NewUserService(userRepo domain.UserRepository) domain.UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (svc *userService) Create(c context.Context, user *domain.User) error {
	return nil
}
func (svc *userService) GetByEmail(c context.Context, email string) (domain.User, error) {
	return domain.User{}, nil
}
func (svc *userService) GetByID(c context.Context, id int64) (domain.User, error) {
	return domain.User{}, nil

}
func (svc *userService) GetProfileByID(c context.Context, userID int64) (domain.UserProfile, error) {
	return domain.UserProfile{}, nil

}

// Login、Register And Refresh Token
func (svc *userService) ExtractIDFromToken(requestToken string, secret string) (int64, error) {
	return 0, nil
}
func (svc *userService) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	return "", nil
}
func (svc *userService) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return "", nil
}
