package main

import (
	"fmt"
	"strconv"
)

type Product struct {
	ID    string
	Name  string
	Price float64
	Stock int
}

var products = []Product{
	{ID: "p1", Name: "Laptop", Price: 999, Stock: 15},
	{ID: "p2", Name: "Phone", Price: 699, Stock: 30},
}

func addProduct(name string, price float64, stock int) Product {
	id := "p" + strconv.Itoa(len(products)+1)
	newProduct := Product{
		ID:    id,
		Name:  name,
		Price: price,
		Stock: stock,
	}
	products = append(products, newProduct)
	return newProduct
}

func getLowStockProducts(threshold int) []Product {
	var lowStock []Product
	for _, p := range products {
		if p.Stock < threshold {
			lowStock = append(lowStock, p)
		}
	}
	return lowStock
}

func updateStock(id string, quantity int) *Product {
	for i := range products {
		if products[i].ID == id {
			products[i].Stock += quantity
			return &products[i]
		}
	}
	return nil
}

func discontinueProduct(id string) {
	for i, p := range products {
		if p.ID == id {
			products = append(products[:i], products[i+1:]...)
			return
		}
	}
}

func printProducts(header string, ps []Product) {
	fmt.Println(header)
	for _, p := range ps {
		fmt.Printf("ID: %s, Name: %s, Price: $%.2f, Stock: %d\n",
			p.ID, p.Name, p.Price, p.Stock)
	}
}

func main() {
	// Test the CRUD operations
	fmt.Println("Initial products:")
	printProducts("", products)

	// CREATE - Add new products
	fmt.Println("\nAdding new products:")
	fmt.Println(addProduct("Tablet", 299, 8))
	fmt.Println(addProduct("Headphones", 99, 25))

	// READ - Get all products and low stock items
	fmt.Println("\nAll products:")
	printProducts("", products)
	fmt.Println("\nLow stock products (threshold: 10):")
	printProducts("", getLowStockProducts(10))

	// UPDATE - Update stock
	fmt.Println("\nUpdating stock:")
	// Sell 5 laptops
	if updated := updateStock("p1", -5); updated != nil {
		fmt.Printf("Updated Laptop stock: %d\n", updated.Stock)
	}
	// Add 10 tablets
	if updated := updateStock("p3", 10); updated != nil {
		fmt.Printf("Updated Tablet stock: %d\n", updated.Stock)
	}

	// DELETE - Discontinue a product
	fmt.Println("\nDiscontinuing product p2:")
	discontinueProduct("p2")
	fmt.Println("Remaining products:")
	printProducts("", products)
}
