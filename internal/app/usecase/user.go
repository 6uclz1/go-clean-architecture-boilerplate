// usecase/user.go

package usecase

import "go-clean-architecture-boilerplate/internal/app/entity"

type UserRepository interface {
    FindByID(id int) (*entity.User, error)
    // 他のメソッド...
}

type UserUseCase struct {
    userRepository UserRepository
}

func NewUserUseCase(userRepository UserRepository) *UserUseCase {
    return &UserUseCase{
        userRepository: userRepository,
    }
}

func (uc *UserUseCase) GetUserByID(id int) (*entity.User, error) {
    user, err := uc.userRepository.FindByID(id)
    if err != nil {
        return nil, err
    }
    // ビジネスロジックの処理...
    return user, nil
}
