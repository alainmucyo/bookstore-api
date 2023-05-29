package books

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

func (s *Service) FindAll() ([]entities.Book, error) {
	books := make([]entities.Book, 0)
	if s.db.DB.Preload("Author").Find(&books).Error != nil {
		return nil, errors.New("error while getting books")
	}
	return books, nil
}

func (s *Service) FindById(id string) (entities.Book, error) {
	var book entities.Book
	if s.db.DB.Preload("Author").Where("id = ?", id).First(&book).Error != nil {
		return entities.Book{}, errors.New("book not found")
	}
	return book, nil
}

func (s *Service) Store(book entities.Book) (entities.Book, error) {
	ctx := s.db.DB.Begin()
	if ctx.Error != nil {
		return entities.Book{}, errors.New("error while creating book")
	}
	if err := ctx.Create(&book).Error; err != nil {
		ctx.Rollback()
		return entities.Book{}, errors.New("error while creating book")
	}
	ctx.Commit()
	return book, nil
}

func (s *Service) Update(id string, book entities.Book) (entities.Book, error) {
	ctx := s.db.DB.Begin()
	if ctx.Error != nil {
		return entities.Book{}, errors.New("error while updating book")
	}
	if err := ctx.Model(&book).Where("id = ?", id).Updates(book).Error; err != nil {
		ctx.Rollback()
		return entities.Book{}, errors.New("error while updating book")
	}
	ctx.Commit()
	return book, nil
}

func (s *Service) Delete(id string) error {
	ctx := s.db.DB.Begin()
	if ctx.Error != nil {
		return errors.New("error while deleting book")
	}
	if err := ctx.Where("id = ?", id).Delete(&entities.Book{}).Error; err != nil {
		ctx.Rollback()
		return errors.New("error while deleting book")
	}
	ctx.Commit()
	return nil
}
