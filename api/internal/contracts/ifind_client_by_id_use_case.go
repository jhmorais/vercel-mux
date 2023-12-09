package contracts

import (
	"context"

	"myapp/api/internal/ports/output"
)

type FindClientByIDUseCase interface {
	Execute(ctx context.Context, clientID int) (*output.FindClient, error)
}
