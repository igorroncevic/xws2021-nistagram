package repositories

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"xws2021-nistagram/models"
)

type UserRepository interface {
	GetAllUsers() ([]models.User, error)
}

type userRepository struct {
	DB *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) UserRepository {
	if db == nil {
		panic("UserRepository not created, pgxpool is nil")
	}
	return &userRepository{
		DB: db,
	}
}

func (repository *userRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User

	query := "select u.id, u.first_name, u.last_name from registered_users u"
	rows, err := repository.DB.Query(context.Background(), query)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
