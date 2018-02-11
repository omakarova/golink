package storage

import (
	"../mymodels"
)

func ValidateNewUser(user mymodels.NewUser) bool {
	return len([]rune(user.Username)) <= 32 && len([]rune(user.Password)) <= 128
}
