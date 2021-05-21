package repositories

import (
	"gorm.io/gorm"
	"xws2021-nistagram/backend/models"
	"xws2021-nistagram/backend/util/auth"
)

type UserRepository interface {
	GetAllUsers() ([]models.User, error)
	CreateUser(user *models.User) error
	CheckPassword(data auth.Credentials) (error)
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) (UserRepository, error) {
	if db == nil {
		panic("UserRepository not created, gorm.DB is nil")
	}

	return &userRepository{ DB: db }, nil
}

func (repository *userRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User

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

func (repository *userRepository) CheckPassword(data auth.Credentials) error {
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

func (repository *userRepository) CreateUser(user *models.User) error {
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