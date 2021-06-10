package repositories

import (
	"context"
	"errors"
	"fmt"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/user_service/model/domain"
	"github.com/david-drvar/xws2021-nistagram/user_service/model/persistence"
	"github.com/david-drvar/xws2021-nistagram/user_service/util"
	"github.com/david-drvar/xws2021-nistagram/user_service/util/encryption"
	"github.com/david-drvar/xws2021-nistagram/user_service/util/images"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type UserRepository interface {
	GetAllUsers(context.Context) ([]persistence.User, error)
	CreateUser(context.Context, *persistence.User) error
	CreateUserWithAdditionalInfo(context.Context, *persistence.User, *persistence.UserAdditionalInfo) (*domain.User, error)
	UpdateUserProfile(ctx context.Context, dto domain.User) (bool, error)
	UpdateUserPassword(ctx context.Context, password domain.Password) (bool, error)
	SearchUsersByUsernameAndName(ctx context.Context, user *domain.User) ([]domain.User, error)
	LoginUser(context.Context, domain.LoginRequest) (persistence.User, error)
	SaveUserProfilePhoto(ctx context.Context, user *persistence.User) (bool, error)
	GetUserAdditionalInfoById(ctx context.Context, id string) (persistence.UserAdditionalInfo, error)
	GetUserByEmail(email string) (domain.User, error)
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
		PhoneNumber: userDTO.PhoneNumber, Sex: userDTO.Sex,ResetCode: userDTO.ResetCode, ApprovedAccount: userDTO.ApprovedAccount, TokenEnd: userDTO.TokenEnd})

	fmt.Println(db.RowsAffected)

	if db.Error != nil {
		return false, db.Error
	} else if db.RowsAffected == 0 {
		return false, errors.New("rows affected is equal to zero")
	}

	userAdditionalInfoUpdate := persistence.UserAdditionalInfo{Website: userDTO.Website, Category: userDTO.Category, Biography: userDTO.Biography}
	if userAdditionalInfoUpdate.Website == "" {
		userAdditionalInfoUpdate.Website = " "
	}
	if userAdditionalInfoUpdate.Biography == "" {
		userAdditionalInfoUpdate.Biography = " "
	}
	db = repository.DB.Model(&userAdditionalInfo).Where("id = ?", userDTO.Id).Updates(userAdditionalInfoUpdate)

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

	filename, err := images.LoadImageToBase64(user.ProfilePhoto)
	if err != nil {
		return domain.User{}, err
	}
	user.ProfilePhoto = filename

	return *user, nil
}

func (repository *userRepository) GetUserByEmail(email string) (domain.User, error) {
	var dbUser persistence.User
	var dbUserAdditionalInfo persistence.UserAdditionalInfo

	db := repository.DB.Where("email = ?", email).Find(&dbUser)
	if db.Error != nil {
		return domain.User{}, db.Error
	}

	db = repository.DB.Where("id = ?", dbUser.Id).Find(&dbUserAdditionalInfo)
	if db.Error != nil {
		return domain.User{}, db.Error
	}
	user := &domain.User{}

	user.GenerateUserDTO(dbUser, dbUserAdditionalInfo)

	filename, err := images.LoadImageToBase64(user.ProfilePhoto)
	if err != nil {
		return domain.User{}, err
	}
	user.ProfilePhoto = filename

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

func (repository *userRepository) LoginUser(ctx context.Context, request domain.LoginRequest) (persistence.User, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "LoginUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var dbUser persistence.User
	result := repository.DB.Where("email = ?", request.Email).First(&dbUser)
	if result.Error != nil || result.RowsAffected != 1 {
		return persistence.User{}, result.Error
	}

	err := encryption.CompareHashAndPassword([]byte(dbUser.Password), []byte(request.Password))
	if err != nil {
		return persistence.User{}, errors.New("passwords do not match")
	}

	return dbUser, nil
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

func (repository *userRepository) CreateUserWithAdditionalInfo(ctx context.Context, user *persistence.User, userAdditionalInfo *persistence.UserAdditionalInfo) (*domain.User, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateUserWithAdditionalInfo")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var userDomain domain.User
	userDomain, _ = repository.GetUserByUsername(user.Username)
	if userDomain.Username == user.Username {
		return nil, errors.New("username already exists")
	}

	if repository.CheckEmailExists(ctx, user.Email) {
		return nil, errors.New("email already exists")
	}

	user.Id = uuid.New().String()
	resultUser := repository.DB.Create(&user)
	if resultUser.Error != nil {
		return nil, resultUser.Error
	}

	_, err := repository.SaveUserProfilePhoto(ctx, user)
	if err != nil {
		return nil, err
	}

	userAdditionalInfo.Id = user.Id
	if userAdditionalInfo.Website == "" {
		userAdditionalInfo.Website = " "
	}
	if userAdditionalInfo.Biography == "" {
		userAdditionalInfo.Biography = " "
	}
	err2 := repository.DB.Create(&userAdditionalInfo)
	print(err2)

	var privacy = persistence.Privacy{}
	privacy.UserId = user.Id
	privacy.IsProfilePublic = true
	privacy.IsDMPublic = true
	privacy.IsTagEnabled = true
	_, err = repository.privacyRepository.CreatePrivacy(ctx, &privacy)
	if err != nil {
		return nil, err
	}

	userReturn := &domain.User{}
	userReturn.GenerateUserDTO(*user, *userAdditionalInfo)

	return userReturn, nil
}

func (repository *userRepository) GetUserAdditionalInfoById(ctx context.Context, id string) (persistence.UserAdditionalInfo, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "SearchUsersByUsernameAndName")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var dbUserAdditionalInfo persistence.UserAdditionalInfo
	db := repository.DB.Where("id = ?", id).Find(&dbUserAdditionalInfo)
	if db.Error != nil {
		return persistence.UserAdditionalInfo{}, db.Error
	}
	return dbUserAdditionalInfo, nil
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
		dbUserAdditionalInfo, err := repository.GetUserAdditionalInfoById(ctx, v.Id)
		if err != nil {
			return nil, err
		}
		user.GenerateUserDTO(v, dbUserAdditionalInfo)
		filename, err := images.LoadImageToBase64(v.ProfilePhoto)
		if err != nil {
			return nil, err
		}
		user.ProfilePhoto = filename
		usersDomain = append(usersDomain, *user)
	}

	return usersDomain, nil
}

func (repository *userRepository) SaveUserProfilePhoto(ctx context.Context, user *persistence.User) (bool, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "SaveUserProfilePhoto")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	mimeType, err := images.GetImageType(user.ProfilePhoto)
	if err != nil {
		return false, err
	}

	t := time.Now()
	formatted := fmt.Sprintf("%s%d%02d%02d%02d%02d%02d%02d", user.Id, t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond())
	name := formatted + "." + mimeType

	err = images.SaveImage(name, user.ProfilePhoto)
	if err != nil {
		return false, err
	}

	user.ProfilePhoto = util.GetContentLocation(name)
	db := repository.DB.Model(&user).Where("id = ?", user.Id).Updates(user)
	if db.Error != nil {
		return false, db.Error
	}

	return true, nil
}
