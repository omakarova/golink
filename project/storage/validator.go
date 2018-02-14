package storage

import (
	"../mymodels"
	"net/url"
	"fmt"
)

func NewUserIsValid(user mymodels.NewUser) bool {
	return len([]rune(user.Username)) <= 32 && len([]rune(user.Password)) <= 128
}

func NewLinkIsValid(link mymodels.NewLink) bool {
	var validURL bool

	_, err := url.Parse(link.URL)

	if err != nil {
		fmt.Println(err)
		validURL = false
	} else {
		validURL = true
	}

	fmt.Printf("%s is a valid URL : %v \n", link.URL, validURL)
	return validURL
}

