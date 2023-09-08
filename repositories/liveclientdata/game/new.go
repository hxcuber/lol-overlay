package game

import (
	"github.com/hxcuber/lol-overlay/model/liveclientdata/allgamedata"
	"github.com/hxcuber/lol-overlay/model/liveclientdata/allgamedata/gamedata"
)

type Repository interface {
	GetAllData() (allgamedata.AllGameData, error)
	GetStats() (gamedata.GameData, error)
}

type impl struct {
	address string
}

func New(address string) Repository {
	return impl{address: address}
}
