package services

import (
	"admin/models"
	"admin/utils"
)

func CreateGame(name string, creatorID string) (string, error) {
	// 作成
	game, err := models.CreateGame(name, creatorID)

	// エラー処理
	if err != nil {
		return "", err
	}

	return game.GameID, nil
}

func GetGames() ([]models.Game, error) {
	// ゲームを取得する
	games, err := models.GetGames()

	// エラー処理
	if err != nil {
		return nil, err
	}

	return games, nil
}

func DeleteGame(gameid string) error {
	// ゲーム取得
	game, err := models.GetGame(gameid)

	// エラー処理
	if err != nil {
		return err
	}

	// ゲームを削除する
	err = game.Delete()

	// エラー処理
	if err != nil {
		return err
	}

	return nil
}

func GetFloor(gameid string) ([]models.Floors, error) {
	// チームを取得する
	game, err := models.GetGame(gameid)

	// エラー処理
	if err != nil {
		return []models.Floors{}, err
	}

	// フロアを取得する
	floors, err := game.GetFloors()

	// エラー処理
	if err != nil {
		return []models.Floors{}, err
	}

	return floors, nil
}


type SetFloorArgs struct {
	Floors []Floor `json:"floors"`
}

type Floor struct {
	FloorNum int `json:"floor"`
	Name string `json:"name"`
}

func SetFloor(gameid string, floors []Floor) error {
	// チームを取得する
	game, err := models.GetGame(gameid)

	// エラー処理
	if err != nil {
		return err
	}

	// フロアを回す
	for _,val := range floors {
		// フロアが1~7までしかない
		if val.FloorNum < 1 || val.FloorNum > 7 {
			continue
		}

		// フロアを追加する
		err = game.AddFloor(val.FloorNum,val.Name)

		// エラー処理
		if err != nil {
			utils.Println("フロア追加失敗 : " + err.Error())
			continue
		}
	}

	return nil
}