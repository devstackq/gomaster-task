package models

import "time"

//entity
type User struct {
	ID          string    `json:"uuid"`
	FirstName   string    `json:"firstname"`
	LastName    string    `json:"lastname"`
	Email       string    `json:"email"`
	Age         int       `json:"age"`
	CreatedTime time.Time `json:"createdtime"`
}
