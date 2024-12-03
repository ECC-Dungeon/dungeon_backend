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

// 使用する階を設定
func SetUseFloors(floors []int) error {
	// すべてのフロアを未使用にする
	allfloors,err := GetUseFloors()

	// エラー処理
	if err != nil {
		return err
	}

	// 未使用にする
	err = SetUnUseFloors(allfloors)

	// エラー処理
	if err != nil {
		return err
	}

	// フロアを保存
	for _, floor := range floors {
		result := dbconn.Model(&Floors{FloorNum: floor}).Update("is_used", true)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}

func SetUnUseFloors(floors []Floors) error {
	// フロアを保存
	for _, floor := range floors {
		result := dbconn.Model(&floor).Update("is_used", false)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
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
