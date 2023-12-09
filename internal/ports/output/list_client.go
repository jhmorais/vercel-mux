package output

import (
	"myapp/internal/domain/entities"
)

type ListClient struct {
	Clients []*entities.Client
}
