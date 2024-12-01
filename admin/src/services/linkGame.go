package services

import (
	"admin/models"
	"admin/utils"
	"net/http"
)

const (
	// 24時間を有効期限とする (秒)
	GameTokenExpired = 24 * 60 * 60
)

func GenGameLink(teamid string) (string, error) {
	// トークンのidを作成
	tokenid := utils.GenID()

	// トークンの有効期限を作成
	expired := utils.Now() + GameTokenExpired

	// トークンを作成
	token, err := GenJwt(GenTokenArgs{
		Teamid:  teamid,
		Tokenid: tokenid,
		Expired: expired,
	})

	// エラー処理
	if err != nil {
		return "", err
	}

	// リンク用のトークンを作成
	err = models.CreateGameLink(teamid, tokenid, expired)

	// エラー処理
	if err != nil {
		return "", err
	}

	return token, nil
}

func UnLink(teamid string, tokenid string) utils.HttpResult {
	//　ゲーム用のトークンを削除
	err := models.DeleteGameLink(teamid)

	// エラー処理
	if err != nil {
		// トークンの削除に失敗した場合は500を返す
		return utils.NewHttpResult(http.StatusInternalServerError, "failed to delete token", err)
	}

	// リンク用のトークンを削除
	err = models.DeleteLinkToken(teamid)

	// エラー処理
	if err != nil {
		// トークンの削除に失敗した場合は500を返す
		return utils.NewHttpResult(http.StatusInternalServerError, "failed to delete token", err)
	}

	return utils.NewHttpResult(http.StatusOK, "success", nil)
}

func VerifyGameToken(token string) (models.Team, utils.HttpResult) {
	// トークンを検証
	args, err := ParseToken(token)

	// エラー処理
	if err != nil {
		// トークンの検証に失敗した場合は400を返す
		return models.Team{}, utils.NewHttpResult(http.StatusForbidden, "failed to verify token", err)
	}

	// チームを取得する
	team, err := models.GetTeam(args.Teamid)

	// エラー処理
	if err != nil {
		// トークンが存在しない場合は404を返す
		return models.Team{}, utils.NewHttpResult(http.StatusNotFound, "token not found", err)
	}

	return team, utils.NewHttpResult(http.StatusOK, "success", nil)
}