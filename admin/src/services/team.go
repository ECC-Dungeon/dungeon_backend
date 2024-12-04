package services

import "admin/models"

func CreateTeam(name string, creatorId string, gameID string) (string, error) {
	// チームを作成
	team,err := models.CreateTeam(name, gameID, creatorId)

	// エラー処理
	if err != nil {
		return "", err
	}

	return team.TeamID, nil
}

func DeleteTeam(teamid string) error {
	// チームを取得
	team, err := models.GetTeam(teamid)

	// エラー処理
	if err != nil {
		return err
	}
	
	// チームを削除
	err = team.Delete()

	// エラー処理
	if err != nil {
		return err
	}

	return nil
}

// ゲームに所属するチームを取得
func ListTeam(gameid string) ([]models.Team, error) {
	// ゲーム取得
	game, err := models.GetGame(gameid)

	// エラー処理
	if err != nil {
		return nil, err
	}

	// チームを取得
	teams, err := game.GetTeams()

	// エラー処理
	if err != nil {
		return nil, err
	}

	return teams, nil
}

func UpdateNickName(teamid string, name string) (models.Team, error) {
	// チームを取得
	team, err := models.GetTeam(teamid)

	// エラー処理
	if err != nil {
		return models.Team{}, err
	}

	// チームを更新
	err = team.UpdateNickName(name)

	// エラー処理
	if err != nil {
		return models.Team{}, err
	}

	return team, nil
}