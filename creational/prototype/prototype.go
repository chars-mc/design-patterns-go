package main

import "fmt"

// user it's the concrete prototype
type user struct {
	name    string
	age     int
	friends []string
	color   *string
	phones  map[string]string
}

// string returns the name and age
func (u *user) string() string {
	return fmt.Sprintf(
		"Name: %s\tAge:%d \tFriends: %v \tColor: %v \tPhones: %v",
		u.name,
		u.age,
		u.friends,
		*u.color,
		u.phones,
	)
}

// clone returns a copy of a user
func (u *user) clone() *user {
	friends := make([]string, len(u.friends))
	copy(friends, u.friends)
	// friends := append(friends, u.friends...)

	color := *u.color

	phones := make(map[string]string)
	for k, v := range u.phones {
		phones[k] = v
	}

	return &user{
		name:    u.name,
		age:     u.age,
		friends: friends,
		color:   &color,
		phones:  phones,
	}
}
