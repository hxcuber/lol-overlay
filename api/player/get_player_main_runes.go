package player

import (
	"encoding/json"
	"fmt"
	"github.com/hxcuber/lol-overlay/response/player/runes"
	"io"
	"net/http"
)

func GetPlayerMainRunes(summonerName string) (runes.Runes, error) {
	resp, err := http.Get(fmt.Sprintf("https://127.0.0.1:2999/liveclientdata/playermainrunes?summonerName=%s", summonerName))
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
