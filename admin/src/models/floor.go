package models

type Floors struct {
	GameID   string `gorm:"primaryKey"`
	FloorNum int    `gorm:"primaryKey"` // フロア番号
	Name     string //フロアで行われているゲーム名
	Enabled  bool   // 有効かどうか
}

func (game *Game) AddFloor(floorNum int, name string, enabled bool) error {
	// 指定のフロアを追加する
	return dbconn.Model(&game).Association("Floors").Append(&Floors{GameID: game.GameID, FloorNum: floorNum, Name: name, Enabled: enabled})
}

func (game *Game) GetFloors() ([]Floors, error) {
	var floors []Floors

	// フロアを取得する
	result := dbconn.Where(&Floors{
		GameID: game.GameID,
	}).Find(&floors)

	return floors, result.Error
}

func (game *Game) ClearFloor() error {
	// db から削除
	result := dbconn.Where(&Floors{
		GameID: game.GameID,
	}).Delete(&Floors{})

	// エラー処理
	if result.Error != nil {
		return result.Error
	}

	// フロアを全て削除
	err := dbconn.Model(&game).Association("Floors").Clear()

	// エラー処理
	if err != nil {
		return err
	}

	return nil
}

// フロアを取得
func (game *Game) GetFloor(floorNum int) (Floors, error) {
	var floor Floors

	// フロアを取得する
	result := dbconn.Where(&Floors{
		GameID:   game.GameID,
		FloorNum: floorNum,
	}).First(&floor)

	return floor, result.Error
}

// フロアを削除
func (game *Game) RemoveFloor(floorNum int) error {
	// フロアを取得する
	floor, err := game.GetFloor(floorNum)

	// エラー処理
	if err != nil {
		return err
	}

	// 指定のフロアを削除する
	err = dbconn.Model(&game).Association("Floors").Delete(&floor)

	// エラー処理
	if err != nil {
		return err
	}

	// フロアを削除する
	err = floor.Delete()

	// エラー処理
	if err != nil {
		return err
	}

	return nil
}

func (floor *Floors) Delete() error {
	// フロアを削除する
	return dbconn.Delete(&floor).Error
}
