package controllers

import (
	"admin/models"
	"admin/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

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
    Name   string `json:"name"`
}

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
