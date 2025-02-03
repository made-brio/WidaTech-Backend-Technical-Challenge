package controllers

import (
	"database/sql"
	"net/http"
	"widatech-technical-challenge/internal/models"
	"widatech-technical-challenge/internal/service"

	"github.com/gin-gonic/gin"
)

// InvoiceController defines the controller layer for invoice operations
type InvoiceController struct {
	InvoiceService *service.InvoiceService
}

// NewInvoiceController creates a new InvoiceController instance
func NewInvoiceController(invoiceService *service.InvoiceService) *InvoiceController {
	return &InvoiceController{InvoiceService: invoiceService}
}

// CreateInvoice handles the creation of a new invoice
func (ic *InvoiceController) CreateInvoice(ctx *gin.Context) {
	var invoice models.Invoice
	if err := ctx.ShouldBindJSON(&invoice); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Use the service layer to create the invoice
	if err := ic.InvoiceService.CreateInvoice(invoice); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create invoice"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Invoice created successfully", "invoice": invoice})
}

// GetInvoice retrieves a single invoice by ID and calculates the total cash and total profit
func (ic *InvoiceController) GetInvoice(ctx *gin.Context) {
	var payload models.InvoiceRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Retrieve invoice, total profit, and total cash from the service
	invoice, totalProfit, totalCash, err := ic.InvoiceService.GetInvoices(payload)
	if err == sql.ErrNoRows {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Invoice not found"})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve invoice"})
		return
	}

	// Return the invoice data along with total profit and total cash
	ctx.JSON(http.StatusOK, gin.H{
		"invoice":     invoice,
		"totalProfit": totalProfit,
		"totalCash":   totalCash,
	})
}

// UpdateInvoice updates an existing invoice
func (ic *InvoiceController) UpdateInvoice(ctx *gin.Context) {

	var invoice models.UpdateInvoiceRequest
	if err := ctx.ShouldBindJSON(&invoice); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := ic.InvoiceService.UpdateInvoice(invoice); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update invoice"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Invoice updated successfully", "invoice": invoice})
}

// DeleteInvoice deletes an invoice by invoice_no
func (ic *InvoiceController) DeleteInvoice(ctx *gin.Context) {
	invoice_no := ctx.Param("invoiceno")

	if err := ic.InvoiceService.DeleteInvoice(invoice_no); err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Invoice not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete invoice"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Invoice deleted successfully"})
}
