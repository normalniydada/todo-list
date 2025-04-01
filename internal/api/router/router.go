package router

import (
	"github.com/labstack/echo/v4"
	"todo-list/internal/api/handlers"
)

func NewRouter(e *echo.Echo, handler handlers.TaskHandler) {
	api := e.Group("/api/task")
	{
		api.GET("/list", handler.List)
		api.GET("/:id", handler.Get)
		api.POST("/create", handler.Create)
		api.PATCH("/update/:id", handler.Update)
		api.DELETE("/delete/:id", handler.Delete)
	}
}
