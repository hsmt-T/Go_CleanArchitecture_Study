package domain

//IDの一意性という ドメインルールを守るため domain/uuid.go

import (
	gouuid "github.com/satori/go.uuid"
)

type UserID string

func NewUserID() UserID {
	return UserID(gouuid.NewV4().String())
}

func IsValidUserID(id string) bool {
	_, err := gouuid.FromString(id)
	return err == nil
}