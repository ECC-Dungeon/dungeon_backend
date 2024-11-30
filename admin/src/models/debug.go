package models

import (
	"admin/utils"
	"log"
	"os"

	"time"
)

func Debug() {
	// データベースを削除する
	os.Remove("./admin.db")

	// 初期化
	Init()

	utils.Println("DBPATH : " + os.Getenv("DBPATH"))

	// チームを作成するユーザー
	doUser := "a7c3897d-1667-49c1-92a6-90a985e540af"
	TokenID := "471b9794-710b-4764-b9a0-ba3af0bea0e0"
	Gameid := "c1734747-e0ad-470e-9e64-e07122ca04f5"
	expired := time.Now().Add(time.Duration(time.Minute * 5))

	// 区切り線表示
	utils.ShowLine()
	teamid, err := debugCreateTeam(doUser, Gameid)

	// エラー処理
	if err != nil {
		utils.Println("チーム作成テスト失敗")
		return
	}

	// 区切り線を表示
	utils.ShowLine()
	// チームを取得する機能をデバッグする
	err = debugGetTeam(teamid)

	//エラー処理
	if err != nil {
		utils.Println("チーム取得エラー : " + err.Error())
		return
	}

	utils.Println("チーム取得成功")
	// 区切り線を表示
	utils.ShowLine()
	utils.Println("リンク作成をテスト")

	// チームリンクをテスト
	err = debugLinkTeam(teamid, TokenID, expired)

	// エラー処理
	if err != nil {
		log.Println("チーム作成テスト失敗")
		return
	}

	utils.Println("存在するチームのリンク作成")

	utils.Println("リンク作成テスト完了")

	utils.ShowLine()
}

func debugCreateTeam(doUser string, gameid string) (string, error) {
	utils.Println("チーム作成をテスト")
	// チームを作成する
	teamid, err := CreateTeam("wao", doUser, gameid)

	// エラー処理
	if err != nil {
		utils.Println("チーム作成エラー : " + err.Error())
		return "", err
	}

	utils.Println("作成したチーム : " + teamid)
	utils.Println("作成成功")

	return teamid, nil
}

func debugLinkTeam(teamid string, tokenid string, expired time.Time) error {
	// // 関連付けを作成する
	err := CreateGameLink(teamid, tokenid, expired.Unix())

	// エラー処理
	if err != nil {
		utils.Println("リンク失敗 : " + err.Error())
		return err
	}

	utils.Println("リンク作成成功")

	// 存在しないチームのリンク作成
	utils.Println("存在しないチームのリンク作成")
	// // 関連付けを作成する
	err = CreateGameLink(teamid+"b", tokenid, expired.Unix())

	// エラー処理
	if err != nil {
		utils.Println("リンク失敗 (成功)")
		return nil
	}

	return err
}

func debugGetTeam(teamid string) error {
	utils.Println("存在するチームを取得")

	// チームを取得する
	team, err := GetTeam(teamid)

	// エラー処理
	if err != nil {
		utils.Println("チーム取得失敗 : " + err.Error())
		return err
	}

	utils.Println("チーム取得成功 チームID: " + team.TeamID)

	_ = team

	// 存在しないチームを取得してみる
	utils.Println("存在しないチームを取得")

	// チームを取得する
	team2, err := GetTeam(teamid + "b")

	// エラー処理
	if err != nil {
		utils.Println("エラーを確認 (成功) エラー : " + err.Error())
		return nil
	}

	utils.Println("存在しないチーム取得 (失敗) 取得したID : " + team2.TeamID)

	return err
}
