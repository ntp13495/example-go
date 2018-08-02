package lendingabook

import (
	"context"

	"github.com/ntp13495/example-go/domain"
)

// Service interface for project service
type Service interface {
	Create(ctx context.Context, p *domain.LendingBooks) error
	Update(ctx context.Context, p *domain.LendingBooks) (*domain.LendingBooks, error)
	Find(ctx context.Context, p *domain.LendingBooks) (*domain.LendingBooks, error)
	FindAll(ctx context.Context) ([]domain.LendingBooks, error)
	Delete(ctx context.Context, p *domain.LendingBooks) error
}
