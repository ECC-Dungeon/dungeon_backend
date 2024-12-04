package models

import "admin/utils"

type GameLink struct {
	TeamID     string `gorm:"primaryKey"` //チームID
	TokenID    string //トークンID
	ExpiryDate int64  //有効期限 (unix time)
	CreatedAt  int64  //作成日
}

// 有効期限を超えたリンクを削除する (削除件数を返す)
func DeleteExpiredGameLink() (int64, error) {
	// 有効期限を超えたリンクを取得する
	gameLinks := []GameLink{}
	result := dbconn.Where("expiry_date < ?", utils.Now()).Find(&gameLinks)

	// エラー処理
	if result.Error != nil {
		return 0, result.Error
	}

	// 削除した件数
	deleteCount := 0

	// 有効期限を超えたリンクを削除する
	for _, gameLink := range gameLinks {
		// チームを取得する
		team, err := GetTeam(gameLink.TeamID)

		// エラー処理
		if err != nil {
			utils.Println("チーム取得失敗 : " + err.Error())
			continue
		}

		// リンクの登録を解除する
		err = team.UnregisterGameLink()

		// エラー処理
		if err != nil {
			utils.Println("リンク削除失敗 : " + err.Error())
			continue
		}

		// 削除した件数を加算する
		deleteCount += 1
	}

	return int64(deleteCount), nil
}

// ゲーム用のトークンの登録を解除
func (team *Team) UnregisterGameLink() error {
	// リンクを削除する
	result := dbconn.Where(&GameLink{
		TeamID: team.TeamID,
	}).Delete(&GameLink{})

	// エラー処理
	if result.Error != nil {
		return result.Error
	}

	// チームを未使用にする
	result = dbconn.Model(&team).Update("status", UnUsed)

	return result.Error
}

// ゲーム用のトークンを登録
func (team *Team) RegisterGameLink(tokenid string, expiryDate int64) error {
	// チームを取得する

	// リンクを作成する
	result := dbconn.Save(&GameLink{
		TeamID:     team.TeamID,
		TokenID:    tokenid,
		ExpiryDate: expiryDate,
		CreatedAt:  utils.Now(),
	})

	// エラー処理
	if result.Error != nil {
		return result.Error
	}

	// チームを使用中にする
	result = dbconn.Model(&team).Update("status", Used)

	return result.Error
}

func GetGameLink(tokenid string) (GameLink, error) {
	// チーム取得
	getGameLink := GameLink{}

	// データベースから取得
	result := dbconn.Where(&GameLink{
		TokenID: tokenid,
	}).First(&getGameLink)

	// エラー処理
	if result.Error != nil {
		return GameLink{}, result.Error
	}

	return getGameLink, nil
}
