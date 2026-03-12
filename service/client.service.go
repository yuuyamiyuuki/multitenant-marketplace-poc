package service

import (
	"github.com/google/uuid"

	"gin-tenant/dto"
	"gin-tenant/middleware"
	"gin-tenant/repository"
)

type ClientService interface {
	CreateClient(input dto.CreateClientInput, tenantID uuid.UUID) (*repository.Client, error)
}

type clientService struct {
	clientRepo  repository.ClientRepository
	userService UserService
}

func NewClientService(cRepo repository.ClientRepository, uSvc UserService) ClientService {
	return &clientService{clientRepo: cRepo, userService: uSvc}
}

func (s *clientService) CreateClient(input dto.CreateClientInput, tenantID uuid.UUID) (*repository.Client, error) {
	userInput := dto.CreateUserInput{Username: input.Username, Role: "customer", Password: input.Password}
	newUser, err := s.userService.CreateTenantUser(userInput, tenantID)
	if err != nil {
		return nil, err
	}

	newClient := &repository.Client{
		Name:     input.Name,
		Birthday: input.Birthday,
		Document: input.Document,
		Phone:    input.Phone,
		UserId:   newUser.ID,
	}

	err = s.clientRepo.Create(newClient)
	if err != nil {
		return nil, middleware.NewInternal("failed to create client in database", err)
	}

	return newClient, nil
}
