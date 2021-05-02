package repository

import (
	"context"

	"github.com/prigotti/cargo/portsservice/domain"
)

type portRepository struct {
}

func NewMongoDBPortRepository() domain.PortRepository {
	return &portRepository{}
}

func (r *portRepository) Save(ctx context.Context, p *domain.Port) error {
	return nil
}

func (r *portRepository) GetWithID(ctx context.Context, id string) (*domain.Port, error) {
	return nil, nil
}
