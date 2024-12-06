package models

// import (
// 	"game/middlewares"
// 	"log"
// 	"slices"
// )

// func Debug() {
// 	GameID := "test"
// 	TeamID1 := "test"
// 	TeamID2 := "test2"
// 	TeamID3 := "test3"

// 	// 使用するフロア情報
// 	UseFloors := []middlewares.Floor{
// 		{GameID: GameID, FloorNum: 1, Name: "test", Enabled: false},
// 		{GameID: GameID, FloorNum: 2, Name: "test", Enabled: true}, //使用
// 		{GameID: GameID, FloorNum: 3, Name: "test", Enabled: false},
// 		{GameID: GameID, FloorNum: 4, Name: "test", Enabled: true}, //使用
// 		{GameID: GameID, FloorNum: 5, Name: "test", Enabled: true}, //使用
// 		{GameID: GameID, FloorNum: 6, Name: "test", Enabled: false},
// 		{GameID: GameID, FloorNum: 7, Name: "test", Enabled: false},
// 	}

// 	// ログを追加する
// 	CreateLog(GameID, TeamID1, 2)
// 	CreateLog(GameID, TeamID1, 4)
// 	CreateLog(GameID, TeamID1, 5)

// 	CreateLog(GameID, TeamID2, 2)
// 	CreateLog(GameID, TeamID2, 4)

// 	CreateLog(GameID, TeamID3, 2)
// 	CreateLog(GameID, TeamID3, 5)
// 	CreateLog(GameID, TeamID3, 2)

// 	// クリア済みのログを取得する
// 	CleardLogs, err := GetCleardFloors(GameID, TeamID2)

// 	// エラー処理
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	log.Println("クリアしたかい", CleardLogs)

// 	log.Println("ALL CLEAR", CheckAllClear(UseFloors, CleardLogs))

// 	// CreateLog(GameID, TeamID2, 5)

// 	_ = UseFloors
// }

// func ConvertLogs(logs []LogModel) ([]int, error) {
// 	CleardFloor := []int{}

// 	// ログをfor文で回す
// 	for _, val := range logs {
// 		CleardFloor = append(CleardFloor, val.FloorNum)
// 	}

// 	return CleardFloor, nil
// }

// func CheckAllClear(floors []middlewares.Floor, CleardFloor []int) bool {
// 	for _, val := range floors {
// 		// フロアが使用しない場合
// 		if !val.Enabled {
// 			continue
// 		}

// 		// クリア済みフロアに含まれているか
// 		if slices.Contains(CleardFloor, val.FloorNum) {
// 			// 含まれている場合
// 			continue
// 		}

// 		// それ以外の場合
// 		return false
// 	}

// 	return true
// }