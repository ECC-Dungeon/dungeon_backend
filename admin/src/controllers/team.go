package controllers

import (
	"admin/middlewares"
	"admin/services"

	"github.com/labstack/echo/v4"
)

func ListTeam(ctx echo.Context) error {
	// チームを取得する
	teams, err := services.ListTeam()

	// エラー処理
	if err != nil {
		return ctx.JSON(500, echo.Map{
			"result": "error",
			"msg":    err.Error(),
		})
	}

	// チームを返す
	return ctx.JSON(200, echo.Map{
		"result": "success",
		"msg":    teams,
	})
}

type CreateArgs struct {
	Name     string `json:"name"`
}

func CreateTeam(ctx echo.Context) error {
	// ユーザーを取得する
	user := ctx.Get("user").(middlewares.UserData)

	// チームを作成する
	args := new(CreateArgs)
	if err := ctx.Bind(args); err != nil {
		return ctx.JSON(500, echo.Map{
			"result": "error",
			"msg":    err.Error(),
		})
	}

	// チームを作成する (ゲームIDは固定)
	teamid, err := services.CreateTeam(args.Name, user.UserID, "f3f9577d-4b90-4180-a396-826ef9676348")

	// エラー処理
	if err != nil {
		return ctx.JSON(500, echo.Map{
			"result": "error",
			"msg":    err.Error(),
		})
	}

	// チームを返す
	return ctx.JSON(200, echo.Map{
		"result": "success",
		"msg":    teamid,
	})
}

type DeleteArgs struct {
	Teamid string `json:"teamid"`
}

func DeleteTeam(ctx echo.Context) error {
	// チームを削除する
	args := new(DeleteArgs)
	if err := ctx.Bind(args); err != nil {
		return ctx.JSON(500, echo.Map{
			"result": "error",
			"msg":    err.Error(),
		})
	}

	// チームを削除する
	err := services.DeleteTeam(args.Teamid)

	// エラー処理
	if err != nil {
		return ctx.JSON(500, echo.Map{
			"result": "error",
			"msg":    err.Error(),
		})
	}

	// 結果を返す
	return ctx.JSON(200, echo.Map{
		"result": "success",
	})
}