package client

import (
	"context"
	"fmt"

	"myapp/internal/contracts"
	"myapp/internal/ports/output"
	"myapp/internal/repositories"
)

type findClientByNameBrandUseCase struct {
	clientRepository repositories.ClientRepository
}

func NewFindClientByNameUseCase(clientRepository repositories.ClientRepository) contracts.FindClientByNameUseCase {

	return &findClientByNameBrandUseCase{
		clientRepository: clientRepository,
	}
}

func (c *findClientByNameBrandUseCase) Execute(ctx context.Context, name string) (*output.ListClient, error) {

	clientEntity, err := c.clientRepository.FindClientByName(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("erro to find client with brand: '%s' at database: '%v'", name, err)
	}

	if len(clientEntity) == 0 {
		return nil, fmt.Errorf("clients not found")
	}

	output := &output.ListClient{
		Clients: clientEntity,
	}

	return output, nil
}
