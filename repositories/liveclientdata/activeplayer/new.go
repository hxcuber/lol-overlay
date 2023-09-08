package activeplayer

import (
	"github.com/hxcuber/lol-overlay/model/liveclientdata/allgamedata/activeplayer"
	"github.com/hxcuber/lol-overlay/model/liveclientdata/allgamedata/activeplayer/abilities"
	"github.com/hxcuber/lol-overlay/model/liveclientdata/allgamedata/activeplayer/fullrunes"
)

type Repository interface {
	Get() (activeplayer.ActivePlayer, error)
	GetAbilities() (abilities.Abilities, error)
	GetName() (string, error)
	GetRunes() (fullrunes.FullRunes, error)
}

type impl struct {
	address string
}

func New(address string) Repository {
	return impl{address: address}
}
