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
    token,err := services.GenToken(args.Teamid, user.UserID)

    // エラー処理
    if err != nil {
        utils.Println("リンク作成失敗 : " + err.Error())
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