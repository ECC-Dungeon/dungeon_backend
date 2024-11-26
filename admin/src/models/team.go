package models

import (
	"admin/utils"
	"time"
)

type Status string

const (
	UnUsed = Status("unused")
	Used   = Status("used")
)

type Team struct {
	TeamID    string    `gorm:"primaryKey"`
	Name      string    //チーム名
	GameID    string    //ゲームID
	Status    Status    //チームステータス
	Creator   string    //作成者ID
	CreatedAt time.Time //作成時間
}

func CreateTeam(name string, creatorId string, gameID string) (string, error) {
	// ID を作成する
	teamID := utils.GenID()

	// チームを作成する
	result := dbconn.Save(&Team{
		TeamID:  teamID,
		Name:    name,
		Status:  UnUsed,
		GameID:  gameID,
		Creator: creatorId,
	})

	// エラー処理
	if result.Error != nil {
		return "", result.Error
	}

	return teamID, nil
}

func GetTeam(teamid string) (Team, error) {
	// チーム取得
	getTeam := Team{}

	// データベースから取得
	result := dbconn.Where(&Team{
		TeamID: teamid,
	}).First(&getTeam)

	// エラー処理
	if result.Error != nil {
		return Team{}, result.Error
	}

	return getTeam, nil
}
