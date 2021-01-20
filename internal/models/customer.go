package models

import "time"

// Customer is the user of this app
type Customer struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	MobileNo  string    `json:"mobile_no"`
	Password  string    `json:"password"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// IsEmpty checks if object is empty
func (c Customer) IsEmpty() bool {
	return c.ID == "" && c.Email == "" && c.MobileNo == "" && c.Password == "" && c.FirstName == "" && c.LastName == "" && c.CreatedAt == time.Time{} && c.UpdatedAt == time.Time{}
}
