package utils

import "time"

func Now() int64 {
	// タイムゾーンを設定
	jst, err := time.LoadLocation("Asia/Tokyo")
    if err != nil {
        panic(err)
    }

	// 現在時刻を取得
    nowJST := time.Now().In(jst)
	return nowJST.Unix()
}