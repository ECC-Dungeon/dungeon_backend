package models

import (
	"admin/utils"
)

type Status string

const (
	UnUsed = Status("unused")
	Used   = Status("used")
)

type Team struct {
	TeamID    string `gorm:"primaryKey"`
	Name      string //チーム名
	GameID    string //ゲームID
	Status    Status //チームステータス
	NickName  string //チーム名 (学生が決める)
	Creator   string //作成者ID
	CreatedAt int64  `gorm:"autoCreateTime"` //作成時間
}

func CreateTeam(name string, creatorId string, gameID string) (string, error) {
	// ID を作成する
	teamID := utils.GenID()

	// チームを作成する
	result := dbconn.Save(&Team{
		TeamID:    teamID,
		Name:      name,
		Status:    UnUsed,
		GameID:    gameID,
		NickName:  name,
		Creator:   creatorId,
		CreatedAt: utils.Now(),
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

func DeleteTeam(teamid string) error {
	// チームを削除する
	result := dbconn.Delete(&Team{
		TeamID: teamid,
	})

	// エラー処理
	if result.Error != nil {
		return result.Error
	}

	// チームのリンクを削除する
	result = dbconn.Delete(&GameLink{
		TeamID: teamid,
	})

	// エラー処理
	if result.Error != nil {
		return result.Error
	}

	// リンク用のトークンも削除する
	result = dbconn.Delete(&LinkToken{
		TeamID: teamid,
	})

	return result.Error
}

func ListTeam() ([]Team, error) {
	// チームを取得する
	var teams []Team

	// チームを取得
	result := dbconn.Find(&teams)

	// エラー処理
	if result.Error != nil {
		return nil, result.Error
	}

	return teams, nil
}

func UpdateNickName(teamid string, name string) (Team, error) {
	// チームを取得
	team, err := GetTeam(teamid)

	// エラー処理
	if err != nil {
		return Team{}, err
	}

	// チーム名を更新
	team.NickName = name

	// チームを更新
	result := dbconn.Save(&team)

	// エラー処理
	if result.Error != nil {
		return Team{}, result.Error
	}

	return team, nil
}