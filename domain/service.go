package domain

import (
	"context"
	"errors"
	"fmt"
)

type PortRepository interface {
	Save(ctx context.Context, port Port) error
	Get(ctx context.Context, id PortID) (*Port, error)
}

type PortReader interface {
	Read(ctx context.Context) (<-chan Port, error)
}

type PortDomainService interface {
	SavePorts(ctx context.Context) error
}

type portService struct {
	reader PortReader
	repo   PortRepository
}

func NewPortService(reader PortReader, repo PortRepository) *portService {
	return &portService{
		reader: reader,
		repo:   repo,
	}
}

func (p *portService) SavePorts(ctx context.Context) error {
	ports, err := p.reader.Read(ctx)
	if err != nil {
		return fmt.Errorf("couldn't read ports from the reader: %w", err)
	}

	for {
		select {
		case port, ok := <-ports:
			if !ok {
				return nil
			}

			err := p.repo.Save(ctx, port)
			if err != nil {
				return fmt.Errorf("got an error when saving a port to the repository: %w", err)
			}
		case <-ctx.Done():
			return errors.New("context is done")
		}
	}
}
