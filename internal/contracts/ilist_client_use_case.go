package contracts

import (
	"context"

	"myapp/internal/ports/output"
)

type ListClientUseCase interface {
	Execute(ctx context.Context) (*output.ListClient, error)
}
