package dao

import (
	"time"
)

//Item - Table representative struc for a item in a list of a user
type Item struct {
	ID          int       `json:"id" db:"id"`
	UserName    string    `json:"userName" db:"userName"`
	Description string    `json:"description" db:"description"`
	Value       float64   `json:"value" db:"value"`
	CreateDate  time.Time `json:"createDate" db:"createDate"`
}
