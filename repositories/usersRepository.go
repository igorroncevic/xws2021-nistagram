package repositories

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	uuid "github.com/satori/go.uuid"
	"xws2021-nistagram/models"
)

type UserRepository interface {
	GetAllUsers() ([]models.User, error)
	CreateUser(user *models.User) error
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

func (repository *userRepository) CreateUser(user *models.User) error {
	tx, err := repository.DB.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	user.ID = uuid.NewV4().String()

	query := "insert into registered_users (id, first_name, last_name) values ($1, $2, $3)"
	_, err = tx.Exec(context.Background(), query, user.ID, user.FirstName, user.LastName)

	if err != nil {
		return err
	}

	return tx.Commit(context.Background())
}