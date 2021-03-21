package models_test

import (
	"fmt"
	"testing"

	"github.com/hbjydev/mangadex-next/models"
)

func TestUserFromJSON(t *testing.T) {
	userString := "{\"id\":\"udhy8dauwdqwu8\",\"username\":\"johndoe\",\"email\":\"johndoe@gmail.com\"}"

	user, err := models.UserFromJSON(userString)
	if err != nil {
		t.Error(err)
	}

	if user.Username != "johndoe" {
		t.Errorf("Expected username to be 'johndoe' but got %v", user.Username)
	}
	fmt.Println(user)
}
