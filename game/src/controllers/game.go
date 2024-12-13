package controllers

import (
	"game/middlewares"
	"game/services"
	"game/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type NextArgs struct {
	ClearFloor int `json:"clear_floor"`
}

func Next(ctx echo.Context) error {
	// チームを取得
	team := ctx.Get("team").(middlewares.Team)
	floors := ctx.Get("floors").([]middlewares.Floor)

	// 引数を取得
	var args NextArgs
	if err := ctx.Bind(&args); err != nil {
		utils.Println(err)
		return ctx.NoContent(http.StatusBadRequest)
	}

	// ログを追加
	nextData, err := services.Next(team, floors, args.ClearFloor)

	// エラー処理
	if err != nil {
		utils.Println(err)
		return ctx.NoContent(http.StatusInternalServerError)
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"msg": nextData,
		"result": "success",
	})
}

func NowFloor(ctx echo.Context) error {
	// チームを取得
	team := ctx.Get("team").(middlewares.Team)

	// 現在の階を取得
	nowFloor,err := services.NowFloor(team)

	// エラー処理
	if err != nil {
		utils.Println(err)
		return ctx.NoContent(http.StatusInternalServerError)
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"msg":  nowFloor,
		"result": "success",
	})
}