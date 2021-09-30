package store

type StoreCreateRequest struct {
	Id   int    `gorm:"column:id" json:"id,omitempty"`
	Name string `gorm:"column:name" json:"store_name"`
	Type int    `gorm:"column:type" json:"type"`
}
