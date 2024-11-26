package main

import (
	"admin/models"
	"log"
	"os"
)

func main() {
	// env など読み込み
	Init()

	// デバッグ
	models.Debug()
}

func mainServer() {
	// 初期化
	Init()

	log.Println("サーバーを起動しています")

	// サーバー初期化
	server := InitServer()

	// サーバー起動
	server.Logger.Fatal(server.Start(os.Getenv("BindAddr")))
}