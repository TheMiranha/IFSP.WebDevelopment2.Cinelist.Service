package dtos

type SignUpDTO struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding="required"`
}

type SignUpResponseDTO struct {
	Success     bool   `json:"success"`
	AccessToken string `json:"accessToken"`
}
