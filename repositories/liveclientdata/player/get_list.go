package player

import (
	"encoding/json"
	"fmt"
	"github.com/hxcuber/lol-overlay/model/liveclientdata/allgamedata/player"
	"github.com/pkg/errors"
	"io"
	"net/http"
)

func (i impl) GetList() ([]player.Player, error) {
	resp, err := http.Get(fmt.Sprintf("%s/playerlist", i.address))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var ps []player.Player
	err = json.Unmarshal(resBody, &ps)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ps, nil
}
