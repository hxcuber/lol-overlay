package allgamedata

import (
	"github.com/hxcuber/lol-overlay/model/liveclientdata/allgamedata/activeplayer"
	"github.com/hxcuber/lol-overlay/model/liveclientdata/allgamedata/events"
	"github.com/hxcuber/lol-overlay/model/liveclientdata/allgamedata/gamedata"
	"github.com/hxcuber/lol-overlay/model/liveclientdata/allgamedata/player"
)

type AllGameData struct {
	ActivePlayer activeplayer.ActivePlayer `json:"activePlayer"`
	AllPlayers   []player.Player           `json:"allPlayers"`
	Events       events.Events             `json:"events"`
	GameData     gamedata.GameData         `json:"gameData"`
}
