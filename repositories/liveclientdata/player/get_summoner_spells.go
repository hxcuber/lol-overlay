package player

import (
	"encoding/json"
	"fmt"
	"github.com/hxcuber/lol-overlay/model/liveclientdata/allgamedata/player/summonerspells"
	"io"
	"net/http"
)

func (i impl) GetSummonerSpells(summonerName string) (summonerspells.SummonerSpells, error) {
	resp, err := http.Get(fmt.Sprintf("%s/playersummonerspells?summonerName=%s", i.address, summonerName))
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
