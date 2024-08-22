package handlers

import (
	"users/storage/postgres"

	"go.uber.org/zap"
)

type Handler struct {
	UserRepo *postgres.UserRepo
	Log         *zap.Logger
}

func NewHandler(users *postgres.UserRepo, log *zap.Logger) *Handler {
	return &Handler{
		UserRepo: users,
		Log:         log,
	}
}
