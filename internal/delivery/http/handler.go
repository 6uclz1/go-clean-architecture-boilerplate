// http/handler.go

package http

import (
	"net/http"

	"go-clean-architecture-boilerplate/internal/app/usecase"
)

type Handler struct {
    userUseCase *usecase.UserUseCase
}

func NewHandler(userUseCase *usecase.UserUseCase) *Handler {
    return &Handler{
        userUseCase: userUseCase,
    }
}

func (h *Handler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
    // リクエストのパラメータからユーザーIDを取得...
    userID := 1

    user, err := h.userUseCase.GetUserByID(userID)
    if err != nil {
        // エラーハンドリング...
        return
    }

    // レスポンスの作成と送信...
}
