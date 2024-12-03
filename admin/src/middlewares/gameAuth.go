package middlewares

import (
	"admin/services"
	"admin/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GameTokenAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			// トークンを取得する
			token := ctx.Request().Header.Get("Authorization")

			// トークンがない場合
			if token == "" {
				return ctx.NoContent(http.StatusUnauthorized)
			}

			// トークンを検証する
			team, err := services.VerifyGameToken(token)

			// エラー処理
			if err.Err != nil {
				utils.Println(err.Error)
				return ctx.NoContent(err.Code)
			}

			// ユーザーを設定
			ctx.Set("team", team)

			// 次を実行
			return next(ctx)
		}
	}
}