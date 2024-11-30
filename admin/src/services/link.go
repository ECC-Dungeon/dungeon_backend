package services

import (
	"admin/models"
	"admin/utils"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

const (
	// 24時間を有効期限とする (秒)
	Expired = 24 * 60 * 60
)

type GenTokenArgs struct {
	Teamid  string `json:"teamid"`
	Userid  string `json:"userid"`
	Tokenid string `json:"tokenid"`
	Expired int64  `json:"expired"`
}

func GenToken(teamid string, userid string) (string, error) {
	// トークンのidを作成
	tokenid := utils.GenID()

	// トークンの有効期限を作成
	expired := utils.Now() + Expired

	// トークンを作成
	token, err := GenJwt(GenTokenArgs{
		Teamid:  teamid,
		Userid:  userid,
		Tokenid: tokenid,
		Expired: expired,
	})

	// エラー処理
	if err != nil {
		return "", err
	}

	// リンクを作成
	err = models.CreateLink(teamid, tokenid, expired)

	// エラー処理
	if err != nil {
		return "", err
	}

	return token, nil
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
