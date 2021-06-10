package security

import (
	"bufio"
	"os"
	"path/filepath"
)

func CheckBlacklistedPassword(password string) bool {
	passwords, err := loadPasswords()
	if err != nil {
		return true
	}

	for _, passwordCheck := range passwords {
		if password == passwordCheck {
			return true
		}
	}

	return false
}

func loadPasswords() ([]string, error){
	pwd, _ := os.Getwd()

	file, err := os.Open(filepath.Join(pwd, "password_blacklist.txt"))
	if err != nil {
		return []string{}, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	passwords := []string{}
	for scanner.Scan() {
		passwords = append(passwords, scanner.Text())
	}

	return passwords, nil
}
