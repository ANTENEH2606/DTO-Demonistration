# Product Inventory Management in Go

This is a simple Go program demonstrating basic CRUD (Create, Read, Update, Delete) operations on a product inventory. It models products with fields like ID, name, price, and stock quantity and provides functions to manage the products list.

## Features

- **Create**: Add new products with unique IDs.
- **Read**: List all products and retrieve low stock items based on a threshold.
- **Update**: Modify stock quantities by product ID (e.g., selling or restocking).
- **Delete**: Remove (discontinue) products by ID.
- Simple console output displaying product details.

## Code Overview

- `Product` struct represents a product with ID, Name, Price, and Stock.
- `products` slice holds the in-memory list of products.
- `addProduct(name, price, stock)` adds a new product with an auto-generated ID.
- `getLowStockProducts(threshold)` returns products whose stock is below the given threshold.
- `updateStock(id, quantity)` updates stock quantity for a product by ID and returns the updated product.
- `discontinueProduct(id)` removes a product from the list by ID.
- `printProducts(header, ps)` prints product details to console.
- The `main()` function tests all CRUD operations with sample products and prints results to the console.

## Usage

1. Run the program by writing "go run main.go" in the therminal.


