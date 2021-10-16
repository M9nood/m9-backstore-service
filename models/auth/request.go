package auth

type RegisterRequest struct {
	UserName  string `json:"username" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Password  string `json:"password" validate:"required"`
	StoreName string `json:"store_name" validate:"required"`
	StoreType int    `json:"store_type" validate:"required"`
}

type LoginRequest struct {
	UserName     string `json:"username" validate:"required"`
	PasswordHash string `json:"password" validate:"required"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}
