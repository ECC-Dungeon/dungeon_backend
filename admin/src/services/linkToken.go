package services

import (
	"admin/models"
	"admin/utils"
	"net/http"
)

const (
	// 15分を有効期限とする (秒)
	LinkTokenExpired = 15 * 60
)

type GenTokenArgs struct {
	Teamid  string `json:"teamid"`
	Userid  string `json:"userid"`
	Tokenid string `json:"tokenid"`
	Expired int64  `json:"expired"`
}

func GenLinkToken(teamid string, userid string) (string, error) {
	// トークンのidを作成
	tokenid := utils.GenID()

	// トークンの有効期限を作成
	expired := utils.Now() + LinkTokenExpired

	// トークンを作成
	token, err := GenJwt(GenTokenArgs{
		Teamid:  "",
		Userid:  "",
		Tokenid: tokenid,
		Expired: expired,
	})

	// エラー処理
	if err != nil {
		return "", err
	}

	// リンク用のトークンを作成
	err = models.CreateLinkToken(teamid, tokenid, expired)

	// エラー処理
	if err != nil {
		return "", err
	}

	return token, nil
}

func InitLink(token string) (string, utils.HttpResult) {
	// トークンを検証
	args, err := ParseToken(token)

	// エラー処理
	if err != nil {
		// トークンの検証に失敗した場合は400を返す
		return "", utils.NewHttpResult(http.StatusBadRequest, "failed to verify token", err)
	}

	// リンク用のトークンを取得
	tokenData, err := models.GetLinkToken(args.Tokenid)

	// エラー処理
	if err != nil {
		// トークンが存在しない場合は403を返す
		return "", utils.NewHttpResult(http.StatusForbidden, "token not found", err)
	}

	// リンク用のトークンを削除
	err = models.DeleteLinkToken(tokenData.TeamID)

	// エラー処理
	if err != nil {
		// トークンの削除に失敗した場合は500を返す
		return "", utils.NewHttpResult(http.StatusInternalServerError, "failed to delete token", err)
	}

	//ゲーム用のトークンを作成
	gameToken, err := GenGameLink(tokenData.TeamID)

	// エラー処理
	if err != nil {
		// トークンの作成に失敗した場合は500を返す
		return "", utils.NewHttpResult(http.StatusInternalServerError, "failed to create game token", err)
	}
	
	return gameToken, utils.NewHttpResult(http.StatusOK, "success", nil)
}