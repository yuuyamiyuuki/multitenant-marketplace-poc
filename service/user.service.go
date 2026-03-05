package service

import (
	"github.com/google/uuid"

	"gin-tenant/dto"
	"gin-tenant/middleware"
	"gin-tenant/repository"
	"gin-tenant/utils"
)

type UserService interface {
	CreateTenantUser(input dto.CreateUserInput, tenantID uuid.UUID) (*repository.User, error)
	UpdatePassword(userID uuid.UUID, tenantID uuid.UUID, input dto.UpdatePasswordInput) error
	Login(input dto.LoginInput) (*dto.LoginResponse, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{userRepo: repo}
}

func (s *userService) CreateTenantUser(input dto.CreateUserInput, tenantID uuid.UUID) (*repository.User, error) {
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return nil, middleware.NewInternal("failed to hash password", err)
	}

	newUser := &repository.User{
		TenantID: tenantID,
		Role:     input.Role,
		Username: input.Username,
		Password: hashedPassword,
	}

	err = s.userRepo.Create(newUser)
	if err != nil {
		return nil, middleware.NewInternal("failed to create user in database", err)
	}

	return newUser, nil
}

func (s *userService) UpdatePassword(userID uuid.UUID, tenantID uuid.UUID, input dto.UpdatePasswordInput) error {
	user, err := s.userRepo.FindByIDAndTenant(userID, tenantID)
	if err != nil {
		return middleware.NewNotFound("user not found", err)
	}

	if !utils.CheckPasswordHash(input.OldPassword, user.Password) {
		return middleware.NewBadRequest("incorrect old password", nil)
	}

	newHashedPassword, err := utils.HashPassword(input.NewPassword)
	if err != nil {
		return middleware.NewInternal("failed to hash new password", err)
	}

	user.Password = newHashedPassword
	err = s.userRepo.Update(user)
	if err != nil {
		return middleware.NewInternal("failed to update password in database", err)
	}

	return nil
}

func (s *userService) Login(input dto.LoginInput) (*dto.LoginResponse, error) {
	user, err := s.userRepo.FindByUsername(input.Username)
	if err != nil {
		return nil, middleware.NewUnauthorized("invalid username or password", err)
	}

	if !utils.CheckPasswordHash(input.Password, user.Password) {
		return nil, middleware.NewUnauthorized("invalid username or password", nil)
	}

	token, expiresAt, err := utils.GenerateToken(user.ID, user.TenantID, user.Role)
	if err != nil {
		return nil, middleware.NewInternal("failed to generate token", err)
	}

	return &dto.LoginResponse{
		Token:     token,
		ExpiresAt: expiresAt,
	}, nil
}
