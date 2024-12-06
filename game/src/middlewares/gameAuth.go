package middlewares

import (
	"encoding/json"
	"errors"
	"fmt"
	"game/utils"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type ReturnGameResult struct {
	Result string `json:"result"`
	Msg    Game   `json:"msg"`
}

type ReturnTeamResult struct {
	Result string `json:"result"`
	Msg    Team   `json:"msg"`
}

type Game struct {
	GameID    string
	Name      string
	CreatorID string
	Status    string
	CreatedAt int64
}

type Team struct {
	TeamID    string 
	Name      string
	GameID    string
	Status    string
	NickName  string
	Creator   string
	CreatedAt int64
}

type ReturnFloorResult struct {
	Result string  `json:"result"`
	Msg    []Floor `json:"msg"`
}

type Floor struct {
	GameID   string
	FloorNum int
	Name     string
	Enabled  bool
}

func GameAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			// ヘッダ取得
			token := ctx.Request().Header.Get("Authorization")

			// トークンを検証する
			game, err := GetGame(token)

			// エラー処理
			if err != nil {
				utils.Println(err)
				return ctx.NoContent(http.StatusUnauthorized)
			}

			// トークンを検証する
			team, err := GetTeam(token)

			// エラー処理
			if err != nil {
				utils.Println(err)
				return ctx.NoContent(http.StatusUnauthorized)
			}

			// フロアを取得
			floors,err := GetFloor(game.GameID)

			// エラー処理
			if err != nil {
				utils.Println(err)
				return ctx.NoContent(http.StatusUnauthorized)
			}

			// ゲームを設定
			ctx.Set("game", game)

			// チームを設定
			ctx.Set("team", team)

			// フロアを設定
			ctx.Set("floors", floors)

			// トークンを設定
			ctx.Set("token", token)

			return next(ctx)
		}
	}
}

func GetGame(token string) (Game, error) {
	// リクエスト送信
	req, _ := http.NewRequest("GET", os.Getenv("ADMIN_URL"), nil)

	// トークンを追加する
	req.Header.Set("Authorization", token)

	// リクエストを送信する
	client := new(http.Client)
	resp, err := client.Do(req)

	// エラー処理
	if err != nil {
		return Game{}, err
	}

	defer resp.Body.Close()

	// エラー処理
	if resp.StatusCode != 200 {
		return Game{}, errors.New(fmt.Sprint("Error: status code", resp.StatusCode))
	}

	var bind_data ReturnGameResult
	// Struct にバインドする
	body, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &bind_data); err != nil {
		return Game{}, err
	}

	// return bind_data.Result, nil
	return bind_data.Msg, nil
}

func GetTeam(token string) (Team, error) {
	// リクエスト送信
	req, _ := http.NewRequest("GET", os.Getenv("TEAM_URL"), nil)

	// トークンを追加する
	req.Header.Set("Authorization", token)

	// リクエストを送信する
	client := new(http.Client)
	resp, err := client.Do(req)

	// エラー処理
	if err != nil {
		return Team{}, err
	}

	defer resp.Body.Close()

	// エラー処理
	if resp.StatusCode != 200 {
		return Team{}, errors.New(fmt.Sprint("Error: status code", resp.StatusCode))
	}

	var bind_data ReturnTeamResult
	// Struct にバインドする
	body, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &bind_data); err != nil {
		return Team{}, err
	}

	// return bind_data.Result, nil
	return bind_data.Msg, nil
}

func GetFloor(gameid string) ([]Floor, error) {
	// リクエスト送信
	req, _ := http.NewRequest("GET", os.Getenv("FLOOR_URL"), nil)

	// トークンを追加する
	req.Header.Set("gameid", gameid)

	// リクエストを送信する
	client := new(http.Client)
	resp, err := client.Do(req)

	// エラー処理
	if err != nil {
		return []Floor{}, err
	}

	defer resp.Body.Close()

	// エラー処理
	if resp.StatusCode != 200 {
		return []Floor{}, errors.New(fmt.Sprint("Error: status code", resp.StatusCode))
	}

	var bind_data ReturnFloorResult
	// Struct にバインドする
	body, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &bind_data); err != nil {
		return []Floor{}, err
	}

	// return bind_data.Result, nil
	return bind_data.Msg, nil
}
