package player

import (
	"encoding/json"
	"fmt"
	"github.com/hxcuber/lol-overlay/response/player/item"
	"io"
	"net/http"
)

func GetPlayerItems(summonerName string) ([]item.Item, error) {
	resp, err := http.Get(fmt.Sprintf("https://127.0.0.1:2999/liveclientdata/playeritems?summonerName=%s", summonerName))
	if err != nil {
		return nil, err
	}
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var is []item.Item
	err = json.Unmarshal(resBody, &is)
	if err != nil {
		return nil, err
	}
	return is, nil
}
