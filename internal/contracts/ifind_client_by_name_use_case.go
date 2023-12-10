package contracts

import (
	"context"

	"myapp/internal/ports/output"
)

type FindClientByNameUseCase interface {
	Execute(ctx context.Context, name string) (*output.ListClient, error)
}
