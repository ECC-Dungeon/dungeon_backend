package controllers

import (
	"game/middlewares"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NextFloor(ctx echo.Context) error {
	// チームを取得
	team := ctx.Get("team").(middlewares.Team)

	
	
	return ctx.String(http.StatusOK, "Hello, World!")
}