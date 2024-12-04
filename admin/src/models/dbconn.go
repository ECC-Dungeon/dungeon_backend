package models

import (
	"admin/utils"
	"log"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	dbconn *gorm.DB = nil
)


func Init() {
	// データベースを開く
	db, err := gorm.Open(sqlite.Open(os.Getenv("DBPATH")), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}
	
	// グローバル変数に格納
	dbconn = db

	// マイグレーション
	db.AutoMigrate(&Team{})
	db.AutoMigrate(&GameLink{})
	db.AutoMigrate(&Game{})
	db.AutoMigrate(&Floors{})

	// 有効期限が過ぎたリンクを削除する (ループさせる)
	go func() {
		// recover
		defer func() {
			// 例外処理
			if rec := recover(); rec != nil {
				utils.Println(rec)
			}
		}()
	
		for {
			// 有効期限を超えたゲームリンクを削除する
			count, err := DeleteExpiredGameLink()

			// エラー処理
			if err != nil {
				// 削除に失敗した時
				utils.Println("ゲームリンク削除失敗 : " + err.Error())
			}

			//削除したときにログを出す 
			if count > 0 {
				utils.Println("削除したゲームリンク数 : " + strconv.FormatInt(count, 10))
			}

			time.Sleep(time.Second * 5)
		}
	}()
}
