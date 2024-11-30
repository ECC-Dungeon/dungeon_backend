package models

import "admin/utils"

type GameLink struct {
	TeamID     string `gorm:"primaryKey"` //チームID
	TokenID    string //トークンID
	ExpiryDate int64  //有効期限 (unix time)
	CreatedAt  int64  //作成日
}

func CreateGameLink(teamid string, tokenid string, expiryDate int64) error {
	// チームを取得する
	team, err := GetTeam(teamid)

	// エラー処理
	if err != nil {
		return err
	}

	// リンクを作成する
	result := dbconn.Save(&GameLink{
		TeamID:     team.TeamID,
		TokenID:    tokenid,
		ExpiryDate: expiryDate,
		CreatedAt:  utils.Now(),
	})

	return result.Error
}

func DeleteGameLink(teamid string) error {
	// リンクを削除する
	result := dbconn.Where(&GameLink{
		TeamID: teamid,
	}).Delete(&GameLink{})

	return result.Error
}
