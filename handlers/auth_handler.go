package handlers

import (
	"github.com/SamPariatIL/rrqd/services"
	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	Login(ctx *gin.Context)
	RefreshToken(ctx *gin.Context)
	UpdatePassword(ctx *gin.Context)
	ResetPassword(ctx *gin.Context)
	VerifyOTP(ctx *gin.Context)
	SendOTP(ctx *gin.Context)
	WarehouseLogin(ctx *gin.Context)
	WarehouseVerifyOTP(ctx *gin.Context)
	DeliveryLogin(ctx *gin.Context)
	DeliveryVerifyOTP(ctx *gin.Context)
	PrincipalLogin(ctx *gin.Context)
	PrincipalResetPassword(ctx *gin.Context)
	UpdatePrincipalPassword(ctx *gin.Context)
}

type authHandler struct {
	authService services.AuthService
}

func NewAuthHandler(as services.AuthService) AuthHandler {
	return &authHandler{authService: as}
}

func (ah *authHandler) Login(ctx *gin.Context) {}

func (ah *authHandler) RefreshToken(ctx *gin.Context) {}

func (ah *authHandler) UpdatePassword(ctx *gin.Context) {}

func (ah *authHandler) ResetPassword(ctx *gin.Context) {}

func (ah *authHandler) VerifyOTP(ctx *gin.Context) {}

func (ah *authHandler) SendOTP(ctx *gin.Context) {}

func (ah *authHandler) WarehouseLogin(ctx *gin.Context) {}

func (ah *authHandler) WarehouseVerifyOTP(ctx *gin.Context) {}

func (ah *authHandler) DeliveryLogin(ctx *gin.Context) {}

func (ah *authHandler) DeliveryVerifyOTP(ctx *gin.Context) {}

func (ah *authHandler) PrincipalLogin(ctx *gin.Context) {}

func (ah *authHandler) PrincipalResetPassword(ctx *gin.Context) {}

func (ah *authHandler) UpdatePrincipalPassword(ctx *gin.Context) {}
