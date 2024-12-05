package models

import "admin/utils"

type Status string

const (
	UnUsed      = Status("unused")
	Used        = Status("used")
	TeamStarted = Status("started")
)

type Team struct {
	TeamID    string `gorm:"primaryKey"`
	Name      string //チーム名
	GameID    string //ゲームID
	Status    Status //チームステータス
	NickName  string //チーム名 (学生が決める)
	Creator   string //作成者ID
	CreatedAt int64  `gorm:"autoCreateTime"` //作成時間
}

func CreateTeam(teamName string, gameID string, creatorID string) (Team, error) {
	// ID を作成する
	TeamID := utils.GenID()

	// チームを作成する
	team := Team{
		TeamID:    TeamID,
		Name:      teamName,
		GameID:    gameID,
		Status:    UnUsed,
		NickName:  teamName,
		Creator:   creatorID,
		CreatedAt: utils.Now(),
	}

	// チームを作成する
	result := dbconn.Save(&team)

	// エラー処理
	if result.Error != nil {
		return Team{}, result.Error
	}

	return team, nil
}

// チームを取得
func GetTeam(teamid string) (Team, error) {
	var team Team

	// チームを取得
	result := dbconn.Where(&Team{
		TeamID: teamid,
	}).First(&team)

	return team, result.Error
}

// チームを削除する処理
func (team *Team) Delete() error {
	// トークンを削除する
	err := team.UnregisterGameLink()

	// エラー処理
	if err != nil {
		return err
	}

	// チームを削除する
	result := dbconn.Delete(team)

	// エラー処理
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// チームのニックネームを更新
func (team *Team) UpdateNickName(name string) error {
	team.NickName = name

	// チームを更新する
	result := dbconn.Save(team)

	// エラー処理
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (team *Team) StartGame() error {
	team.Status = TeamStarted

	// チームを更新する
	result := dbconn.Save(&team)

	// エラー処理
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (team *Team) EndGame() error {
	team.Status = Used

	// チームを更新する
	result := dbconn.Save(&team)

	// エラー処理
	if result.Error != nil {
		return result.Error
	}

	return nil
}