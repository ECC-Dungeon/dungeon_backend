package services

import "admin/models"

func AdminGameStart() error {
	// ゲームを開始
	err := models.SetSetting(models.IsGameStart, "true")

	return err
}

func AdminGameStop() error {
	// ゲームを終了
	err := models.SetSetting(models.IsGameStart, "false")

	return err
}

func IsGameStarted() (bool, error) {
	// ゲームを開始しているかどうかを返す
	isStarted, err := models.GetSetting(models.IsGameStart)

	if err != nil {
		return false, err
	}

	return isStarted == "true", nil
}

func GetUseFloors() ([]models.Floors, error) {
	// 使用する階を返す
	return models.GetUseFloors()
}

func SetUseFloors(floorNums []int) error {
	// 使用する階を設定
	return models.SetUseFloors(floorNums)
}
