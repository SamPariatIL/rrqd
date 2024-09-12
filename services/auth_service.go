package services

import "github.com/SamPariatIL/rrqd/repository"

type AuthService interface {
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

type authService struct {
	authRepository repository.AuthRepository
}

func NewAuthService(ar repository.AuthRepository) AuthService {
	return &authService{authRepository: ar}
}

func (as *authService) Login() {}

func (as *authService) RefreshToken() {}

func (as *authService) UpdatePassword() {}

func (as *authService) ResetPassword() {}

func (as *authService) VerifyOTP() {}

func (as *authService) SendOTP() {}

func (as *authService) WarehouseLogin() {}

func (as *authService) WarehouseVerifyOTP() {}

func (as *authService) DeliveryLogin() {}

func (as *authService) DeliveryVerifyOTP() {}

func (as *authService) PrincipalLogin() {}

func (as *authService) PrincipalResetPassword() {}

func (as *authService) UpdatePrincipalPassword() {}
