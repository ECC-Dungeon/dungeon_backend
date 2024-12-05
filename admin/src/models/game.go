package models

import "admin/utils"

type GameStatus string

const (
	Started GameStatus = "started"	// 開始
	Stopped GameStatus = "stopped"	// 停止
)

type Game struct {
	GameID    string   `gorm:"primaryKey"`
	Name      string   //ゲーム名
	CreatorID string   //作成者
	Teams     []Team   `gorm:"foreignKey:GameID;references:GameID;constraints:OnDelete:CASCADE"` //チーム
	Floors    []Floors `gorm:"foreignKey:GameID;references:GameID;constraints:OnDelete:CASCADE"` //使用するフロア
	Status    GameStatus	// ゲーム開始済みか
	CreatedAt int64 //作成日
}

func CreateGame(name string, creatorID string) (Game, error) {
	// ID を作成する
	GameId := utils.GenID()

	// ゲームを作成する
	game := Game{
		GameID:    GameId,
		Name:      name,
		CreatorID: creatorID,
		Teams:     []Team{},
		Floors:    []Floors{},
		Status:    Stopped,
		CreatedAt: utils.Now(),
	}

	// ゲームを作成する
	result := dbconn.Save(&game)

	// エラー処理
	if result.Error != nil {
		return Game{}, result.Error
	}

	// フロアを作成する
	for i := 1; i < 8; i++ {
		// フロアを作成する
		err := game.AddFloor(i, "フロア名", false)

		// エラー処理
		if err != nil {
			utils.Println("フロア作成失敗 : " + err.Error())
			continue
		}
	}

	return game, nil
}

// チームを追加
func (game *Game) AddTeam(teamID string, teamName string, creatorID string) error {
	// チームを作成
	team, err := CreateTeam(teamName, game.GameID, creatorID)

	// エラー処理
	if err != nil {
		return err
	}

	// チームを追加する
	return dbconn.Model(&game).Association("Teams").Append(&team)
}

// チームを削除
func (game *Game) RemoveTeam(teamID string) error {
	// チームを取得する
	team, err := GetTeam(teamID)

	// エラー処理
	if err != nil {
		return err
	}

	// チームの削除を行う
	if err := team.Delete(); err != nil {
		return err
	}

	// チームを削除する
	return nil
}

func (game *Game) Delete() error {
	// チーム一覧を取得
	teams, err := game.GetTeams()

	// エラー処理
	if err != nil {
		return err
	}

	// チームを削除する
	for _, team := range teams {
		// チームを削除
		if err := team.Delete(); err != nil {
			return err
		}
	}

	// フロア一覧を取得
	floors, err := game.GetFloors()

	// エラー処理
	if err != nil {
		return err
	}

	// フロアを削除する
	for _, floor := range floors {
		// フロアを削除
		if err := floor.Delete(); err != nil {
			return err
		}
	}

	// チームを削除する
	return dbconn.Delete(&game).Error
}

// ゲームを取得
func GetGame(gameid string) (Game, error) {
	var game Game

	// ゲームを取得する
	result := dbconn.Where(&Game{
		GameID: gameid,
	}).First(&game)

	return game, result.Error
}

// ゲームに所属するチームを取得
func (game *Game) GetTeams() ([]Team, error) {
	// チームリスト
	teams := []Team{}

	// チームを取得する
	err := dbconn.Model(&game).Association("Teams").Find(&teams)

	return teams, err
}

// ゲーム一覧を取得
func GetGames() ([]Game, error) {
	var games []Game

	// ゲームを取得する
	result := dbconn.Where(&Game{}).Find(&games)

	return games, result.Error
}

func (game *Game) Start() error {
	return dbconn.Model(&game).Update("status", Started).Error
}

func (game *Game) Stop() error {
	// チーム一覧を取得
	teams, err := game.GetTeams()

	// エラー処理
	if err != nil {
		return err
	}

	// チームを停止する
	for _, team := range teams {
		utils.Println("チーム停止 : " + team.TeamID)

		// チームを停止
		if err := team.EndGame(); err != nil {
			utils.Println("チーム停止失敗 : " + err.Error())
			continue
		}
	}

	return dbconn.Model(&game).Update("status", Stopped).Error
}