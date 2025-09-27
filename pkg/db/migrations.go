package db

import (
	"log"
	"schoolmanagement/internal/order"
	"schoolmanagement/internal/product"
	"schoolmanagement/internal/user"
)

func Migrate() {
	err := Database.AutoMigrate(
		&user.User{},
		&product.Product{},
		&order.Order{},
	)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	log.Println("Database migration completed")
}
