package models

// Product represents the products table in the database.
type Product struct {
    ID         int     `json:"id" db:"id" binding:"required"`                   // Unique ID for the product
    InvoiceNo  string  `json:"invoice_no" db:"invoice_no" binding:"required"`   // Foreign key to invoice
    ItemName   string  `json:"item_name" db:"item_name" binding:"required,minLength=5"` // Name of the product, minLength: 5
    Quantity   int     `json:"quantity" db:"quantity" binding:"required,min=1"`   // Product quantity, minValue: 1
    TotalCost  float64 `json:"total_cost" db:"total_cost" binding:"required,min=0"` // Cost of the product sold, minValue: 0
    TotalPrice float64 `json:"total_price" db:"total_price" binding:"required,min=0"` // Price of the product sold, minValue: 0
}
