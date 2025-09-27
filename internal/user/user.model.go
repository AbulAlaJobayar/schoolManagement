package user

import "time"

type User struct {
	ID        int `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
}