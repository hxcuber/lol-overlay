package player

import (
	"encoding/json"
	"fmt"
	"github.com/hxcuber/lol-overlay/response/player/summonerspells"
	"io"
	"net/http"
)

func GetPlayerSummonerSpells(summonerName string) (summonerspells.SummonerSpells, error) {
	resp, err := http.Get(fmt.Sprintf("https://127.0.0.1:2999/liveclientdata/playersummonerspells?summonerName=%s", summonerName))
	if err != nil {
		return summonerspells.SummonerSpells{}, err
	}
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return summonerspells.SummonerSpells{}, err
	}
	var s summonerspells.SummonerSpells
	err = json.Unmarshal(resBody, &s)
	if err != nil {
		return summonerspells.SummonerSpells{}, err
	}
	return s, nil
}
