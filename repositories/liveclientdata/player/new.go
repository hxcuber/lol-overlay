package player

import (
	"github.com/hxcuber/lol-overlay/model/liveclientdata/allgamedata/player"
	"github.com/hxcuber/lol-overlay/model/liveclientdata/allgamedata/player/item"
	"github.com/hxcuber/lol-overlay/model/liveclientdata/allgamedata/player/runes"
	"github.com/hxcuber/lol-overlay/model/liveclientdata/allgamedata/player/score"
	"github.com/hxcuber/lol-overlay/model/liveclientdata/allgamedata/player/summonerspells"
)

type Repository interface {
	GetList() ([]player.Player, error)
	GetItems(summonerName string) ([]item.Item, error)
	GetMainRunes(summonerName string) (runes.Runes, error)
	GetScore(summonerName string) (score.Score, error)
	GetSummonerSpells(summonerName string) (summonerspells.SummonerSpells, error)
}

type impl struct {
	address string
}

func New(address string) Repository {
	return impl{address: address}
}
