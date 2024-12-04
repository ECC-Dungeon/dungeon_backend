package services

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)


type GenTokenArgs struct {
	Teamid  string `json:"teamid"`
	Userid  string `json:"userid"`
	Tokenid string `json:"tokenid"`
	Expired int64  `json:"expired"`
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