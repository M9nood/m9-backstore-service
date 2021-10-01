package store

type StoreCreateRequest struct {
	Id   int    `gorm:"column:id" json:"id,omitempty"`
	Name string `gorm:"column:name" json:"store_name"`
	Type int    `gorm:"column:type" json:"type"`
}

type StoreOwnerCreateRequest struct {
	Id        int `gorm:"column:id" json:"id,omitempty"`
	OwnerType int `gorm:"column:owner_type" json:"owner_type"`
	UserId    int `gorm:"column:user_id" json:"user_id"`
	StoreId   int `gorm:"column:store_id" json:"store_id"`
}
