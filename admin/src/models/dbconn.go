package models

import (
	"admin/utils"
	"log"
	"os"
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

	// マイグレーション
	db.AutoMigrate(&Team{})
	db.AutoMigrate(&GameLink{})
	db.AutoMigrate(&LinkToken{})

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
				utils.Println("削除したゲームリンク数 : " + string(rune(count)))
			}

			// 有効期限を超えたリンクを削除する
			count, err = DeleteExpiredLinkToken()

			// エラー処理
			if err != nil {
				// 削除に失敗した時
				utils.Println("リンク削除失敗 : " + err.Error())
			}

			//削除したときにログを出す 
			if count > 0 {
				utils.Println("削除したリンク数 : " + string(rune(count)))
			}

			time.Sleep(time.Minute * 5)
		}
	}()

	// グローバル変数に格納
	dbconn = db
}
