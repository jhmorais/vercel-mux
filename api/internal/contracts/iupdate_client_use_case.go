package contracts

import (
	"context"

	"myapp/api/internal/ports/input"
	"myapp/api/internal/ports/output"
)

type UpdateClientUseCase interface {
	Execute(ctx context.Context, updateClient *input.UpdateClient) (*output.CreateClient, error)
}
