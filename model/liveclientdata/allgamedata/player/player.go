package player

import (
	"github.com/hxcuber/lol-overlay/model/liveclientdata/allgamedata/player/item"
	"github.com/hxcuber/lol-overlay/model/liveclientdata/allgamedata/player/runes"
	"github.com/hxcuber/lol-overlay/model/liveclientdata/allgamedata/player/score"
	"github.com/hxcuber/lol-overlay/model/liveclientdata/allgamedata/player/summonerspells"
)

type Player struct {
	ChampionName    string                        `json:"championName"`
	IsBot           bool                          `json:"isBot"`
	IsDead          bool                          `json:"isDead"`
	Items           []item.Item                   `json:"items"`
	Level           int                           `json:"level"`
	Position        string                        `json:"position"`
	RawChampionName string                        `json:"rawChampionName"`
	RespawnTimer    float64                       `json:"respawnTimer"`
	Runes           runes.Runes                   `json:"runes"`
	Scores          score.Score                   `json:"scores"`
	SkinID          int                           `json:"skinID"`
	SummonerName    string                        `json:"summonerName"`
	SummonerSpells  summonerspells.SummonerSpells `json:"summonerSpells"`
	Team            string                        `json:"team"`
}
