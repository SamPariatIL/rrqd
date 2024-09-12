package repository

import "go.mongodb.org/mongo-driver/mongo"

type AuthRepository interface {
	Login()
	RefreshToken()
	UpdatePassword()
	ResetPassword()
	VerifyOTP()
	SendOTP()
	WarehouseLogin()
	WarehouseVerifyOTP()
	DeliveryLogin()
	DeliveryVerifyOTP()
	PrincipalLogin()
	PrincipalResetPassword()
	UpdatePrincipalPassword()
}

type authRepository struct {
	db *mongo.Database
}

func NewAuthRepository(db *mongo.Database) AuthRepository {
	return &authRepository{db: db}
}

func (ar *authRepository) Login() {}

func (ar *authRepository) RefreshToken() {}

func (ar *authRepository) UpdatePassword() {}

func (ar *authRepository) ResetPassword() {}

func (ar *authRepository) VerifyOTP() {}

func (ar *authRepository) SendOTP() {}

func (ar *authRepository) WarehouseLogin() {}

func (ar *authRepository) WarehouseVerifyOTP() {}

func (ar *authRepository) DeliveryLogin() {}

func (ar *authRepository) DeliveryVerifyOTP() {}

func (ar *authRepository) PrincipalLogin() {}

func (ar *authRepository) PrincipalResetPassword() {}

func (ar *authRepository) UpdatePrincipalPassword() {}
