package users

import (
	"bookstore/store/entities"
	"bookstore/store/hash"
	"bookstore/store/postgres"
	"errors"
)

type Service struct {
	db   *postgres.Database
	hash *hash.Hash
}

func New(db *postgres.Database, hash *hash.Hash) *Service {
	return &Service{db: db, hash: hash}
}

func (s *Service) FindAll() ([]entities.User, error) {
	users := make([]entities.User, 0)
	if s.db.DB.Find(&users).Error != nil {
		return nil, errors.New("error while getting users")
	}
	return users, nil
}

func (s *Service) VerifyUser(email string, password string) (bool, entities.User) {
	user, err := s.FindByEmail(email)

	if err != nil {
		return false, entities.User{}
	}
	return s.hash.VerifyPassword(password, user.Password), user
}

func (s *Service) FindByEmail(email string) (entities.User, error) {
	var user entities.User
	if s.db.DB.Where("email = ?", email).First(&user).Error != nil {
		return entities.User{}, errors.New("user not found")
	}
	return user, nil
}

func (s *Service) Store(user entities.User) (entities.User, error) {
	_, err := s.FindByEmail(user.Email)

	if err == nil {
		return entities.User{}, errors.New("email already used")
	}

	// Hash password
	hashedPassword, err := s.hash.HashPassword(user.Password)
	if err != nil {
		return entities.User{}, errors.New("password error")
	}
	user.Password = hashedPassword

	// save user
	ctx := s.db.DB.Begin()
	if ctx.Error != nil {
		return user, errors.New("error start")
	}
	if s.db.DB.Save(&user).Error != nil {
		ctx.Rollback()
		return user, errors.New("error save")
	}
	if ctx.Commit().Error != nil {
		ctx.Rollback()
		return user, errors.New("error commit")
	}
	return user, nil
}
