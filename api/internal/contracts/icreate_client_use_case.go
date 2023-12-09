package contracts

import (
	"context"

	"myapp/api/internal/ports/input"
	"myapp/api/internal/ports/output"
)

type CreateClientUseCase interface {
	Execute(ctx context.Context, createClient *input.CreateClient) (*output.CreateClient, error)
}
