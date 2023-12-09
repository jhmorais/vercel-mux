package client

import (
	"context"
	"fmt"

	"myapp/api/internal/contracts"
	"myapp/api/internal/ports/output"
	"myapp/api/internal/repositories"
)

type deleteClientUseCase struct {
	clientRepository repositories.ClientRepository
}

func NewDeleteClientUseCase(clientRepository repositories.ClientRepository) contracts.DeleteClientUseCase {

	return &deleteClientUseCase{
		clientRepository: clientRepository,
	}
}

func (c *deleteClientUseCase) Execute(ctx context.Context, clientID int) (*output.DeleteClient, error) {

	clientEntity, err := c.clientRepository.FindClientByID(ctx, clientID)
	if err != nil {
		return nil, fmt.Errorf("failed to find client '%d' at database: '%w'", clientID, err)
	}

	if clientEntity == nil || clientEntity.ID == 0 {
		return nil, fmt.Errorf("clientID not found")
	}

	err = c.clientRepository.DeleteClient(ctx, clientEntity)
	if err != nil {
		return nil, fmt.Errorf("failed to delete client '%d'", clientEntity.ID)
	}

	output := &output.DeleteClient{
		ClientID:   clientEntity.ID,
		ClientName: clientEntity.Name,
	}

	return output, nil
}
