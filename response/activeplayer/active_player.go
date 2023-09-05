package activeplayer

import (
	"github.com/hxcuber/lol-overlay/response/activeplayer/abilities"
	"github.com/hxcuber/lol-overlay/response/activeplayer/championstats"
	"github.com/hxcuber/lol-overlay/response/activeplayer/fullrunes"
)

type ActivePlayer struct {
	Abilities          abilities.Abilities         `json:"abilities"`
	ChampionStats      championstats.ChampionStats `json:"championStats"`
	CurrentGold        float64                     `json:"currentGold"`
	FullRunes          fullrunes.FullRunes         `json:"fullRunes"`
	Level              int                         `json:"level"`
	SummonerName       string                      `json:"summonerName"`
	TeamRelativeColors bool                        `json:"teamRelativeColors"`
}
