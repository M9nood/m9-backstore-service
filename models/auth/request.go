package auth

type RegisterRequest struct {
	UserName  string `json:"username" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Password  string `json:"password" validate:"required"`
	StoreName string `json:"store_name" validate:"required"`
	StoreType int    `json:"store_type" validate:"required"`
}
