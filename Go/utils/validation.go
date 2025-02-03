package utils

import (
	"errors"
	"strings"
	"widatech-technical-challenge/internal/models"
)

func ValidateInvoiceFields(invoice models.Invoice) error {
	var validationErrors []string

	// Validate individual fields and append error messages.
	if len(invoice.InvoiceNo) < 1 {
		validationErrors = append(validationErrors, "invoice_no must have at least 1 character")
	}
	if invoice.Date.IsZero() {
		validationErrors = append(validationErrors, "date is required")
	}
	if len(invoice.CustomerName) < 2 {
		validationErrors = append(validationErrors, "customer_name must have at least 2 characters")
	}
	if len(invoice.SalespersonName) < 2 {
		validationErrors = append(validationErrors, "salesperson_name must have at least 2 characters")
	}
	if err := ValidateInvoicePaymentType(invoice); err != nil {
		validationErrors = append(validationErrors, err.Error())
	}
	if invoice.Notes != "" && len(invoice.Notes) < 5 {
		validationErrors = append(validationErrors, "notes must have at least 5 characters if provided")
	}
	if len(invoice.Products) == 0 {
		validationErrors = append(validationErrors, "products list cannot be empty")
	}

	// If there are validation errors, join them into a single string with newlines and return.
	if len(validationErrors) > 0 {
		return errors.New(strings.Join(validationErrors, ";\n"))
	}
	return nil
}

// ValidateInvoicePaymentType ensures that the PaymentType of an Invoice is valid.
func ValidateInvoicePaymentType(invoice models.Invoice) error {
	validPaymentTypes := map[string]bool{
		"CASH":   true,
		"CREDIT": true,
	}
	if !validPaymentTypes[invoice.PaymentType] {
		return errors.New("invalid payment type: must be 'CASH' or 'CREDIT'")
	}
	return nil
}

// validateProduct checks that the product meets the specified requirements.
func ValidateProduct(product models.Product) error {
	if len(product.ItemName) < 5 {
		return errors.New("item_name must have at least 5 characters")
	}
	if product.Quantity < 1 {
		return errors.New("quantity must be at least 1")
	}
	if product.TotalCost < 0 {
		return errors.New("total_cost must be non-negative")
	}
	if product.TotalPrice < 0 {
		return errors.New("total_price must be non-negative")
	}
	return nil
}
