package game

import (
	"encoding/json"
	"fmt"
	"github.com/hxcuber/lol-overlay/model/liveclientdata/allgamedata/gamedata"
	"io"
	"net/http"
)

func (i impl) GetStats() (gamedata.GameData, error) {
	resp, err := http.Get(fmt.Sprintf("%s/gamestats", i.address))
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
