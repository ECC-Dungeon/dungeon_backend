package main

import (
	"admin/gamerpc"
	"log"
	"os"
)

func main() {
	// env など読み込み
	Init()

	// デバッグ
	// models.Debug()

	// サーバー起動
	mainServer()
}

func mainServer() {
	// 初期化
	Init()

	// GRPC 初期化
	gamerpc.Init()

	log.Println("サーバーを起動しています")

	// サーバー初期化
	server := InitServer()

	// サーバー起動
	server.Logger.Fatal(server.Start(os.Getenv("BindAddr")))
}