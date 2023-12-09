package client

import (
	"context"
	"fmt"

	"myapp/api/internal/contracts"
	"myapp/api/internal/ports/output"
	"myapp/api/internal/repositories"
)

type findClientByIDUseCase struct {
	clientRepository repositories.ClientRepository
}

func NewFindClientByIDUseCase(clientRepository repositories.ClientRepository) contracts.FindClientByIDUseCase {

	return &findClientByIDUseCase{
		clientRepository: clientRepository,
	}
}

func (c *findClientByIDUseCase) Execute(ctx context.Context, clientID int) (*output.FindClient, error) {

	clientEntity, err := c.clientRepository.FindClientByID(ctx, clientID)
	if err != nil {
		return nil, fmt.Errorf("erro to find client '%d' at database: '%v'", clientID, err)
	}

	if clientEntity == nil || clientEntity.ID == 0 {
		return nil, fmt.Errorf("clientID not found")
	}

	output := &output.FindClient{
		Client: clientEntity,
	}

	return output, nil
}
