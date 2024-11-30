package controllers

import (
	"admin/middlewares"
	"admin/services"
	"admin/utils"

	"github.com/labstack/echo/v4"
)

type CreateLinkArgs struct {
    Teamid string `json:"teamid"`
}

func GenToken(ctx echo.Context) error {
    // ユーザーを取得する
    user := ctx.Get("user").(middlewares.UserData)

    // チームIDを取得する
    args := new(CreateLinkArgs)
    if err := ctx.Bind(args); err != nil {
        utils.Println("リンク作成失敗 : " + err.Error())
        return ctx.JSON(400, echo.Map{
            "result": "error",
            "msg":    err.Error(),
        })
    }

    // リンクを作成する
    token,err := services.GenLinkToken(args.Teamid, user.UserID)

    // エラー処理
    if err != nil {
        utils.Println("リンク用トークン作成失敗 : " + err.Error())
        return ctx.JSON(500, echo.Map{
            "result": "error",
            "msg":    err.Error(),
        })
    }

    // リンクを返す
    return ctx.JSON(200, echo.Map{
        "result": "success",
        "msg":    token,
    })
}

// トークンを取得
type InitLinkArgs struct {
    Token string `json:"token"`
}

func InitLink(ctx echo.Context) error {
    // トークンを取得する
    args := new(InitLinkArgs)
    if err := ctx.Bind(args); err != nil {
        utils.Println("リンク初期化失敗 : " + err.Error())
        return ctx.JSON(400, echo.Map{
            "result": "error",
            "msg":    err.Error(),
        })
    }

    // ゲーム用のトークンを作成する
    gameToken, result := services.InitLink(args.Token)

    // エラー処理
    if result.Err != nil {
        utils.Println("ゲームトークン生成失敗 : " + result.Message)
        return ctx.JSON(result.Code, echo.Map{
            "result": "error",
            "msg":    result.Message,
        })
    }

    // リンクを返す
    return ctx.JSON(200, echo.Map{
        "result": "success",
        "msg":    gameToken,
    })
}

// トークン取得
type UnLinkArgs struct {
    Teamid string `json:"teamid"`
}

func UnLink(ctx echo.Context) error {
    // ユーザーを取得する
    user := ctx.Get("user").(middlewares.UserData)

    // トークンを取得する
    args := new(UnLinkArgs)
    if err := ctx.Bind(args); err != nil {
        utils.Println("リンク解除失敗 : " + err.Error())
        return ctx.JSON(400, echo.Map{
            "result": "error",
            "msg":    err.Error(),
        })
    }

    // リンクを解除する
    result := services.UnLink(args.Teamid, user.UserID)

    // エラー処理
    if result.Err != nil {
        utils.Println("リンク解除失敗 : " + result.Message)
        return ctx.JSON(result.Code, echo.Map{
            "result": "error",
            "msg":    result.Message,
        })
    }

    // レスポンスを返す
    return ctx.JSON(200, echo.Map{
        "result": "success",
    })
}