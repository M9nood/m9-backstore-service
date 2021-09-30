package user

type UserCreateRequest struct {
	Id       int    `gorm:"column:id" json:"id,omitempty"`
	UserName string `gorm:"column:username" json:"username"`
	Email    string `gorm:"column:email" json:"email"`
	Password string `gorm:"column:password" json:"password,omitempty"`
}
