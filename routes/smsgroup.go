package routes

import (
	"sendsmsgroup-producer/controllers"

	"github.com/labstack/echo/v4"
)

func SMSGroupRoute(e *echo.Echo) {
	e.GET("/healthcheck", controllers.HealthCheck)
	e.POST("/api/v1/sendsmsgroup", controllers.AddSMSQueue)
}
