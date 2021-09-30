package store

import "github.com/guregu/null"

type StoreSchema struct {
	Id         int       `gorm:"column:id" json:"id"`
	Name       string    `gorm:"column:store_name" json:"store_name"`
	Type       int       `gorm:"column:type" json:"type"`
	StoreUuid  string    `gorm:"column:store_uuid" json:"store_uuid"`
	CreatedAt  null.Time `gorm:"column:created_at" json:"created_at"`
	UpatedAt   null.Time `gorm:"column:updated_at" json:"updated_at"`
	DeleteFlag int       `gorm:"column:delete_flag" json:"delete_flag"`
}
