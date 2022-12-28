package controller

import (
	"net/http"

	"github.com/gin-contrib/cors"

	"github.com/rs/zerolog"

	"gitlab.com/alexandrevinet/leasing/business/usecases"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "gitlab.com/alexandrevinet/leasing/docs"
)

// NewRouter will register all routes in gin handler.
func NewRouter(handler *gin.Engine, l *zerolog.Logger, uc *usecases.UseCases) {
	// Options
	handler.Use(
		gin.Logger(),
		gin.Recovery(),
		cors.Default(),
	)

	// Swagger
	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	handler.GET("/swagger/*any", swaggerHandler)

	// K8s probe
	handler.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	// Prometheus metrics
	handler.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Routers
	h := handler.Group("/api/v1")
	{
		newCarRoutes(h, uc.Car, l)
		newCustomerRoutes(h, uc.Customer, l)
		newScheduleRoutes(h, uc.Schedule, l)
	}
}
