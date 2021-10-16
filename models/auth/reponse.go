package auth

type LoginResponse struct {
	Id       int    `json:"id"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	StoreId  *int   `json:"store_id"`
	Token    Token  `json:"token,omitempty"`
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
