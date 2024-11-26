package models

import "time"

type Link struct {
	TeamID     string    `gorm:"primaryKey"` //チームID
	TokenID    string    //トークンID
	ExpiryDate int64     //有効期限 (unix time)
	CreatedAt  time.Time //作成日
}

func CreateLink(teamid string, tokenid string, expiryDate int64) error {
	// チームを取得する
	team,err := GetTeam(teamid)

	// エラー処理
	if err != nil {
		return err
	}

	// リンクを作成する
	result := dbconn.Save(&Link{
		TeamID:     team.TeamID,
		TokenID:    tokenid,
		ExpiryDate: expiryDate,
	})

	return result.Error
}
