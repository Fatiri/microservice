package models

//Product for migration table in database
type Product struct {
	ProductID   uint64 `gorm:"primary_key; auto_increment; unique" json:"product_id"`
	ProductName string `gorm:"not null; type:varchar(255);" json:"product_name,omitempty"`
	Price       string `gorm:"not null; type:varchar(255);" json:"price,omitempty"`
}
