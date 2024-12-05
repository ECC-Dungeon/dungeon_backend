package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

type FloorRsult struct {
	Result string `json:"result"`
	Msg    Floor  `json:"msg"`
}

type Floor struct {
	GameID   string `json:"GameID"`
	FloorNum int    `json:"FloorNum"`
	Name     string `json:"Name"`
	Enabled  bool   `json:"Enabled"`
}

func GetTeam(token string) (Floor, error) {
	// リクエスト送信
	req, _ := http.NewRequest("GET", os.Getenv("FLOOR_URL"), nil)

	// トークンを追加する
	req.Header.Set("Authorization", token)

	// リクエストを送信する
	client := new(http.Client)
	resp, err := client.Do(req)

	// エラー処理
	if err != nil {
		return Floor{}, err
	}

	defer resp.Body.Close()

	// エラー処理
	if resp.StatusCode != 200 {
		return Floor{}, errors.New(fmt.Sprint("Error: status code", resp.StatusCode))
	}

	var bind_data FloorRsult
	// Struct にバインドする
	body, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &bind_data); err != nil {
		return Floor{}, err
	}

	// return bind_data.Result, nil
	return bind_data.Msg, nil
}
