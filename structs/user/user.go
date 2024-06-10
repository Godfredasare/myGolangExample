package user

import (
	"errors"
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthday  string
	createdAt time.Time
}

func (u user) PrintOutput() {
	fmt.Println(u.firstName, u.lastName, u.birthday)
}

func (u *user) ClearUser() {
	u.firstName = " "
	u.lastName = " "
}

func New(firstName, lastName, birthday string) (*user, error) {
	if firstName == "" || lastName == "" || birthday == "" {
		return &user{}, errors.New("firstname, lastname, birthdate are required")
	}

	user := &user{
		firstName: firstName,
		lastName:  lastName,
		birthday:  birthday,
		createdAt: time.Now(),
	}

	return user, nil
}
