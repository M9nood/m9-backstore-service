package store

import (
	"database/sql"

	"github.com/guregu/null"
)

type StoreSchema struct {
	Id         int            `gorm:"column:id" json:"id,omitempty"`
	Name       string         `gorm:"column:name" json:"store_name"`
	Type       int            `gorm:"column:type" json:"type"`
	StoreUuid  sql.NullString `gorm:"column:store_uuid" json:"store_uuid,omitempty"`
	CreatedAt  null.Time      `gorm:"column:created_at" json:"created_at"`
	UpatedAt   null.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeleteFlag int            `gorm:"column:delete_flag" json:"delete_flag"`
}
