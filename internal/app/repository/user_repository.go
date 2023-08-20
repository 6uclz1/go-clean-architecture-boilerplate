// repository/user_repository.go

package repository

import (
	"go-clean-architecture-boilerplate/internal/app/entity"
)

type UserRepository struct {
    // データベース接続やクエリの実行に必要なものを含む
}

func NewUserRepository() *UserRepository {
    // 初期化処理...
    return &UserRepository{}
}

func (r *UserRepository) FindByID(id int) (*entity.User, error) {
    // データベースクエリの実行...
    // エンティティの作成と返却...
}
