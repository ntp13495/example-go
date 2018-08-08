package lendingabook

import (
	"context"

	"github.com/jinzhu/gorm"

	"github.com/ntp13495/example-go/domain"
)

// pgService implmenter for LendingBooks serivce in postgres
type pgService struct {
	db *gorm.DB
}

// NewPGService create new PGService
func NewPGService(db *gorm.DB) Service {
	return &pgService{
		db: db,
	}
}

// Create implement Create for LendingBooks service
func (s *pgService) Create(_ context.Context, p *domain.LendingBooks) error {
	res := []domain.LendingBooks{}
	s.db.Find(&res)
	for _, iterator := range res {
		if p.BookID == iterator.BookID && p.From.Before(iterator.To) {
			return ErrBookInUse
		}
	}

	resBook := []domain.Book{}
	s.db.Find(&resBook)
	flag := 0
	for _, iterator := range resBook {
		if p.BookID == iterator.ID {
			flag = 1
		}
	}
	if flag == 0 {
		return ErrInvalidBook
	}

	resUser := []domain.User{}
	s.db.Find(&resUser)
	flag = 0
	for _, iterator := range resUser {
		if p.UserID == iterator.ID {
			flag = 1
		}
	}
	if flag == 0 {
		return ErrInvalidUser
	}

	return s.db.Create(p).Error
}

// Update implement Update for LendingBooks service
func (s *pgService) Update(_ context.Context, p *domain.LendingBooks) (*domain.LendingBooks, error) {
	old := domain.LendingBooks{Model: domain.Model{ID: p.ID}}
	if err := s.db.Find(&old).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}

	resBook := []domain.Book{}
	s.db.Find(&resBook)
	flag := 0
	for _, iterator := range resBook {
		if p.BookID == iterator.ID {
			flag = 1
		}
	}
	if flag == 0 {
		return nil, ErrInvalidBook
	}

	resUser := []domain.User{}
	s.db.Find(&resUser)
	flag = 0
	for _, iterator := range resUser {
		if p.UserID == iterator.ID {
			flag = 1
		}
	}
	if flag == 0 {
		return nil, ErrInvalidUser
	}

	old.BookID = p.BookID
	old.UserID = p.UserID
	old.From = p.From
	old.To = p.To

	return &old, s.db.Save(&old).Error
}

// Find implement Find for LendingBooks service
func (s *pgService) Find(_ context.Context, p *domain.LendingBooks) (*domain.LendingBooks, error) {
	res := p
	if err := s.db.Find(&res).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return res, nil
}

// FindAll implement FindAll for LendingBooks service
func (s *pgService) FindAll(_ context.Context) ([]domain.LendingBooks, error) {
	res := []domain.LendingBooks{}
	return res, s.db.Find(&res).Error
}

// Delete implement Delete for LendingBooks service
func (s *pgService) Delete(_ context.Context, p *domain.LendingBooks) error {
	old := domain.LendingBooks{Model: domain.Model{ID: p.ID}}
	if err := s.db.Find(&old).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrNotFound
		}
		return err
	}
	return s.db.Delete(old).Error
}
