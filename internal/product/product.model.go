package product
type Product struct{
	ID int `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Price int `json:"price"`
}