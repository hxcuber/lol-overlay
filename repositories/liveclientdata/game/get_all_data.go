package game

import (
	"encoding/json"
	"fmt"
	"github.com/hxcuber/lol-overlay/model/liveclientdata/allgamedata"
	"github.com/pkg/errors"
	"io"
	"net/http"
)

func (i impl) GetAllData() (allgamedata.AllGameData, error) {
	resp, err := http.Get(fmt.Sprintf("%s/allgamedata", i.address))
	if err != nil {
		return allgamedata.AllGameData{}, err
	}
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return allgamedata.AllGameData{}, err
	}
	var agd allgamedata.AllGameData
	err = json.Unmarshal(resBody, &agd)
	if err != nil {
		return allgamedata.AllGameData{}, err
	}
	return agd, errors.WithStack(err)
}
