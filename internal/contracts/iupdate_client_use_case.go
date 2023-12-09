package contracts

import (
	"context"

	"myapp/internal/ports/input"
	"myapp/internal/ports/output"
)

type UpdateClientUseCase interface {
	Execute(ctx context.Context, updateClient *input.UpdateClient) (*output.CreateClient, error)
}
