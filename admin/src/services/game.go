package services

import (
	"admin/gamerpc"
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
	FloorNum int    `json:"floorNum"`
	Checked  bool   `json:"checked"`
	Name     string `json:"name"`
}

func SetFloor(gameid string, floors []Floor) error {
	// チームを取得する
	game, err := models.GetGame(gameid)

	// エラー処理
	if err != nil {
		return err
	}

	// フロアをリセット
	err = game.ClearFloor()

	// エラー処理
	if err != nil {
		return err
	}

	// フロアを回す
	for _, val := range floors {
		// フロアが1~7までしかない
		if val.FloorNum < 1 || val.FloorNum > 7 {
			continue
		}

		// フロアを追加する
		err = game.AddFloor(val.FloorNum, val.Name, val.Checked)

		// エラー処理
		if err != nil {
			utils.Println("フロア追加失敗 : " + err.Error())
			continue
		}
	}

	return nil
}

// ゲームを取得する
func GetGame(gameid string) (models.Game, error) {
	// ゲームを取得する
	return models.GetGame(gameid)
}

func StartGame(gameid string) error {
	// ゲームを取得する
	game, err := models.GetGame(gameid)

	// エラー処理
	if err != nil {
		return err
	}

	// チームのリスト作成
	sendTeams := []*gamerpc.Team{}

	// チームを取得する
	teams, err := game.GetTeams()

	// エラー処理
	if err != nil {
		return err
	}

	// チームを回す
	for _, team := range teams {
		if team.Status != models.UnUsed {
			// 使用中のチームを追加
			sendTeams = append(sendTeams, &gamerpc.Team{
				Id:   team.TeamID,
				Name: team.Name,
			})
		}
	}

	// フロアも設定
	sendFloors := []*gamerpc.Floor{}

	// フロアを取得する
	floors,err := game.GetFloors()

	// エラー処理
	if err != nil {
		return err
	}

	// フロアを追加
	for _, floor := range floors {
		// 使用するフロアのみ

		if !floor.Enabled {
			continue
		}

		sendFloors = append(sendFloors, &gamerpc.Floor{
			Name: floor.Name,
			Num:  int32(floor.FloorNum),
		})
	}

	// GPRC 経由で送信
	err = gamerpc.StartGame(gameid, sendTeams, sendFloors)

	// エラー処理
	if err != nil {
		return err
	}

	// ゲームを開始する
	return game.Start()
}

func StopGame(gameid string) error {
	// ゲームを取得する
	game, err := models.GetGame(gameid)

	// エラー処理
	if err != nil {
		return err
	}

	// ゲームを停止する
	return game.Stop()
}

// スマホ側でゲームを開始する
func StartGame2(teamid string) error {
	// チームを取得する
	team, err := models.GetTeam(teamid)

	// エラー処理
	if err != nil {
		return err
	}

	// ゲームを開始する
	return team.StartGame()
}