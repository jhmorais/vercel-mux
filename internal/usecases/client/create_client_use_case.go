package client

import (
	"context"
	"fmt"
	"time"

	"myapp/internal/contracts"
	"myapp/internal/domain/entities"
	"myapp/internal/ports/input"
	"myapp/internal/ports/output"
	"myapp/internal/repositories"
)

type createClientUseCase struct {
	clientRepository repositories.ClientRepository
}

func NewCreateClientUseCase(clientRepository repositories.ClientRepository) contracts.CreateClientUseCase {

	return &createClientUseCase{
		clientRepository: clientRepository,
	}
}

func (c *createClientUseCase) Execute(ctx context.Context, createClient *input.CreateClient) (*output.CreateClient, error) {

	//max 250
	if len(createClient.Name) > 250 {
		//will discard the rest
		createClient.Name = createClient.Name[:250]
	}

	if createClient.Name == "" {
		return nil, fmt.Errorf("cannot create a client without name")
	}

	if createClient.Name == "" {
		return nil, fmt.Errorf("cannot create a client without brand")
	}

	client, err := c.clientRepository.FindClientByName(ctx, createClient.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to get client")
	}

	if len(client) > 0 && client[0].ID > 0 {
		return nil, fmt.Errorf("failed, already exists client with the same name")
	}

	clientEntity := &entities.Client{
		Name:      createClient.Name,
		PixType:   createClient.PixType,
		PixKey:    createClient.PixKey,
		PartnerID: createClient.PartnerID,
		CreatedAt: time.Now(),
	}

	err = c.clientRepository.CreateClient(ctx, clientEntity)
	if err != nil {
		return nil, fmt.Errorf("cannot save client at database: %v", err)
	}

	createClientOutput := &output.CreateClient{
		ClientID: clientEntity.ID,
	}

	return createClientOutput, nil
}
