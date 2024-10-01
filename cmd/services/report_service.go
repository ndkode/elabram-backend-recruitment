package services

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/ndkode/elabram-backend-recruitment/cmd/configs"
	"github.com/ndkode/elabram-backend-recruitment/cmd/repositories"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type reportService struct {
	Repo repositories.ReportRepository
}

type ReportService interface {
	GenerateProductReport(ctx *gin.Context, isOptimized bool) (map[string]interface{}, error)
}

func NewReportService(repo repositories.ReportRepository) *reportService {
	return &reportService{Repo: repo}
}

func (s *reportService) GenerateProductReport(ctx *gin.Context, isOptimized bool) (map[string]interface{}, error) {
	rdb := configs.ClientRedis()
	// Try to fetch the cached report
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))
	key := fmt.Sprintf("product_report_%d_page_size%d", page, pageSize)
	cachedReport, err := rdb.Get(ctx, key).Result()

	if err == redis.Nil || cachedReport == "" { // Cache miss, regenerate report
		fmt.Print("Cache miss, regenerate report isOptimized:", isOptimized, "\n")
		var (
			report map[string]interface{}
			err    error
		)
		if isOptimized {
			report, err = s.Repo.GenerateProductReportWithGoroutines(ctx)
		} else {
			report, err = s.Repo.GenerateProductReport(ctx)
		}

		if err != nil {
			return nil, err
		}
		// Cache the report for 5 minutes
		reportJson, err := json.Marshal(report)
		if err != nil {
			return nil, err
		}
		rdb.Set(ctx, key, string(reportJson), 5*time.Second)
		return report, err
	} else {
		fmt.Print("Cache hit\n")
		var cachedReportMap map[string]interface{}
		err = json.Unmarshal([]byte(cachedReport), &cachedReportMap)
		if err != nil {
			return nil, err
		}
		report := cachedReportMap // Use cached report
		return report, err

	}
}
