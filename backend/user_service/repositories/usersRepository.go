package repositories

import (
	"context"
	"errors"
	"fmt"
	"github.com/david-drvar/xws2021-nistagram/common"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/user_service/model/domain"
	"github.com/david-drvar/xws2021-nistagram/user_service/model/persistence"
	"github.com/david-drvar/xws2021-nistagram/user_service/util/encryption"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUsers(context.Context) ([]persistence.User, error)
	CreateUser(context.Context, *persistence.User) error
	CreateUserWithAdditionalInfo(context.Context, *persistence.User, *persistence.UserAdditionalInfo) error
	CheckPassword(data common.Credentials) error
	UpdateUserProfile(ctx context.Context, dto domain.User) (bool, error)
	UpdateUserPassword(ctx context.Context, password domain.Password) (bool, error)
	SearchUsersByUsernameAndName(ctx context.Context, user *domain.User) ([]domain.User, error)
}

type userRepository struct {
	DB                *gorm.DB
	privacyRepository PrivacyRepository
}

func NewUserRepo(db *gorm.DB) (*userRepository, error) {
	if db == nil {
		panic("UserRepository not created, gorm.DB is nil")
	}
	privacyRepository, _ := NewPrivacyRepo(db)

	return &userRepository{DB: db, privacyRepository: privacyRepository}, nil
}

func (repository *userRepository) UpdateUserPassword(ctx context.Context, password domain.Password) (bool, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "UpdateUserPassword")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var user *persistence.User

	db := repository.DB.Select("password").Where("id = ?", password.Id).Find(&user)
	if db.Error != nil {
		return false, db.Error
	} else if db.RowsAffected == 0 {
		return false, errors.New("rows affected is equal to zero")
	}

	err := encryption.CompareHashAndPassword([]byte(user.Password), []byte(password.OldPassword))
	if err != nil {
		return false, err
	}

	db = repository.DB.Model(&user).Where("id = ?", password.Id).Updates(persistence.User{Password: encryption.HashAndSalt([]byte(password.NewPassword))})
	if db.Error != nil {
		return false, db.Error
	} else if db.RowsAffected == 0 {
		return false, errors.New("rows affected is equal to zero")
	}

	return true, nil
}

func (repository *userRepository) UpdateUserProfile(ctx context.Context, userDTO domain.User) (bool, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "UpdateUserProfile")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var user persistence.User
	var userAdditionalInfo persistence.UserAdditionalInfo

	db := repository.DB.Model(&user).Where("id = ?", userDTO.Id).Updates(persistence.User{FirstName: userDTO.FirstName, LastName: userDTO.LastName, Email: userDTO.Email, Username: userDTO.Username, BirthDate: userDTO.BirthDate,
		PhoneNumber: userDTO.PhoneNumber, Sex: userDTO.Sex})

	fmt.Println(db.RowsAffected)

	if db.Error != nil {
		return false, db.Error
	} else if db.RowsAffected == 0 {
		return false, errors.New("rows affected is equal to zero")
	}
	db = repository.DB.Model(&userAdditionalInfo).Where("id = ?", userDTO.Id).Updates(persistence.UserAdditionalInfo{Website: userDTO.Website, Category: userDTO.Category, Biography: userDTO.Biography})

	if db.Error != nil {
		return false, db.Error
	} else if db.RowsAffected == 0 {
		return false, errors.New("rows affected is equal to zero")
	}

	return true, nil
}

func (repository *userRepository) GetUserByUsername(username string) (domain.User, error) {
	var dbUser persistence.User
	var dbUserAdditionalInfo persistence.UserAdditionalInfo

	db := repository.DB.Where("username = ?", username).Find(&dbUser)
	if db.Error != nil {
		return domain.User{}, db.Error
	}

	db = repository.DB.Where("id = ?", dbUser.Id).Find(&dbUserAdditionalInfo)
	if db.Error != nil {
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

	user.Id = uuid.New().String()
	result := repository.DB.Create(&user)

	return result.Error
}

func (repository *userRepository) CheckEmailExists(ctx context.Context, email string) bool {
	span := tracer.StartSpanFromContextMetadata(ctx, "CheckEmailExists")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	count := int64(0)
	repository.DB.Model(persistence.User{}).Where("email = ?", email).Count(&count)
	return count > 0

}

func (repository *userRepository) CreateUserWithAdditionalInfo(ctx context.Context, user *persistence.User, userAdditionalInfo *persistence.UserAdditionalInfo) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateUserWithAdditionalInfo")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var userPersistence domain.User
	userPersistence, _ = repository.GetUserByUsername(user.Username)
	if userPersistence.Username == user.Username {
		return errors.New("username already exists")
	}

	if repository.CheckEmailExists(ctx, user.Email) {
		return errors.New("email already exists")
	}

	user.Id = uuid.New().String()
	resultUser := repository.DB.Create(&user)
	if resultUser.Error != nil {
		return resultUser.Error
	}



	userAdditionalInfo.Id = user.Id
	userAdditionalInfo.Category=user
	resultUserAdditionalInfo := repository.DB.Create(&userAdditionalInfo)

	var privacy = persistence.Privacy{}
	privacy.UserId = user.Id
	privacy.IsProfilePublic = true
	privacy.IsDMPublic = true
	privacy.IsTagEnabled = true
	_, err := repository.privacyRepository.CreatePrivacy(ctx, &privacy)
	if err != nil {
		return err
	}

	return resultUserAdditionalInfo.Error
}

func (repository *userRepository) SearchUsersByUsernameAndName(ctx context.Context, user *domain.User) ([]domain.User, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "SearchUsersByUsernameAndName")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var users []persistence.User

	if user.Username != "" && user.FirstName != "" && user.LastName != "" {
		repository.DB.Where("username = ? AND first_name = ? AND last_name = ?", user.Username, user.FirstName, user.LastName).Find(&users)
	} else if user.Username != "" && user.FirstName != "" && user.LastName == "" {
		repository.DB.Where("username = ? AND first_name = ?", user.Username, user.FirstName).Find(&users)
	} else if user.Username != "" && user.FirstName == "" && user.LastName != "" {
		repository.DB.Where("username = ? AND last_name = ?", user.Username, user.LastName).Find(&users)
	} else if user.Username == "" && user.FirstName != "" && user.LastName != "" {
		repository.DB.Where("first_name = ? AND last_name = ?", user.FirstName, user.LastName).Find(&users)
	} else if user.Username != "" && user.FirstName == "" && user.LastName == "" {
		repository.DB.Where("username = ?", user.Username).Find(&users)
	} else if user.Username == "" && user.FirstName != "" && user.LastName == "" {
		repository.DB.Where("first_name = ?", user.FirstName).Find(&users)
	} else if user.Username == "" && user.FirstName == "" && user.LastName != "" {
		repository.DB.Where("last_name = ?", user.LastName).Find(&users)
	}

	var usersDomain []domain.User

	for _, v := range users { //i - index, v - user
		user := &domain.User{}
		user.GenerateUserDTO(v, persistence.UserAdditionalInfo{})
		usersDomain = append(usersDomain, *user)
	}

	return usersDomain, nil
}
