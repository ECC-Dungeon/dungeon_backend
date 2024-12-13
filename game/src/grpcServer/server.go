package grpcserver

import (
	"context"
	"game/gamerpc"
	"google.golang.org/grpc"
	"log"
	"net"
)

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
	return &gamerpc.StartResult{}, nil
}
