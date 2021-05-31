package repositories

import (
	"context"
	"errors"
	"fmt"
	"github.com/david-drvar/xws2021-nistagram/common"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/user_service/model/domain"
	"github.com/david-drvar/xws2021-nistagram/user_service/model/persistence"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUsers(context.Context) ([]persistence.User, error)
	CreateUser(context.Context, *persistence.User) error
	CheckPassword(data common.Credentials) error
	UpdateUserProfile(dto domain.User) (bool, error)
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


func (repository *userRepository) UpdateUserProfile(userDTO domain.User) (bool, error) {
	var user persistence.User
	var userAdditionalInfo persistence.UserAdditionalInfo

	db := repository.DB.Model(&user).Where("id = ?", userDTO.Id ).Updates(persistence.User{FirstName: userDTO.FirstName, LastName: userDTO.LastName, Email: userDTO.Email, Username: userDTO.Username, BirthDate: userDTO.BirthDate,
	    	 PhoneNumber: userDTO.PhoneNumber, Sex: userDTO.Sex})

	fmt.Println(db.RowsAffected)

	if db.Error != nil  {
		return false, db.Error
	}else if db.RowsAffected == 0 {
		return false, errors.New("rows affected is equal to zero")
	}
	db = repository.DB.Model(&userAdditionalInfo).Where("id = ?", userDTO.Id ).Updates(persistence.UserAdditionalInfo{Website: userDTO.Website, Category: userDTO.Category, Biography: userDTO.Biography}).Find(1)

	if db.Error != nil {
		return false, db.Error
	}else if db.RowsAffected == 0 {
		return false, errors.New("rows affected is equal to zero")
	}

	return true, nil
}

func(repository *userRepository) GetUserByUsername(username string) (domain.User, error) {
	var dbUser persistence.User
	var dbUserAdditionalInfo persistence.UserAdditionalInfo

	db := repository.DB.Where("username = ?", username).Find(&dbUser)
	if db.Error != nil{
		return domain.User{}, db.Error
	}

	db = repository.DB.Where("id = ?", dbUser.Id).Find(&dbUserAdditionalInfo)
	if db.Error != nil{
		return domain.User{}, db.Error
	}
	user := &domain.User{}

	user.GenerateUserDTO(dbUser, dbUserAdditionalInfo)

	return *user, nil

}


func (repository *userRepository) GetAllUsers(ctx context.Context) ([]persistence.User, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllUsers")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var users []persistence.User

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


func (repository *userRepository) CreateUser(ctx context.Context, user *persistence.User) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
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