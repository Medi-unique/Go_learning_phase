package usecases

import (
    "context"
    "task-manager/Domain"
    "task-manager/Repositories"
)

type UserUsecase struct {
    userRepo *repositories.UserRepository
}

func NewUserUsecase(userRepo *repositories.UserRepository) *UserUsecase {
    return &UserUsecase{
        userRepo: userRepo,
    }
}

func (u *UserUsecase) RegisterUser(ctx context.Context, user *domain.User) error {
    return u.userRepo.CreateUser(ctx, user)
}

func (u *UserUsecase) LoginUser(ctx context.Context, username, password string) (*domain.User, error) {
    user, err := u.userRepo.GetUserByUsername(ctx, username)
    if err != nil {
        return nil, err
    }
    // Add password verification logic here
    return user, nil
}