package repository

import (
	"database/sql"
	"native-go-clean-arch/internal/domain/model"
)

type UserRepository interface {
	FetchAll() ([]model.User, error)
}

type userRepositoryImpl struct {
	db *sql.DB
}

// Constructor
func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepositoryImpl{db: db}
}

// Implementasi FetchAll
func (r *userRepositoryImpl) FetchAll() ([]model.User, error) {
	rows, err := r.db.Query("SELECT id,account_no FROM partner")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.AccountNo); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
