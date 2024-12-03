package models

import "admin/utils"

type SettingKey string

const (
	IsInit      SettingKey = "init"
	IsGameStart SettingKey = "game_start"
)

type Setting struct {
	Key   SettingKey `gorm:"primaryKey"`
	Value string
}

type Floors struct {
	FloorNum int  `gorm:"primaryKey"` // フロア番号
	IsUsed   bool // 使用するかの設定
}

func GetSetting(key SettingKey) (string, error) {
	var setting Setting
	result := dbconn.Where(&Setting{Key: key}).First(&setting)
	return setting.Value, result.Error
}

func SetSetting(key SettingKey, value string) error {
	setting := Setting{Key: key, Value: value}
	result := dbconn.Save(&setting)
	return result.Error
}

// 使用する階を取得
func GetUseFloors() ([]Floors, error) {
	// フロアを取得
	var floors []Floors

	// 階を取得
	result := dbconn.Find(&floors)

	return floors, result.Error
}

func InitSetting() error {
	// 初期化
	err := SetSetting(IsInit, "true")

	// エラー処理
	if err != nil {
		// 初期化に失敗した時
		utils.Println("初期化失敗 : " + err.Error())
	}

	// ゲーム開始を false にする
	err = SetSetting(IsGameStart, "false")

	// エラー処理
	if err != nil {
		// 初期化に失敗した時
		utils.Println("初期化失敗 : " + err.Error())
	}

	// 使用する階を入れる
	// 7階ループ
	for i := 1; i <= 7; i++ {
		// データに保存
		result := dbconn.Save(&Floors{FloorNum: i, IsUsed: false})

		// エラー処理
		if result.Error != nil {
			// 初期化に失敗した時
			utils.Println("初期化失敗 : " + result.Error.Error())
			continue
		}
	}

	return nil
}
