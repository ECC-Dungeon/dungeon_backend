package services

import "admin/models"

func CreateTeam(name string, creatorId string, gameID string) (string, error) {
	// チームを作成
	teamid,err := models.CreateTeam(name, creatorId, gameID)

	// エラー処理
	if err != nil {
		return "", err
	}

	return teamid, nil
}

func DeleteTeam(teamid string) error {
	// チームを削除
	err := models.DeleteTeam(teamid)

	// エラー処理
	if err != nil {
		return err
	}

	return nil
}

func ListTeam() ([]models.Team, error) {
	// チームを取得
	teams, err := models.ListTeam()

	// エラー処理
	if err != nil {
		return nil, err
	}

	return teams, nil
}

func UpdateNickName(teamid string, name string) (models.Team, error) {
	// チームを更新
	team, err := models.UpdateNickName(teamid, name)

	// エラー処理
	if err != nil {
		return models.Team{}, err
	}

	return team, nil
}