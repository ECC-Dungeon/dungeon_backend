package controllers

import (
	"admin/models"
	"admin/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

// ゲーム用のトークンを検証する エンドポイント
func CheckToken(ctx echo.Context) error {
	// チームを取得する
	team := ctx.Get("team").(models.Team)

	// チームを返す
	return ctx.JSON(http.StatusOK, echo.Map{
		"result": "success",
		"msg":    team,
	})
}

type UpdateTeamNameArgs struct {
	Name string `json:"name"`
}

// チーム名を更新する エンドポイント
func UpdateTeamName(ctx echo.Context) error {
	// チームを取得する
	team := ctx.Get("team").(models.Team)

	// チームを更新する
	args := new(UpdateTeamNameArgs)
	if err := ctx.Bind(args); err != nil {
		return ctx.JSON(500, echo.Map{
			"result": "error",
			"msg":    err.Error(),
		})
	}

	// バリエーション
	if args.Name == "" {
		return ctx.JSON(400, echo.Map{
			"result": "error",
			"msg":    "チーム名を入力してください",
		})
	}

	// チームを更新する
	team, err := services.UpdateNickName(team.TeamID, args.Name)
	if err != nil {
		return ctx.JSON(500, echo.Map{
			"result": "error",
			"msg":    err.Error(),
		})
	}

	// チームを返す
	return ctx.JSON(http.StatusOK, echo.Map{
		"result": "success",
		"msg":    team,
	})
}

func AdminGameStart(ctx echo.Context) error {
	// ゲームを開始する
	err := services.AdminGameStart()
	if err != nil {
		return ctx.JSON(500, echo.Map{
			"result": "error",
			"msg":    err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"result": "success",
	})
}

// 管理者側でゲームを終了するエンドポイント
func AdminGameStop(ctx echo.Context) error {
	// ゲームを終了する
	err := services.AdminGameStop()
	if err != nil {
		return ctx.JSON(500, echo.Map{
			"result": "error",
			"msg":    err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"result": "success",
	})
}

// スマホ側でゲームを開始するエンドポイント
func MobileGameStart(ctx echo.Context) error {
	// ゲームが開始されているか
	isStarted, err := services.IsGameStarted()

	// エラー処理
	if err != nil {
		return ctx.JSON(500, echo.Map{
			"result": "error",
			"msg":    err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"result": "success",
		"msg":    isStarted,
	})
}

// 仕様フロアを取得するエンドポイント
func UseFloors(ctx echo.Context) error {
	// フロアを使用する
	floors, err := services.GetUseFloors()
	if err != nil {
		return ctx.JSON(500, echo.Map{
			"result": "error",
			"msg":    err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"result": "success",
		"msg":    floors,
	})
}

type SetFloorsArgs struct {
	Floors []int `json:"floors"`
}

func SetFloors(ctx echo.Context) error {
	//bind
	args := new(SetFloorsArgs)
	if err := ctx.Bind(args); err != nil {
		return ctx.JSON(500, echo.Map{
			"result": "error",
			"msg":    err.Error(),
		})
	}

	// 使用する階を設定
	err := services.SetUseFloors(args.Floors)

	// エラー処理
	if err != nil {
		return ctx.JSON(500, echo.Map{
			"result": "error",
			"msg":    err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"result": "success",
	})
}
