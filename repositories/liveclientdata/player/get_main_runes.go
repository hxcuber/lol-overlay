package player

import (
	"encoding/json"
	"fmt"
	"github.com/hxcuber/lol-overlay/model/liveclientdata/allgamedata/player/runes"
	"io"
	"net/http"
)

func (i impl) GetMainRunes(summonerName string) (runes.Runes, error) {
	resp, err := http.Get(fmt.Sprintf("%s/playermainrunes?summonerName=%s", i.address, summonerName))
	if err != nil {
		return runes.Runes{}, err
	}
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return runes.Runes{}, err
	}
	var r runes.Runes
	err = json.Unmarshal(resBody, &r)
	if err != nil {
		return runes.Runes{}, err
	}
	return r, nil
}
