package liveclientdata

import (
	"github.com/hxcuber/lol-overlay/repositories/liveclientdata/activeplayer"
	"github.com/hxcuber/lol-overlay/repositories/liveclientdata/events"
	"github.com/hxcuber/lol-overlay/repositories/liveclientdata/game"
	"github.com/hxcuber/lol-overlay/repositories/liveclientdata/player"
)

type Registry interface {
	ActivePlayer() activeplayer.Repository
	Events() events.Repository
	Game() game.Repository
	Player() player.Repository
}

type impl struct {
	activePlayerRepo activeplayer.Repository
	eventsRepo       events.Repository
	gameRepo         game.Repository
	playerRepo       player.Repository
}

func (i impl) ActivePlayer() activeplayer.Repository {
	return i.activePlayerRepo
}
func (i impl) Events() events.Repository {
	return i.eventsRepo
}
func (i impl) Game() game.Repository {
	return i.gameRepo
}
func (i impl) Player() player.Repository {
	return i.playerRepo
}

func New(address string) Registry {
	return impl{
		activePlayerRepo: activeplayer.New(address),
		eventsRepo:       events.New(address),
		gameRepo:         game.New(address),
		playerRepo:       player.New(address),
	}
}
