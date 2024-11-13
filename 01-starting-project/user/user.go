package user

import (
  "fmt"
  "time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

type Admin struct {
  email string
  password string
  user User
}

func NewAdmin(email, password string) Admin {
  return Admin {
    email: email,
    password: password,
    user: User {
      firstName: "ADMIN",
      lastName: "ADMIN",
      birthDate: "-",
      createdAt: time.Now(),
    },
  }
}

func New(firstName, lastName, birthDate string) *User {
  return &User {
    firstName: firstName,
    lastName: lastName,
    birthDate: birthDate,
    createdAt: time.Now(),
  }
}

func (u User) Print() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func (u *User) Murder() {
  fmt.Printf("%v %v has been successfully murdered\n", u.firstName, u.lastName)
  u.firstName = ""
  u.lastName = ""
}
