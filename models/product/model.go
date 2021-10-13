package product

import (
	"fmt"

	"github.com/guregu/null"
)

type ProductSchema struct {
	Id          int        `gorm:"column:id" json:"id"`
	ProductName string     `gorm:"column:product_name" json:"product_name"`
	Description string     `gorm:"column:description" json:"description"`
	DispCode    string     `gorm:"column:disp_code" json:"disp_code"`
	InStock     int        `gorm:"column:in_stock" json:"in_stock"`
	ProductUuid *string    `gorm:"-;column:product_uuid" json:"product_uuid,omitempty"`
	StoreId     int        `gorm:"column:store_id" json:"store_id"`
	ImageKey    string     `gorm:"column:image_key" json:"image_key"`
	Cost        null.Float `gorm:"column:cost" json:"cost"`
	Price       null.Float `gorm:"column:price" json:"price"`
}

func ToLineMessage(p ProductSchema) string {
	return fmt.Sprintf("%s   x%d\n", p.ProductName, p.InStock)
}
