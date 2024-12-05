package controllers

import (
	"admin/middlewares"
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

type CreateGameArgs struct {
	Name string `json:"name"`
}

func CreateGame(ctx echo.Context) error {
	// ゲームを作成する
	args := new(CreateGameArgs)
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
			"msg":    "ゲーム名を入力してください",
		})
	}

	// ユーザーを取得する
	user := ctx.Get("user").(middlewares.UserData)

	// ゲームを作成する
	game, err := services.CreateGame(args.Name,user.UserID)

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
		"msg":    game,
	})
}

func GetGames(ctx echo.Context) error {
	// ゲームを取得する
	games, err := services.GetGames()

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
		"msg":    games,
	})
}

func DeleteGame(ctx echo.Context) error {
	// ゲームID を取得する
	gameid := ctx.Request().Header.Get("gameid")

	// バリデーション
	if gameid == "" {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"result": "error",
			"msg":    "ゲームIDを入力してください",
		})
	}

	// チームを削除する
	err := services.DeleteGame(gameid)

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


func GetFloor(ctx echo.Context) error {
	// チームID を取得する
	gameid := ctx.Request().Header.Get("gameid")

	// バリデーション
	if gameid == "" {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"result": "error",
			"msg":    "ゲームIDを入力してください",
		})
	}

	// チームを取得する
	floor, err := services.GetFloor(gameid)

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
		"msg":    floor,
	})
}

func SetFloor(ctx echo.Context) error {
	// チームID を取得する
	gameid := ctx.Request().Header.Get("gameid")

	// バリデーション
	if gameid == "" {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"result": "error",
			"msg":    "ゲームIDを入力してください",
		})
	}

	// フロア取得
	args := new(services.SetFloorArgs)
	if err := ctx.Bind(args); err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"result": "error",
			"msg":    err.Error(),
		})
	}

	// フロアを設定する
	err := services.SetFloor(gameid, args.Floors)

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

func GetGameInfo(ctx echo.Context) error {
	// チーム取得
	team := ctx.Get("team").(models.Team)

	// ゲームID を取得する
	game,err := services.GetGame(team.GameID)

	// エラー処理
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"result": "error",
			"msg":    err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"result": "success",
		"msg":    game,
	})
}

func StartGame(ctx echo.Context) error {
	// ゲームID を取得する
	gameid := ctx.Request().Header.Get("gameid")

	// ゲームを開始する
	err := services.StartGame(gameid)

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

func StopGame(ctx echo.Context) error {
	// ゲームID を取得する
	gameid := ctx.Request().Header.Get("gameid")

	// ゲームを停止する
	err := services.StopGame(gameid)

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

// スマホ側でゲームを開始する
func StartGame2(ctx echo.Context) error {
	// チームを取得
	team := ctx.Get("team").(models.Team)

	// ゲームを取得する
	game,err := models.GetGame(team.GameID)

	// エラー処理
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"result": "error",
			"msg":    err.Error(),
		})
	}

	// ゲームが開始していないばあい
	if game.Status != models.Started {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"result": "error",
			"msg":    "ゲームは開始していません",
		})
	}

	// ゲームを開始する
	err = services.StartGame2(team.TeamID)

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