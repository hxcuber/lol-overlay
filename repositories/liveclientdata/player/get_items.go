package player

import (
	"encoding/json"
	"fmt"
	"github.com/hxcuber/lol-overlay/model/liveclientdata/allgamedata/player/item"
	"github.com/pkg/errors"
	"io"
	"net/http"
)

func (i impl) GetItems(summonerName string) ([]item.Item, error) {
	resp, err := http.Get(fmt.Sprintf("%s/playeritems?summonerName=%s", i.address, summonerName))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var is []item.Item
	err = json.Unmarshal(resBody, &is)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return is, nil
}
