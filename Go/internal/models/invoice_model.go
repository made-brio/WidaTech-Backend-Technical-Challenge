package models

import "time"

// Invoice represents the invoice table in the database.
type Invoice struct {
	ID              int       `json:"id" db:"id" binding:"required"`                             // Unique ID for the invoice
	InvoiceNo       string    `json:"invoice_no" db:"invoice_no" binding:"required"`             // Invoice number, required field
	Date            time.Time `json:"date" db:"date" binding:"required"`                         // Date of the invoice creation
	CustomerName    string    `json:"customer_name" db:"customer_name" binding:"required"`       // Name of the customer, required field
	SalespersonName string    `json:"salesperson_name" db:"salesperson_name" binding:"required"` // Name of the salesperson, required field
	PaymentType     string    `json:"payment_type" db:"payment_type" binding:"required"`         // Payment type (Enum: CASH | CREDIT)
	Notes           string    `json:"notes,omitempty" db:"notes"`                                // Optional field for additional notes
	Products        []Product `json:"products,omitempty" binding:"required"`       // List of products sold, stored in the product table
}
