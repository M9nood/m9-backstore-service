package user

import (
	"database/sql"

	"github.com/guregu/null"
)

type UserSchema struct {
	Id         int            `gorm:"column:id" json:"id"`
	UserName   string         `gorm:"column:username" json:"username"`
	Email      string         `gorm:"column:email" json:"email"`
	Password   string         `gorm:"column:password" json:"password,omitempty"`
	UserUuid   sql.NullString `gorm:"column:user_uuid"  json:"user_uuid"`
	CreatedAt  null.Time      `gorm:"column:created_at" json:"created_at"`
	UpatedAt   null.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeleteFlag int            `gorm:"column:delete_flag" json:"delete_flag"`
}
