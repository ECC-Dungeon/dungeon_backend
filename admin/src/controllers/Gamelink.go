package controllers

import (
	"admin/models"
	"admin/services"
	"admin/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GenGameToken(ctx echo.Context) error {
    // チームID を取得する
    teamid := ctx.Request().Header.Get("teamid")

    // チームID を検証する
    if teamid == "" {
        return ctx.JSON(http.StatusBadRequest, echo.Map{
            "result": "error",
            "msg":    "チームIDを入力してください",
        })
    }

    // ゲーム用のトークンを作成する
    gameToken, result := services.GenGameLink(teamid)

    // エラー処理
    if result.Err != nil {
        utils.Println("ゲームトークン生成失敗 : " + result.Message)
        return ctx.JSON(result.Code, echo.Map{
            "result": "error",
            "msg":    result.Message,
        })
    }

    // リンクを返す
    return ctx.JSON(http.StatusOK, echo.Map{
        "result": "success",
        "msg":    utils.FormatResponse(gameToken),
    })
}

// トークン取得
type UnLinkArgs struct {
    Teamid string `json:"teamid"`
}

func UnLink(ctx echo.Context) error {
    // トークンを取得する
    args := new(UnLinkArgs)
    if err := ctx.Bind(args); err != nil {
        utils.Println("リンク解除失敗 : " + err.Error())
        return ctx.JSON(http.StatusBadRequest, echo.Map{
            "result": "error",
            "msg":    err.Error(),
        })
    }

    // リンクを解除する
    result := services.UnLink(args.Teamid)

    // エラー処理
    if result.Err != nil {
        utils.Println("リンク解除失敗 : " + result.Message)
        return ctx.JSON(result.Code, echo.Map{
            "result": "error",
            "msg":    result.Message,
        })
    }

    // レスポンスを返す
    return ctx.JSON(http.StatusOK, echo.Map{
        "result": "success",
    })
}

func GetTeam(ctx echo.Context) error {
    // チームを取得する
    team := ctx.Get("team").(models.Team)

    // チームを返す
    return ctx.JSON(http.StatusOK, echo.Map{
        "result": "success",
        "msg":    team,
    })
}