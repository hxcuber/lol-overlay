package summonerspells

import "github.com/hxcuber/lol-overlay/response/player/summonerspells/summonerspell"

type SummonerSpells struct {
	SummonerSpellOne summonerspell.SummonerSpell `json:"summonerSpellOne"`
	SummonerSpellTwo summonerspell.SummonerSpell `json:"summonerSpellTwo"`
}
