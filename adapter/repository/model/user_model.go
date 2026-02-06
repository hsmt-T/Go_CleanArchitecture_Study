package model

import("time")

type UserModel struct {
	ID string  `gorm:"primaryKey"`
	Name		string
	Email		string
	CreatedAt	time.Time
}

func (UserModel) TableName() string {
	return "users"
}