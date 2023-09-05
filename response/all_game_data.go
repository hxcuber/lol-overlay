package response

import (
	"github.com/hxcuber/lol-overlay/response/activeplayer"
	"github.com/hxcuber/lol-overlay/response/events"
	"github.com/hxcuber/lol-overlay/response/gamedata"
	"github.com/hxcuber/lol-overlay/response/player"
)

type AllGameData struct {
	ActivePlayer activeplayer.ActivePlayer `json:"activePlayer"`
	AllPlayers   []player.Player           `json:"allPlayers"`
	Events       events.Events             `json:"events"`
	GameData     gamedata.GameData         `json:"gameData"`
}
