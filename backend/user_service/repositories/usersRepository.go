package repositories

import (
	"context"
	"github.com/david-drvar/xws2021-nistagram/common"
	"github.com/david-drvar/xws2021-nistagram/user_service/model"
	userspb "github.com/david-drvar/xws2021-nistagram/user_service/proto"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUsers() ([]model.User, error)
	CreateUser(context.Context, *userspb.User) error
	CheckPassword(data common.Credentials) error
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) (*userRepository, error) {
	if db == nil {
		panic("UserRepository not created, gorm.DB is nil")
	}

	return &userRepository{ DB: db }, nil
}

func (repository *userRepository) GetAllUsers() ([]model.User, error) {
	var users []model.User

	/*query := "select u.id, u.first_name, u.last_name, u.email from registered_users u"
	rows, err := repository.DB.Query(context.Background(), query)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}*/

	return users, nil
}

func (repository *userRepository) CheckPassword(data common.Credentials) error {
	/*var user models.User

	query := "select u.id, u.password from registered_users u where u.email = $1"
	rows, err := repository.DB.Query(context.Background(), query, data.Email)
	defer rows.Close()
	if err != nil {
		return err
	}

	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Password)
		if err != nil {
			return err
		}
	}

	err = encryption.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if err != nil{
		return err
	}
	*/

	return nil
}

func (repository *userRepository) CreateUser(ctx context.Context, user *userspb.User) error {
	/*tx, err := repository.DB.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	user.ID = uuid.NewV4().String()

	query := `insert into registered_users (id, first_name, last_name, "email", "password") values ($1, $2, $3, $4, $5)`
	_, err = tx.Exec(context.Background(), query, user.ID, user.FirstName, user.LastName, user.Email, user.Password)

	if err != nil {
		return err
	}

	return tx.Commit(context.Background())*/
	return nil
}