package service

import (
	"database/sql"
	"fmt"
	"io"
	"strconv"
	"time"
	"widatech-technical-challenge/internal/models"
	"widatech-technical-challenge/internal/repository"

	"github.com/xuri/excelize/v2"
)

// ImportService defines the service layer for importing invoices and products
type ImportService struct {
	DB *sql.DB
}

// NewImportService creates a new ImportService instance
func NewImportService(db *sql.DB) *ImportService {
	return &ImportService{DB: db}
}

// ProcessXLSXFile processes and validates the uploaded XLSX file
func (is *ImportService) ProcessXLSXFile(file io.Reader) ([]map[string]string, error) {
	f, err := excelize.OpenReader(file)
	if err != nil {
		return nil, fmt.Errorf("failed to parse XLSX file: %w", err)
	}

	var errors []map[string]string

	// Process the "product_sold" sheet
	productRows, err := f.GetRows("product sold")
	if err != nil {
		return nil, fmt.Errorf("failed to read product_sold sheet: %w", err)
	}

	// Process the "invoice" sheet
	invoiceRows, err := f.GetRows("invoice")
	if err != nil {
		return nil, fmt.Errorf("failed to read invoice sheet: %w", err)
	}

	for i, row := range invoiceRows {
		if i == 0 {
			continue // Skip the header
		}
		invoiceID := row[0]
		var products []models.Product

		// Associate products with the corresponding invoice
		for j, productRow := range productRows {
			if j == 0 {
				continue // Skip the header
			}
			if productRow[0] == invoiceID {
				quantity, _ := strconv.Atoi(productRow[2])
				totalCost, _ := strconv.ParseFloat(productRow[3], 64)
				totalPrice, _ := strconv.ParseFloat(productRow[4], 64)

				product := models.Product{
					InvoiceNo:  productRow[0],
					ItemName:   productRow[1],
					Quantity:   quantity,
					TotalCost:  totalCost,
					TotalPrice: totalPrice,
				}
				products = append(products, product)
			}
		}

		err := validateAndInsertInvoice(row, products, is)
		if err != nil {
			errors = append(errors, map[string]string{
				"invoice_id": invoiceID,
				"error":      err.Error(),
			})
		}
	}

	return errors, nil
}

// validateAndInsertInvoice validates and inserts invoice data into the database
func validateAndInsertInvoice(row []string, products []models.Product, is *ImportService) error {
	if len(row) < 6 || row[0] == "" || row[1] == "" || row[2] == "" || row[3] == "" || row[4] == "" || row[5] == "" {
		return fmt.Errorf("missing required invoice fields")
	}

	parsedDate, err := time.Parse("02-01-06", row[1]) // Adjust the layout based on your date format
	if err != nil {
		return fmt.Errorf("invalid date format: %w", err)
	}

	exists, err := repository.CheckInvoiceExists(is.DB, row[0])
	if err != nil {
		return fmt.Errorf("error checking invoice duplication: %w", err)
	}
	if exists {
		return fmt.Errorf("duplicate invoice ID found")
	}

	//attach invoice
	invoice := models.Invoice{
		InvoiceNo:       row[0],
		Date:            parsedDate,
		CustomerName:    row[2],
		SalespersonName: row[3],
		PaymentType:     row[4],
		Notes:           row[5],
		Products:        products,
	}

	return repository.CreateInvoice(is.DB, invoice)
}
