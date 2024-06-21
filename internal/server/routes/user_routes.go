package routes

import (
    "generator-go-code/internal/server/handler"
    "github.com/labstack/echo/v4"
)

func SetupUserRoutes(e *echo.Echo, UserHandler *handler.UserHandler) {
    e.POST("/user", UserHandler.CreateUser)
    e.GET("/user/list", UserHandler.GetAllUsers)
    e.GET("/user/:id", UserHandler.GetUserByID)
    e.PUT("/user/:id", UserHandler.UpdateUser)
    e.DELETE("/user/:id", UserHandler.DeleteUser)
}
