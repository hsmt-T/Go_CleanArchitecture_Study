package domain

import (
	"time"
)

type UserID string

//　ここではUserの概念だけを定義
//	小文字なのは外から直接変更できなくするため
//　代わりに下でゲッターを作って読み込み出来るようにしている
type User struct {
	id 			UserID
	name		string
	email		string
	createdAt	time.Time
}

// NewUser ユーザーを新しく作るためのコンストラクタ
func NewUser(id UserID, name, email string, createdAt time.Time) User {
	return User{
		id:			id,
		name:		name,
		email:		email,
		createdAt: 	createdAt,
	}
}

//外から読み取るようの　ゲッター
func (u User) ID() UserID {
	return u.id
}
func (u User) Name() string {
	return u.name
}

func (u User) Email() string {
	return u.email
}

func (u User) CreatedAt() time.Time {
	return u.createdAt
}


// 	Repository インターフェース
//	何をするかだけここで定義して再利用できるようにしている
type UserRepository interface {
	Create(user User) (User, error)
	FindByID(id UserID) (User, error)
}