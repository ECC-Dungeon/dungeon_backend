package services

import "admin/models"

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