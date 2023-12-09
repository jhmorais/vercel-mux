package output

import (
	"myapp/api/internal/domain/entities"
)

type ListClient struct {
	Clients []*entities.Client
}
