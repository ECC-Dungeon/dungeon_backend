package main

import (
	// "game/models"
	grpcserver "game/grpcServer"
	"log"
	"os"
)

func main() {
	// ENV を読み込み
	loadEnv()

	// 諸々初期化
	Init()

	// GRPC サーバー起動
	go grpcserver.RunGRPC()

	// サーバー起動
	RunServer()

	// デバッグ起動
	// models.Debug()
}

func RunServer() {
	log.Println("サーバーを起動しています")

	// サーバー初期化
	server := InitServer()

	// サーバー起動
	server.Logger.Fatal(server.Start(os.Getenv("BindAddr")))

}
