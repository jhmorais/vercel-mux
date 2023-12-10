package contracts

import (
	"context"

	"myapp/internal/ports/output"
)

type FindClientByIDUseCase interface {
	Execute(ctx context.Context, clientID int) (*output.FindClient, error)
}
