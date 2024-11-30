package services

import (
	"admin/models"
	"admin/utils"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

const (
	// 24時間を有効期限とする (秒)
	GameTokenExpired = 24 * 60 * 60

	// 5分を有効期限とする (秒)
	LinkTokenExpired = 5 * 60
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
		return "", utils.NewHttpResult(400, "failed to verify token", err)
	}

	// リンク用のトークンを取得
	tokenData, err := models.GetLinkToken(args.Tokenid)

	// エラー処理
	if err != nil {
		// トークンが存在しない場合は403を返す
		return "", utils.NewHttpResult(403, "token not found", err)
	}

	// リンク用のトークンを削除
	err = models.DeleteLinkToken(tokenData.TeamID)

	// エラー処理
	if err != nil {
		// トークンの削除に失敗した場合は500を返す
		return "", utils.NewHttpResult(500, "failed to delete token", err)
	}

	//ゲーム用のトークンを作成
	gameToken, err := GenGameLink(tokenData.TeamID)

	// エラー処理
	if err != nil {
		// トークンの作成に失敗した場合は500を返す
		return "", utils.NewHttpResult(500, "failed to create game token", err)
	}
	
	return gameToken, utils.NewHttpResult(200, "success", nil)
}

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

func ParseToken(tokenString string) (GenTokenArgs, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// トークンの検証
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRETKEY")), nil
	})

	//検証に成功したか
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// 検証に成功した時
		return GenTokenArgs{
			Teamid:  claims["teamid"].(string),
			Userid:  claims["userid"].(string),
			Tokenid: claims["tokenid"].(string),
			Expired: int64(claims["exp"].(float64)),
		}, nil
	} else {
		// 検証に失敗した時
		return GenTokenArgs{}, err
	}
}

func GenJwt(args GenTokenArgs) (string, error) {
	// クライムを作成
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"teamid":  args.Teamid,
		"userid":  args.Userid,
		"tokenid": args.Tokenid,
		"exp":     args.Expired,
	})

	//トークンに署名
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRETKEY")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
