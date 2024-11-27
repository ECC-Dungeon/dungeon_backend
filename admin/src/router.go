package main

import (
	// "admin/controllers"
	"admin/controllers"
	"admin/middlewares"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitServer() *echo.Echo {
	// サーバー作成
	server := echo.New()

	// ミドルウェア
	server.Use(middleware.Logger())
	// server.Use(middleware.Recover())

	server.POST("/", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "Hello, World!")
	}, middlewares.PocketAuth())

	// ルーティング
	teamg := server.Group("/team")
	teamg.Use(middlewares.PocketAuth())
	{
		teamg.GET("/list", controllers.ListTeam)
		teamg.POST("/create", controllers.CreateTeam)
		teamg.DELETE("/delete", controllers.DeleteTeam)
	}
	return server
}
