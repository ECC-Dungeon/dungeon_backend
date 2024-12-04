package controllers

import (
	"admin/middlewares"
	"admin/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ListTeam(ctx echo.Context) error {
	// ゲームID を取得する
	gameid := ctx.Request().Header.Get("gameid")

	//バリデーション
	if gameid == "" {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"result": "error",
			"msg":    "ゲームIDを入力してください",
		})
	}

	// チームを取得する
	teams, err := services.ListTeam(gameid)

	// エラー処理
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"result": "error",
			"msg":    err.Error(),
		})
	}

	// チームを返す
	return ctx.JSON(http.StatusOK, echo.Map{
		"result": "success",
		"msg":    teams,
	})
}

type CreateArgs struct {
	Name   string `json:"name"`
	GameID string `json:"gameid"`
}

func CreateTeam(ctx echo.Context) error {
	// ユーザーを取得する
	user := ctx.Get("user").(middlewares.UserData)

	// チームを作成する
	args := new(CreateArgs)
	if err := ctx.Bind(args); err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"result": "error",
			"msg":    err.Error(),
		})
	}

	// バリデーション
	if args.Name == "" {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"result": "error",
			"msg":    "チーム名を入力してください",
		})
	}

	// チームを作成する (ゲームIDは固定)
	teamid, err := services.CreateTeam(args.Name, user.UserID, args.GameID)

	// エラー処理
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"result": "error",
			"msg":    err.Error(),
		})
	}

	// チームを返す
	return ctx.JSON(http.StatusOK, echo.Map{
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
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"result": "error",
			"msg":    err.Error(),
		})
	}

	// チームを削除する
	err := services.DeleteTeam(args.Teamid)

	// エラー処理
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"result": "error",
			"msg":    err.Error(),
		})
	}

	// 結果を返す
	return ctx.JSON(http.StatusOK, echo.Map{
		"result": "success",
	})
}
