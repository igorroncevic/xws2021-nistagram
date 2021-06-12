package domain

import (
	"errors"
	"regexp"
	"strings"
)

const (
	passwordRegex = "^(?=.*[\\d])(?=.*[A-Z])(?=.*[a-z])(?=.*[!@#$%^&*])[\\w!@#$%^&*]{8,}$"
)

func (user *User) CheckAllFields() (bool, error) {
	checkRegex, err := user.checkRegexValidation()
	if !checkRegex { return false, errors.New("field does not obey to regex rule") }
	if err != nil { return false, err }

	checkRegex, err = user.checkField(user.FirstName)
	if !checkRegex { return false, errors.New("field does not obey to regex rule") }
	if err != nil { return false, err }

	checkRegex, err = user.checkField(user.LastName)
	if !checkRegex { return false, errors.New("field does not obey to regex rule") }
	if err != nil { return false, err }

	checkRegex, err = user.checkField(user.Username)
	if !checkRegex { return false, errors.New("field does not obey to regex rule") }
	if err != nil { return false, err }

	checkRegex, err = user.checkField(user.Email)
	if !checkRegex { return false, errors.New("field does not obey to regex rule") }
	if err != nil { return false, err }

	checkRegex, err = user.checkField(user.Website)
	if !checkRegex { return false, errors.New("field does not obey to regex rule") }
	if err != nil { return false, err }

	checkRegex, err = user.checkField(user.Biography)
	if !checkRegex { return false, errors.New("field does not obey to regex rule") }
	if err != nil { return false, err }

	checkRegex, err = user.checkField(user.PhoneNumber)
	if !checkRegex { return false, errors.New("field does not obey to regex rule") }
	if err != nil { return false, err }

	return true, nil
}

func (user *User) checkField(field string) (bool, error){
	specialWordCheck, err := checkSpecialWordsValidation(field)
	if !specialWordCheck { return false, errors.New("field contains special word") }
	if err != nil { return false, err }

	return true, nil
}

func checkSpecialWordsValidation(s string) (bool, error) {
	res := strings.Split(s, " ")
	for _, item := range(res) {
		match, _ := regexp.MatchString("\\b(ALTER|CREATE|DELETE|DROP|EXEC(UTE){0,1}|INSERT( +INTO){0,1}|MERGE|SELECT|UPDATE|UNION( +ALL){0,1})\\b", item)
		if match {
			return false, errors.New("bad request!")
		}
		match, _ = regexp.MatchString("\\b(alter|create|delete|drop|exec(ute){0,1}|insert( +into){0,1}|merge|select|update|union( +all){0,1})\\b", item)
		if match {
			return false, errors.New("bad request!")
		}
	}

	return true, nil
}

func (user *User) checkRegexValidation() (bool, error){
	match, err := regexp.MatchString("^[a-zA-Z ,.'-]+$", user.FirstName)
	if !match { return false, err }

	checkSpecialWords, err := checkSpecialWordsValidation(user.FirstName)
	if !checkSpecialWords { return false, errors.New("first name contains invalid characters") }
	if err != nil { return false, err }

	match, err = regexp.MatchString("^[a-zA-Z ,.'-]+$", user.LastName)
	if !match { return false, err }

	match, err = regexp.MatchString("^[a-zA-Z ,.'-]+$", user.Username)
	if !match { return false, err }

	//match, err = regexp.MatchString("^[a-zA-Z ,.'-]+$", user.Username)
	match, err = regexp.MatchString("[A-Z0-9.%+-][A-Z0-9.-][A-Z]{0,500}$", user.Biography)
	if !match { return false, err }

	match, err = regexp.MatchString("[A-Z0-9.%+-][A-Z0-9.-][A-Z]{0,300}$", user.Website)
	if !match { return false, err }

	match, err = regexp.MatchString("^[a-zA-Z]+$", user.Sex)
	if !match { return false, err }

	match, err = regexp.MatchString("^[A-Z0-9.%+-]+@[A-Z0-9.-]+\\.[A-Z]{2,64}$", user.Email)
	if !match { return false, err }

	match, err = regexp.MatchString("^[+]?[0-9]{8,12}$", user.PhoneNumber)
	if !match { return false, err }

	return true, nil
}

func (password *Password) CheckAllFields() (bool, error) {
	checkRegex, err := password.checkRegexValidation()
	if !checkRegex { return false, errors.New("field does not obey to regex rule") }
	if err != nil { return false, err }

	checkRegex, err = password.checkField(password.NewPassword)
	if !checkRegex { return false, errors.New("field does not obey to regex rule") }
	if err != nil { return false, err }

	checkRegex, err = password.checkField(password.OldPassword)
	if !checkRegex { return false, errors.New("field does not obey to regex rule") }
	if err != nil { return false, err }

	checkRegex, err = password.checkField(password.RepeatedPassword)
	if !checkRegex { return false, errors.New("field does not obey to regex rule") }
	if err != nil { return false, err }

	return true, nil
}

func (password *Password) checkField(field string) (bool, error){
	specialWordCheck, err := checkSpecialWordsValidation(field)
	if !specialWordCheck { return false, errors.New("field contains special word") }
	if err != nil { return false, err }

	return true, nil
}

func (password *Password) checkRegexValidation() (bool, error){
	/*match, err := regexp.MatchString(passwordRegex, password.OldPassword)
	if !match { return false, err }

	match, err = regexp.MatchString(passwordRegex, password.NewPassword)
	if !match { return false, err }

	match, err = regexp.MatchString(passwordRegex, password.RepeatedPassword)
	if !match { return false, err }*/

	return true, nil
}

