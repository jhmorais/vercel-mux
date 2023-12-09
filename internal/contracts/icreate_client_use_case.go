package contracts

import (
	"context"

	"myapp/internal/ports/input"
	"myapp/internal/ports/output"
)

type CreateClientUseCase interface {
	Execute(ctx context.Context, createClient *input.CreateClient) (*output.CreateClient, error)
}
