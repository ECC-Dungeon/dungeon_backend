package utils

import (
	"github.com/google/uuid"
)

func GenID() string {
	// ID を生成
	uid,_ := uuid.NewRandom()

	return uid.String()
}