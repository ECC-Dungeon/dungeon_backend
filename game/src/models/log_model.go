package models

import "game/utils"

type LogModel struct {
	GameID    string `gorm:"primary_key"` // ゲームID
	TeamID    string `gorm:"primary_key"` // チームID
	FloorNum  int    `gorm:"primary_key"` // フロア番号
	CreatedAt int64  //タイムスタンプ
}

func CreateLog(gameID string, teamID string, floorNum int) error {
	// DBに保存
	retult := dbconn.Save(&LogModel{GameID: gameID, TeamID: teamID, FloorNum: floorNum, CreatedAt: utils.Now()})

	return retult.Error
}
