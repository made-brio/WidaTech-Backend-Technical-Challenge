package routes

import (
	"database/sql"
	"widatech-technical-challenge/internal/controllers"
	"widatech-technical-challenge/internal/service"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, db *sql.DB) {

	// Initialize Services
	invoiceService := service.NewInvoiceService(db)
	importService := service.NewImportService(db)

	// Invoice
	invoiceController := controllers.NewInvoiceController(invoiceService)
	invoiceRoutes := router.Group("/api/invoice")
	{
		invoiceRoutes.POST("/", invoiceController.CreateInvoice)
		invoiceRoutes.GET("/", invoiceController.GetInvoice)
		invoiceRoutes.PUT("/", invoiceController.UpdateInvoice)
		invoiceRoutes.DELETE("/:invoiceno", invoiceController.DeleteInvoice)
	}
	// XLSX Import Routes
	xlsxController := controllers.NewImportController(importService) // Assuming you have an XLSX controller
	xlsxRoutes := router.Group("/api/xlsx")
	{
		// Route to upload an XLSX file
		xlsxRoutes.POST("/import", xlsxController.ImportInvoices) // Handles the XLSX import
	}
}
