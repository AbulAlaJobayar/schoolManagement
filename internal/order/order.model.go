package order
import "schoolmanagement/internal/product"

type Order struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	ProductId int            `json:"productId" gorm:"not null"`
	Product   product.Product `gorm:"foreignKey:ProductId"` // link ProductId to Product
}