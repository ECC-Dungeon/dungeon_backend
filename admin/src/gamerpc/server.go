package gamerpc

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

var (
	Client GameServiceClient
)

func Init() {
	var conn *grpc.ClientConn

	// ゲームサーバーに接続
	conn, err := grpc.Dial("game:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}

	// クライアントを生成
	client := NewGameServiceClient(conn)

	// グローバル変数に格納
	Client = client
}

func StartGame(gameid string,teams []*Team,floors []*Floor) error {

	// ゲームを開始する
	Client.Start(context.Background(), &StartArgs{
		Gameid:        gameid,
		Floors:        floors,
		Teams:         teams,
	})

	return nil
}