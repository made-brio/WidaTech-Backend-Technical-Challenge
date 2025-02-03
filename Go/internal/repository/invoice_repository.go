package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"widatech-technical-challenge/internal/models"
	"widatech-technical-challenge/utils"
)

// CreateInvoice inserts a new invoice record into the database.
func CreateInvoice(db *sql.DB, invoice models.Invoice) error {
	// Start a transaction
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Validate the Invoice Fields before proceeding.
	if err = utils.ValidateInvoiceFields(invoice); err != nil {
		return err
	}

	// Check for duplicate invoice
	exists, err := CheckInvoiceExists(db, invoice.InvoiceNo)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("duplicate invoice number")
	}

	// Insert the invoice
	sqlQuery := `INSERT INTO invoices (invoice_no, date, customer_name, salesperson_name, payment_type, notes) 
	             VALUES ($1, $2, $3, $4, $5, $6)`
	_, err = db.Exec(sqlQuery, invoice.InvoiceNo, invoice.Date, invoice.CustomerName, invoice.SalespersonName, invoice.PaymentType, invoice.Notes)
	if err != nil {
		return err
	}

	// Validate the Products Fields before proceeding.
	for _, product := range invoice.Products {
		if err = utils.ValidateProduct(product); err != nil {
			return err
		}
	}

	// Insert associated products
	productQuery := `INSERT INTO products (invoice_no, item_name, quantity, total_cost, total_price)
	                 VALUES ($1, $2, $3, $4, $5)`
	for _, product := range invoice.Products {
		_, err = tx.Exec(productQuery, invoice.InvoiceNo, product.ItemName, product.Quantity, product.TotalCost, product.TotalPrice)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

// GetInvoices retrieves a list of invoices based on the provided parameters (date, size, page)
// It also calculates and returns the total profit and total cash transactions for the given date
func GetInvoices(db *sql.DB, payload models.InvoiceRequest) (invoices []models.Invoice, totalProfit, totalCash float64, err error) {
	sqlQuery := `SELECT invoice_no, date, customer_name, salesperson_name, payment_type, notes 
	             FROM invoices 
	             WHERE date = $1 
	             LIMIT $2 OFFSET $3`

	rows, err := db.Query(sqlQuery, payload.Date, payload.Size, (payload.Page-1)*(payload.Size))
	if err != nil {
		return nil, 0, 0, err
	}
	defer rows.Close()

	// Retrieve invoices
	for rows.Next() {
		var invoice models.Invoice
		if err := rows.Scan(&invoice.InvoiceNo, &invoice.Date, &invoice.CustomerName, &invoice.SalespersonName, &invoice.PaymentType, &invoice.Notes); err != nil {
			return nil, 0, 0, err
		}

		// Retrieve products for the current invoice
		productsQuery := `SELECT id, invoice_no, item_name, quantity, total_cost, total_price 
		                  FROM products 
		                  WHERE invoice_no = $1`
		productRows, err := db.Query(productsQuery, invoice.InvoiceNo)
		if err != nil {
			return nil, 0, 0, err
		}

		var products []models.Product
		for productRows.Next() {
			var product models.Product
			if err := productRows.Scan(&product.ID, &product.InvoiceNo, &product.ItemName, &product.Quantity, &product.TotalCost, &product.TotalPrice); err != nil {
				return nil, 0, 0, err
			}
			products = append(products, product)
		}
		productRows.Close()

		invoice.Products = products
		invoices = append(invoices, invoice)
	}

	// Calculate total profit and total cash transactions
	for _, inv := range invoices {
		for _, product := range inv.Products {
			totalProfit += product.TotalPrice - product.TotalCost
			if inv.PaymentType == "CASH" {
				totalCash += product.TotalPrice
			}
		}
	}

	return invoices, totalProfit, totalCash, nil
}

func UpdateInvoice(db *sql.DB, invoice models.UpdateInvoiceRequest) error {
	// Validate if at least one field is provided for the update
	if invoice.Date.IsZero() && invoice.CustomerName == "" && invoice.SalespersonName == "" && invoice.PaymentType == "" && invoice.Notes == "" {
		return errors.New("no fields to update")
	}

	// Dynamic query
	query := "UPDATE invoices SET"
	args := []interface{}{}
	argCount := 1

	if !invoice.Date.IsZero() {
		query += fmt.Sprintf(" date = $%d,", argCount)
		args = append(args, invoice.Date)
		argCount++
	}
	if invoice.CustomerName != "" {
		query += fmt.Sprintf(" customer_name = $%d,", argCount)
		args = append(args, invoice.CustomerName)
		argCount++
	}
	if invoice.SalespersonName != "" {
		query += fmt.Sprintf(" salesperson_name = $%d,", argCount)
		args = append(args, invoice.SalespersonName)
		argCount++
	}
	if invoice.PaymentType != "" {
		query += fmt.Sprintf(" payment_type = $%d,", argCount)
		args = append(args, invoice.PaymentType)
		argCount++
	}
	if invoice.Notes != "" {
		query += fmt.Sprintf(" notes = $%d,", argCount)
		args = append(args, invoice.Notes)
		argCount++
	}

	// Remove trailing comma and add WHERE clause
	query = strings.TrimSuffix(query, ",")
	query += fmt.Sprintf(" WHERE invoice_no = $%d", argCount)
	args = append(args, invoice.InvoiceNo)

	_, err := db.Exec(query, args...)
	return err
}

// DeleteInvoice removes an invoice from the database
func DeleteInvoice(db *sql.DB, invoiceNo string) (err error) {
	sqlQuery := `DELETE FROM invoices WHERE invoice_no = $1`
	_, err = db.Exec(sqlQuery, invoiceNo)
	return err
}

// CheckInvoiceExists checks if an invoice with the given invoice number already exists in the database.
func CheckInvoiceExists(db *sql.DB, invoiceNo string) (bool, error) {
	sqlQuery := `SELECT COUNT(1) FROM invoices WHERE invoice_no = $1`
	var count int
	err := db.QueryRow(sqlQuery, invoiceNo).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
