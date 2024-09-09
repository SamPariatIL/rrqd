package handlers

import (
	"github.com/SamPariatIL/rrqd/services"
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	GetUsers(ctx *gin.Context)
	GetUserById(ctx *gin.Context)
	CreateUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	UpdateUserStatus(ctx *gin.Context)
}

type userHandler struct {
	userService services.UserService
}

func NewUserHandler(us services.UserService) UserHandler {
	return &userHandler{userService: us}
}

func (uh *userHandler) GetUsers(ctx *gin.Context) {}

func (uh *userHandler) GetUserById(ctx *gin.Context) {}

func (uh *userHandler) CreateUser(ctx *gin.Context) {}

func (uh *userHandler) UpdateUser(ctx *gin.Context) {}

func (uh *userHandler) UpdateUserStatus(ctx *gin.Context) {}
