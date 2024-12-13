package grpcserver

import (
	"context"
	"game/gamerpc"
	"game/models"
	"log"
	"net"

	"google.golang.org/grpc"
)

// 階を渡すと分けてくれる関数
func rotateAndReturn(slice []int) [][]int {
	n := len(slice)
	result := make([][]int, n)

	for i := 0; i < n; i++ {
		rotated := make([]int, n)
		for j := 0; j < n; j++ {
			rotated[j] = slice[(i+j)%n]
		}
		result[i] = rotated
	}

	return result
}

func RunGRPC() {
	log.Print("main start")

	// 9000番ポートでクライアントからのリクエストを受け付けるようにする
	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	// Sample構造体のアドレスを渡すことで、クライアントからGetDataリクエストされると
	// GetDataメソッドが呼ばれるようになる
	gamerpc.RegisterGameServiceServer(grpcServer, &GameRPC{})

	// 以下でリッスンし続ける
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

	log.Print("main end")
}

type GameRPC struct {

}

// Start implements gamerpc.GameServiceServer.
func (game *GameRPC) Start(ctx context.Context, args *gamerpc.StartArgs) (*gamerpc.StartResult, error) {
	log.Println("start")
	log.Println(args.Teams)
	log.Println(args.Floors)
	log.Println(args)

	// チームのデータを削除
	err := models.DeleteALLTeam(args.Gameid)

	// エラー処理
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// フロアのリスト
	floors := []int{}

	for _, floor := range args.Floors {
		// int型に変換
		floors = append(floors, int(floor.Num))
	}

	// チームのリスト
	reloatedList := rotateAndReturn(floors)

	// チームを登録
	for index, team := range args.Teams {
		// チームを登録
		err := models.RegisterTeam(args.Gameid, team.Id, reloatedList[index % len(reloatedList)])

		// エラー処理
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}

	return &gamerpc.StartResult{}, nil
}
