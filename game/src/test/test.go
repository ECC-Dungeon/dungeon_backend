package main

import (
	"game/gamerpc"
	"game/models"
	"log"
	"os"
)

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

func main() {
	teams := []gamerpc.Team{
		{
			Id:   "1",
			Name: "team1",
		},
		{
			Id:   "2",
			Name: "team2",
		},
		{
			Id:   "3",
			Name: "team3",
		},
		{
			Id:   "4",
			Name: "team4",
		},
		{
			Id:   "5",
			Name: "team5",
		},
		{
			Id:   "6",
			Name: "team6",
		},
		{
			Id:   "7",
			Name: "team7",
		},
		{
			Id:   "8",
			Name: "team8",
		},
		{
			Id:   "9",
			Name: "team9",
		},
		{
			Id:   "10",
			Name: "team10",
		},
		{
			Id:   "11",
			Name: "team11",
		},
		{
			Id:   "12",
			Name: "team12",
		},
		{
			Id:   "13",
			Name: "team13",
		},
	}

	// ファイルを削除する
	os.Remove("./test.db")

	// 環境変数設定
	os.Setenv("DBPATH","./test.db")

	models.Init()
	TeamdiVision(teams, []int{2, 4, 5, 6})
}

func TeamdiVision(teams []gamerpc.Team, floors []int) {
	// slice := []int{2, 4, 5, 6}
	slice := []int{2, 4, 5, 6}
	rotatedList := rotateAndReturn(slice)

	for index, team := range teams {
		models.RegisterTeam("test",team.Id, rotatedList[index % len(rotatedList)])	
	}

	for index, team := range teams {
		models.RegisterTeam("test2",team.Id, rotatedList[index % len(rotatedList)])	
	}

	// チームを取得する
	team, err := models.GetTeam("test", "1")

	// エラー処理
	if err != nil {
		panic(err)
	}

	team.UpdateChallenges([]int{4, 5, 6})

	// 使用する階を取得する
	challenges, err := team.GetChallenges()

	// エラー処理
	if err != nil {
		panic(err)
	}

	log.Println(challenges)

	team2,err := models.GetTeam("test2", "1")

	// エラー処理
	if err != nil {
		panic(err)
	}

	challenges2, err := team2.GetChallenges()

	// エラー処理
	if err != nil {
		panic(err)
	}

	log.Println(challenges2)
}