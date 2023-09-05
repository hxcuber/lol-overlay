package game

import (
	"encoding/json"
	"github.com/hxcuber/lol-overlay/response/gamedata"
	"io"
	"net/http"
)

func GetGameStats() (gamedata.GameData, error) {
	resp, err := http.Get("https://127.0.0.1:2999/liveclientdata/gamestats")
	if err != nil {
		return gamedata.GameData{}, err
	}
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return gamedata.GameData{}, err
	}
	var gd gamedata.GameData
	err = json.Unmarshal(resBody, &gd)
	if err != nil {
		return gamedata.GameData{}, err
	}
	return gd, nil
}
