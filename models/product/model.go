package product

import "fmt"

type ProductSchema struct {
	Id          int    `gorm:"column:id" json:"id"`
	ProductName string `gorm:"column:product_name" json:"product_name"`
	InStock     int    `gorm:"column:in_stock" json:"in_stock"`
}

func ToLineMessage(p ProductSchema) string {
	return fmt.Sprintf("%s   x%d\n", p.ProductName, p.InStock)
}
