package contracts

import (
	"context"

	"myapp/internal/ports/output"
)

type FindClientUseCase interface {
	Execute(ctx context.Context, partnerID int, name string) (*output.FindClient, error)
}
