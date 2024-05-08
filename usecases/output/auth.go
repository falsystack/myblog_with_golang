package output

type AuthResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name" gorm:"unique"`
	Email string `json:"email" gorm:"unique"`
}

func NewAuthResponse(id uint, name, email string) *AuthResponse {
	return &AuthResponse{
		ID:    id,
		Name:  name,
		Email: email,
	}
}
