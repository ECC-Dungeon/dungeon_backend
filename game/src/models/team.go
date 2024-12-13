package models

import (
	"encoding/json"
	"game/utils"
)

type TeamTable struct {
	GameID        string `gorm:"primary_key"`
	TeamID        string `gorm:"primary_key"`
	Challenges    string	 // 挑戦する階のリスト
	ClearedFloors string	 // クリアした階のリスト
	CreatedAt     int64	 // 作成日時
}

// チームを登録する
func RegisterTeam(gameID string, teamID string,Challenges []int) error {
	// json 文字列にする
	jsonChallenges, err := json.Marshal(Challenges)

	// エラー処理
	if err != nil {
		return err
	}

	// json 文字列にする
	StrChallenges := string(jsonChallenges)

	// DBに保存
	retult := dbconn.Save(&TeamTable{GameID: gameID, TeamID: teamID, Challenges: StrChallenges, CreatedAt: utils.Now(), ClearedFloors: "[]"})

	return retult.Error
}

// チームを削除する
func DeleteALLTeam(gameID string) error {
	utils.Println(gameID)
	// DBに保存
	retult := dbconn.Delete(&TeamTable{GameID: gameID})

	return retult.Error
}

func (team *TeamTable) GetChallenges() ([]int, error) {
	// json 文字列から戻す
	var Challenges []int
	err := json.Unmarshal([]byte(team.Challenges), &Challenges)

	return Challenges, err
}

func (team *TeamTable) UpdateChallenges(Challenges []int) error {
	// json 文字列にする
	jsonChallenges, err := json.Marshal(Challenges)

	// エラー処理
	if err != nil {
		return err
	}

	// json 文字列にする
	team.Challenges = string(jsonChallenges)

	// DBに保存
	result := dbconn.Model(&team).Update("Challenges", team.Challenges)

	return result.Error
}

func (team *TeamTable) GetClearedFloors() ([]int, error) {
	// json 文字列から戻す
	var clearedFloors []int
	err := json.Unmarshal([]byte(team.ClearedFloors), &clearedFloors)

	return clearedFloors, err
}

func (team *TeamTable) UpdateClearedFloors(clearedFloors []int) error {
	// json 文字列にする
	jsonClearedFloors, err := json.Marshal(clearedFloors)

	// エラー処理
	if err != nil {
		return err
	}

	// json 文字列にする
	team.ClearedFloors = string(jsonClearedFloors)

	// DBに保存
	result := dbconn.Model(&team).Update("ClearedFloors", team.ClearedFloors)

	return result.Error
}

func GetTeam(gameid string,teamID string) (TeamTable, error) {
	var team TeamTable
	result := dbconn.Where(&TeamTable{
		TeamID: teamID,
	}).First(&team)

	return team, result.Error
}