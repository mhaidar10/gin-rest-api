// models/users_m.go
package models

// User is a simple struct representing a user.
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Users is a slice of User representing a list of users.
var users = []User{
	{ID: 1, Name: "John"},
	{ID: 2, Name: "Jane"},
	{ID: 3, Name: "Doe"},
}
