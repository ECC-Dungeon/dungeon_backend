package main

import (
	// "admin/controllers"
	"admin/controllers"
	"admin/middlewares"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitServer() *echo.Echo {
	// サーバー作成
	server := echo.New()

	// ミドルウェア
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())

	server.POST("/", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "Hello, World!")
	}, middlewares.PocketAuth())

	// ルーティング
	teamg := server.Group("/team")
	teamg.Use(middlewares.PocketAuth())
	{
		// チームを作成する エンドポイント
		teamg.POST("/create", controllers.CreateTeam)   //確認済み

		// チームを削除する エンドポイント
		teamg.DELETE("/delete", controllers.DeleteTeam) //確認済み

		// ゲーム用のトークンを生成する エンドポイント
		teamg.POST("/link", controllers.GenGameToken)   //確認済み

		// ゲーム用のトークンを削除する エンドポイント
		teamg.DELETE("/unlink", controllers.UnLink)     //確認済み

		// ゲームのチーム一覧
		teamg.GET("/list", controllers.ListTeam) //確認済み
	}

	// チーム情報を取得するエンドポイント
	gameg := server.Group("/game")
	{
		// ゲーム一覧を取得するエンドポイント
		gameg.GET("/list", controllers.GetGames, middlewares.PocketAuth())

		// ゲームを作成するエンドポイント
		gameg.POST("/create", controllers.CreateGame, middlewares.PocketAuth()) //確認済み

		// チームを取得する
		gameg.GET("/team", controllers.GetTeam, middlewares.GameTokenAuth()) //確認済み

		// チーム名を更新する
		gameg.PUT("/tname", controllers.UpdateTeamName, middlewares.GameTokenAuth()) //確認済み

		// ゲームを削除する
		gameg.DELETE("/delete", controllers.DeleteGame, middlewares.PocketAuth()) //確認済み

		// floor を設定する
		gameg.POST("/floor", controllers.SetFloor, middlewares.GameTokenAuth())
	}

	return server
}
