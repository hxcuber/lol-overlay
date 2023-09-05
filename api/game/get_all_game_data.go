package game

import (
	"encoding/json"
	"github.com/hxcuber/lol-overlay/response"
	"io"
	"net/http"
)

func GetAllGameData() (response.AllGameData, error) {
	resp, err := http.Get("https://127.0.0.1:2999/liveclientdata/allgamedata")
	if err != nil {
		return response.AllGameData{}, err
	}
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return response.AllGameData{}, err
	}
	var agd response.AllGameData
	err = json.Unmarshal(resBody, &agd)
	if err != nil {
		return response.AllGameData{}, err
	}
	return agd, err
}
