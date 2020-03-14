package main

import "fmt"

type User struct {
	Username string
	Name     string
	Surname  string
}

func NUsers(n int) []User {
	var users []User
	for i := 0; i < n; i++ {
		users = append(users, User{
			Username: fmt.Sprintf("username %d", i),
			Name:     fmt.Sprintf("name %d", i),
			Surname:  fmt.Sprintf("surname %d", i),
		})
	}

	return users
}
