package authors

import (
	"bookstore/store/entities"
	"bookstore/store/postgres"
	"errors"
)

type Service struct {
	db *postgres.Database
}

func New(db *postgres.Database) *Service {
	return &Service{db: db}
}

func (s *Service) FindAll() ([]entities.Author, error) {
	authors := make([]entities.Author, 0)
	if s.db.DB.Find(&authors).Error != nil {
		return nil, errors.New("error while getting authors")
	}
	return authors, nil
}

func (s *Service) FindById(id string) (entities.Author, error) {
	var author entities.Author
	if s.db.DB.Preload("Books").Where("id = ?", id).First(&author).Error != nil {
		return entities.Author{}, errors.New("author not found")
	}
	return author, nil
}

func (s *Service) Store(author entities.Author) (entities.Author, error) {
	ctx := s.db.DB.Begin()
	if ctx.Error != nil {
		return entities.Author{}, errors.New("error while creating author")
	}
	if err := ctx.Create(&author).Error; err != nil {
		ctx.Rollback()
		return entities.Author{}, errors.New("error while creating author")
	}
	ctx.Commit()
	return author, nil
}

func (s *Service) Update(id string, author entities.Author) (entities.Author, error) {
	ctx := s.db.DB.Begin()
	if ctx.Error != nil {
		return entities.Author{}, errors.New("error while updating author")
	}
	if err := ctx.Model(&author).Where("id = ?", id).Updates(author).Error; err != nil {
		ctx.Rollback()
		return entities.Author{}, errors.New("error while updating author")
	}
	ctx.Commit()
	return author, nil
}

func (s *Service) Delete(id string) error {
	ctx := s.db.DB.Begin()
	if ctx.Error != nil {
		return errors.New("error while deleting author")
	}
	if err := ctx.Where("id = ?", id).Delete(&entities.Author{}).Error; err != nil {
		ctx.Rollback()
		return errors.New("error while deleting author")
	}
	ctx.Commit()
	return nil
}
