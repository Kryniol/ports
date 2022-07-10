package infrastructure

import (
	"context"
	"fmt"
	"sync"

	"github.com/Kryniol/ports/domain"
)

type inMemoryPortRepository struct {
	ports sync.Map
}

func NewInMemoryPortRepository() *inMemoryPortRepository {
	return &inMemoryPortRepository{}
}

func (r *inMemoryPortRepository) Save(_ context.Context, port domain.Port) error {
	r.ports.Store(port.ID, port)

	return nil
}

func (r *inMemoryPortRepository) Get(_ context.Context, id domain.PortID) (*domain.Port, error) {
	res, ok := r.ports.Load(id)
	if !ok {
		return nil, fmt.Errorf("%w: port with ID: %s does not exist", domain.ErrNotFound, id)
	}

	port := res.(domain.Port)

	return &port, nil
}
