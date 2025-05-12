package dto

type UserResponse struct {
	Id               string  `json:"id"`
	FullName         string  `json:"fullName"`
	Email            string  `json:"email"`
	PhotoProfile     string  `json:"photoProfile"`
	IsVerified       bool    `json:"isVerified"`
	BodyWeight       float64 `json:"bodyWeight"`
	DateOfBirth      string  `json:"dateOfBirth"`
	Address          string  `json:"address"`
	CaregiverContact string  `json:"caregiverContact"`
}

type UpdateUserRequest struct {
	FullName         string  `json:"fullName" validate:"required"`
	Email            string  `json:"email" validate:"required,email"`
	PhotoProfile     string  `json:"photoProfile"`
	BodyWeight       float64 `json:"bodyWeight" validate:"required"`
	DateOfBirth      string  `json:"dateOfBirth" validate:"required"`
	Address          string  `json:"address"`
	CaregiverContact string  `json:"caregiverContact"`
}
