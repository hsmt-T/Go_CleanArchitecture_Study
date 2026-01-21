package repository

import (
	"errors"
	"go_cleanArchitecture_study/domain"
	"sync"
)

// In-memory User Repository（DBの代わり）
// 小文字で直接使わせない/この型を必ず経由して生成させたい
type userMemoryRepository struct {
	mu    sync.Mutex
	users map[domain.UserID]domain.User
}

//コンストラクタ

func NewUserMemoryRepository() domain.UserRepository {
	return &userMemoryRepository{
		users: make(map[domain.UserID]domain.User),
	}
}

//Create
// domain.userを返してインターフェイスを返している
func (r *userMemoryRepository) Create(user domain.User) (domain.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.users[user.ID()]; exists {
		return domain.User{}, errors.New("このIDは使用されています")
	}

	r.users[user.ID()] = user
	return user, nil
}

//FindByID
func (r *userMemoryRepository) FindByID(id domain.UserID) (domain.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	user, exists := r.users[id]
	if !exists {
		return  domain.User{}, errors.New("このIDのユーザーは存在しません")
	}

	return user, nil
}

// domainで定義した関数をrepositoryで作る　（名前は違うが前開発した時と同じ）