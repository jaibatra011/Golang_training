package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	sql := "root:admin123@tcp(localhost:3306)/golang"
	db, err := gorm.Open(mysql.Open(sql), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.Debug().AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	db.First(&product, 1) // find product with integer primary key
	fmt.Println("                         ", product)
	db.First(&product, "code = ?", "D42") // find product with code D42
	fmt.Println("    ", product)
	fmt.Println(product.Code)

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	fmt.Println(product)
	// Delete - delete product
	//db.Where("Code=?","F42").Delete(&product)
	//db.Delete(&product, 1)

	fmt.Println("end")
}
