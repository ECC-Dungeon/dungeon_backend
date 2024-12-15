package models

import (
	"game/utils"
	"log"
)

type LogModel struct {
	GameID    string `gorm:"primary_key"` // ゲームID
	TeamID    string `gorm:"primary_key"` // チームID
	FloorNum  int    `gorm:"primary_key"` // フロア番号
	CreatedAt int64  //タイムスタンプ
}

func CreateLog(gameID string, teamID string, floorNum int) error {
	// reconver 
	defer func () {
		if rec := recover(); rec != nil {
			log.Println("reconver")
		}
	}()

	// DBに保存
	retult := dbconn.Save(&LogModel{GameID: gameID, TeamID: teamID, FloorNum: floorNum, CreatedAt: utils.Now()})

	return retult.Error
}

// チームのログを取得する
func GetLogs(gameID string, teamID string) ([]LogModel, error) {
	var logs []LogModel
	retult := dbconn.Where(&LogModel{GameID: gameID, TeamID: teamID}).Find(&logs)

	return logs, retult.Error
}

// ゲーム内のフロア数を取得する
func CountFloors(gameID string,FloorNum int) (int64, error) {
	var count int64
	retult := dbconn.Where(&LogModel{GameID: gameID, FloorNum: FloorNum}).Count(&count)

	return count, retult.Error
}