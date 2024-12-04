package models

import "admin/utils"

type LinkToken struct {
	TeamID     string `gorm:"primaryKey"` //チームID
	TokenID    string //トークンID
	ExpiryDate int64  //有効期限 (unix time)
	CreatedAt  int64  //作成日
}

//有効期限を超えたリンクを削除する (削除件数を返す)
func DeleteExpiredLinkToken() (int64,error) {
	// 有効期限を超えたリンクを削除する
	result := dbconn.Where("expiry_date < ?", utils.Now()).Delete(&LinkToken{})

	return result.RowsAffected, result.Error	
}

func CreateLinkToken(teamid string, tokenid string, expiryDate int64) error {
	// チームを取得する
	team, err := GetTeam(teamid)

	// エラー処理
	if err != nil {
		return err
	}

	// リンクを作成する
	result := dbconn.Save(&LinkToken{
		TeamID:     team.TeamID,
		TokenID:    tokenid,
		ExpiryDate: expiryDate,
		CreatedAt:  utils.Now(),
	})

	return result.Error
}

func DeleteLinkToken(teamid string) error {
	// リンクを削除する
	result := dbconn.Delete(&LinkToken{
		TeamID: teamid,
	})

	return result.Error
}

func GetLinkToken(tokenid string) (LinkToken, error) {
	// チーム取得
	getLinkToken := LinkToken{}

	// データベースから取得
	result := dbconn.Where(&LinkToken{
		TokenID: tokenid,
	}).First(&getLinkToken)

	// エラー処理
	if result.Error != nil {
		return LinkToken{}, result.Error
	}

	return getLinkToken, nil
}