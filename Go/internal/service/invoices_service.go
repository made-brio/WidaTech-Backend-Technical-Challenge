package service

import (
	"database/sql"
	"widatech-technical-challenge/internal/models"
	"widatech-technical-challenge/internal/repository"
)

// InvoiceService defines the service layer for invoice operations
type InvoiceService struct {
	DB *sql.DB
}

// NewInvoiceService creates a new InvoiceService instance
func NewInvoiceService(db *sql.DB) *InvoiceService {
	return &InvoiceService{DB: db}
}

// CreateInvoice creates a new invoice
func (is *InvoiceService) CreateInvoice(invoiceData models.Invoice) error {
	return repository.CreateInvoice(is.DB, invoiceData)
}

// GetInvoices retrieves a list of invoices
func (is *InvoiceService) GetInvoices(payload models.InvoiceRequest) ([]models.Invoice, float64, float64, error) {
	return repository.GetInvoices(is.DB, payload)
}

// UpdateInvoice updates an existing invoice
func (is *InvoiceService) UpdateInvoice(invoiceData models.UpdateInvoiceRequest) error {
	return repository.UpdateInvoice(is.DB, invoiceData)
}

// DeleteInvoice deletes an invoice by its invoice number
func (is *InvoiceService) DeleteInvoice(invoiceNo string) error {
	return repository.DeleteInvoice(is.DB, invoiceNo)
}
