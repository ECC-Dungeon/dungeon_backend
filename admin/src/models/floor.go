package models

type Floors struct {
	GameID   string `gorm:"primaryKey"`
	FloorNum int    `gorm:"primaryKey"` // フロア番号
}

func (game *Game) AddFloor(floorNum int) error {
	// 指定のフロアを追加する
	return dbconn.Model(&game).Association("Floors").Append(Floors{GameID: game.GameID, FloorNum: floorNum})
}

func (game *Game) RemoveFloor(floorNum int) error {
	// 指定のフロアを削除する
	return dbconn.Model(&game).Association("Floors").Delete(Floors{GameID: game.GameID, FloorNum: floorNum})
}

func (floor *Floors) Delete() error {
	// フロアを削除する
	return dbconn.Delete(&floor).Error
}