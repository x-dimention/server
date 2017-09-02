package model

import (
	"time"
)

type Model struct {
	key       int64     `xorm:"pk autoincr"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
	State     string
}

type ModeRec struct {
	key        int64
	ActionTime time.Time
	ActionType int64
}

// type stringSlice s
