package dtos

type SignInDTO struct {
	Email    string `json:"email" binding="required"`
	Password string `json:"password" binding="required"`
}

type SignInResponseDTO struct {
	Success     bool   `json:"success"`
	AccessToken string `json:"accessToken"`
}
