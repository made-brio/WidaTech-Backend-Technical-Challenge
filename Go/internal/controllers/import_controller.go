package controllers

import (
	"net/http"
	"widatech-technical-challenge/internal/service"

	"github.com/gin-gonic/gin"
)

// ImportController defines the controller layer for importing operations
type ImportController struct {
	ImportService *service.ImportService
}

// NewImportController creates a new ImportController instance
func NewImportController(importService *service.ImportService) *ImportController {
	return &ImportController{ImportService: importService}
}

// ImportInvoices handles the import of invoices and products from an XLSX file
func (ic *ImportController) ImportInvoices(ctx *gin.Context) {
	// Parse the XLSX file from the request
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}

	// Open the uploaded file
	f, err := file.Open()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
		return
	}
	defer f.Close()

	// Process the file using the service layer
	errors, err := ic.ImportService.ProcessXLSXFile(f)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process file"})
		return
	}

	// Respond with any validation or processing errors
	if len(errors) > 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": errors})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "File imported successfully"})
}
