package product

type ProductSchema struct {
	Id          int    `gorm:"column:id" json:"id"`
	ProductName string `gorm:"column:product_name" json:"product_name"`
}

func ToLineMessage(p ProductSchema) string {
	return p.ProductName + "\n"
}
