package auth

type LoginResponse struct {
	Id       int    `json:"id"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Token    Token  `json:"token"`
}

type Token struct {
	AccessToken string `json:"access_token"`
}
