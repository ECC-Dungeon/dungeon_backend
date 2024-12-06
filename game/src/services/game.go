package services

import (
	"game/middlewares"
	"game/models"
	"log"
	"slices"
)

type NextData struct {
	NextNum int		//次のフロア
	AllClear bool	//全ての階をクリアしたか
	CleardFloor []int	//クリアした階
}

func Next(team middlewares.Team, floors []middlewares.Floor,ClearFloor int) (NextData,error) {
	// 現在のクリア状況を取得
	CleardFloor,err := GetCleardFloors(team.GameID, team.TeamID)

	// エラー処理
	if err != nil {
		log.Println(err)
		return NextData{}, err
	}

	// 全ての階をクリアした場合
	if CheckAllClear(floors, CleardFloor) {
		return NextData{
			NextNum:     0,
			AllClear:    true,
			CleardFloor: CleardFloor,
		}, nil
	}

	// クリアしてない場合
	// 初期化時
	if ClearFloor == -1 {
		// 一番人が少ないフロアを返す
		low_floor := models.GetLowFloor(team.GameID, floors)

		// チャンレンジを作成する
		err := models.CreateChallenge(team.GameID, team.TeamID, low_floor)

		// エラー処理
		if err != nil {
			log.Println(err)
			return NextData{}, err
		}
		
		// 次のフロアを返す
		return NextData{
			NextNum:     low_floor,
			AllClear:    false,
			CleardFloor: CleardFloor,
		}, nil
	}

	// それ以外の時
	// 既存のチャレンジを消す
	err = models.DeleteChallenge(team.GameID, team.TeamID, ClearFloor)

	// エラー処理
	if err != nil {
		log.Println(err)
		return NextData{}, err
	}

	// クリアのログを返す
	err = models.CreateLog(team.GameID, team.TeamID, ClearFloor)

	// エラー処理
	if err != nil {
		log.Println(err)
		return NextData{}, err
	}

	// 現在のクリア状況を取得
	CleardFloor,err = GetCleardFloors(team.GameID, team.TeamID)

	// エラー処理
	if err != nil {
		log.Println(err)
		return NextData{}, err
	}

	// クリア済みフロアを返す
	checked_floors := []middlewares.Floor{}

	for _, val := range floors {
		// フロアが使用しない場合
		if !val.Enabled {
			continue
		}

		// クリア済みに含まれている場合
		if slices.Contains(CleardFloor, val.FloorNum) {
			val.Enabled = false
		}

		// 追加
		checked_floors = append(checked_floors, val)
	}

	// 次のフロア取得
	low_floor := models.GetLowFloor(team.GameID, checked_floors)

	// 次のフロアを返す
	return NextData{
		NextNum:     low_floor,
		AllClear:    CheckAllClear(floors, CleardFloor),
		CleardFloor: CleardFloor,
	}, nil
}


func ConvertLogs(logs []models.LogModel) ([]int,error) {
	CleardFloor := []int{}

	// ログをfor文で回す
	for _, val := range logs {
		CleardFloor = append(CleardFloor, val.FloorNum)
	}

	return CleardFloor, nil
}

// クリアしたかいを取得する
func GetCleardFloors(gameID string, teamID string) ([]int, error) {
	// ログを取得する
	logs,err := models.GetLogs(gameID, teamID)

	// エラー処理
	if err != nil {
		log.Println(err)
		return []int{}, err
	}

	// ログを変換
	CleardFloor,err := ConvertLogs(logs)

	// エラー処理
	if err != nil {
		log.Println(err)
		return []int{}, err
	}

	return CleardFloor, nil
}


func CheckAllClear(floors []middlewares.Floor, CleardFloor []int) bool {
	for _, val := range floors {
		// フロアが使用しない場合
		if !val.Enabled {
			continue
		}

		// クリア済みフロアに含まれているか
		if slices.Contains(CleardFloor, val.FloorNum) {
			// 含まれている場合
			continue
		}

		// それ以外の場合
		return false
	}

	return true
}
