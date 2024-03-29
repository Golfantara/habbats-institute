package dtos

type ResUser struct {
	ID          int    `json:"id"`
	RoleID      int    `json:"role_id"`
	Fullname    string `json:"fullname"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	Gender      string `json:"gender"`
	Email       string `json:"email"`
}

type LoginResponse struct {
	Fullname     string `json:"fullname"`
	Email        string `json:"email"`
	RoleID       int    `json:"role_id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type ResJWT struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}