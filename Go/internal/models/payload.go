package models

import "time"

type InvoiceRequest struct {
	Page int       `json:"page" binding:"required"`
	Size int       `json:"size" binding:"required"`
	Date time.Time `json:"date" binding:"required"`
}

type UpdateInvoiceRequest struct {
	InvoiceNo       string    `json:"invoice_no" db:"invoice_no" `             // Invoice number, required field
	Date            time.Time `json:"date" db:"date" `                         // Date of the invoice creation
	CustomerName    string    `json:"customer_name" db:"customer_name" `       // Name of the customer, required field
	SalespersonName string    `json:"salesperson_name" db:"salesperson_name" ` // Name of the salesperson, required field
	PaymentType     string    `json:"payment_type" db:"payment_type" `         // Payment type (Enum: CASH | CREDIT)
	Notes           string    `json:"notes,omitempty" db:"notes"`              // Optional field for additional notes
}
