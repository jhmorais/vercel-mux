package contracts

import (
	"context"

	"myapp/api/internal/ports/output"
)

type DeleteClientUseCase interface {
	Execute(ctx context.Context, clientID int) (*output.DeleteClient, error)
}
