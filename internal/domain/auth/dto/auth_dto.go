package dto

type SignUpRequest struct {
	FullName        string `json:"fullName" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=8,password"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,min=8,password"`
}

type SignInRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,password"`
}

type VerifyEmailRequest struct {
	Code string `json:"code" validate:"required"`
}

type SignUpResponse struct {
	Id           string `json:"id"`
	FullName     string `json:"fullName"`
	Email        string `json:"email" `
	PhotoProfile string `json:"photoProfile"`
}

type SignInResponse struct {
	TokenType   string `json:"tokenType"`
	AccessToken string `json:"accessToken"`
	ExpiresAt   int64  `json:"expiresAt"`
}

type ChangePasswordRequest struct {
	NewPassword        string `json:"newPassword" validate:"required,min=8,password"`
	ConfirmNewPassword string `json:"confirmNewPassword" validate:"required,min=8,password"`
}

type VerifyForgotPasswordRequest struct {
	Code string `json:"code" validate:"required"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email" validate:"required"`
}
