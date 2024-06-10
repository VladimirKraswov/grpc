package auth

import "context"

type AuthService struct {
	
}

func New() *AuthService {
	return &AuthService{}
}

func (a AuthService) RegisterNewUser(ctx context.Context,  email string, password string) (userID int64, err error) {
	return 1, nil
}

func (a AuthService) Login(ctx context.Context,  email string, password string) (token string, err error) {
	return "token", nil
}