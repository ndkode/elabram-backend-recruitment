package controllers

import (
	"net/http"

	"github.com/ndkode/elabram-backend-recruitment/cmd/services"

	"github.com/gin-gonic/gin"
)

type reportController struct {
	Service services.ReportService
}

type ReportController interface {
	GetProductReport(ctx *gin.Context)
}

func NewReportController(service services.ReportService) *reportController {
	return &reportController{Service: service}
}

func (c *reportController) GetProductReport(ctx *gin.Context) {
	isOptimized := ctx.DefaultQuery("is_optimized", "false") == "true"
	// Generate the report
	report, err := c.Service.GenerateProductReport(ctx, isOptimized)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Return the report
	ctx.JSON(http.StatusOK, report)
}
