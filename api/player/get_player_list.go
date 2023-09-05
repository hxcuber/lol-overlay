package player

import (
	"encoding/json"
	"github.com/hxcuber/lol-overlay/response/player"
	"io"
	"net/http"
)

func GetPlayerList() ([]player.Player, error) {
	resp, err := http.Get("https://127.0.0.1:2999/liveclientdata/playerlist")
	if err != nil {
		return nil, err
	}
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var ps []player.Player
	err = json.Unmarshal(resBody, &ps)
	if err != nil {
		return nil, err
	}
	return ps, nil
}
