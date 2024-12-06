package models

import (
	"game/middlewares"
	"game/utils"
	"log"
)

// 挑戦中のチームを記録するテーブル
type ChallengeTable struct {
	GameID    string `gorm:"primary_key"`
	TeamID    string `gorm:"primary_key"`
	FloorNum  int    `gorm:"primary_key"`
	CreatedAt int64  // タイムスタンプ
}

func CreateChallenge(gameID string, teamID string, floorNum int) error {
	// DBに保存
	retult := dbconn.Save(&ChallengeTable{GameID: gameID, TeamID: teamID, FloorNum: floorNum, CreatedAt: utils.Now()})

	return retult.Error
}

func DeleteChallenge(gameID string, teamID string, floorNum int) error {
	// DBに保存
	retult := dbconn.Delete(&ChallengeTable{GameID: gameID, TeamID: teamID, FloorNum: floorNum})

	return retult.Error
}

func GetChallenge(gameID string, teamID string) (ChallengeTable, error) {
	var challenge ChallengeTable
	retult := dbconn.Where(&ChallengeTable{GameID: gameID, TeamID: teamID}).First(&challenge)

	return challenge, retult.Error
}

// フロアをチャレンジしているチーム数を取得
func CountChallenge(gameID string,floorNum int) (int, error) {
	var challenges []ChallengeTable
	
	// DBからカウント
	retult := dbconn.Where(&ChallengeTable{GameID: gameID, FloorNum: floorNum}).Find(&challenges)

	return len(challenges), retult.Error
}


// チャレンジ中のチームが一番少ない階を返す
func GetLowFloor(gameid string,floors []middlewares.Floor) int {
	// 返すフロア
	lowFloor := 0

	// チャレンジ数保存用変数
	challenge_count := 0

	for _, val := range floors {
		// フロアが使用しない場合
		if !val.Enabled {
			continue
		}

		// フロアにいるチャレンジ数を取得
		challengeC,err := CountChallenge(gameid,val.FloorNum)

		// エラー処理
		if err != nil {
			log.Println("failed to count challenge : ",err)
			continue
		}

		// チャレンジ数が0の場合
		if challenge_count == 0 {
			// チャレンジ数を保存
			lowFloor = val.FloorNum
			challenge_count = challengeC
			continue
		}

		// 前のチャレンジ数が現在のチャレンジ数より大きい場合
		if challenge_count > challengeC {
			// チャレンジ数を保存
			lowFloor = val.FloorNum
			challenge_count = challengeC
			continue
		}
	}

	return lowFloor
}