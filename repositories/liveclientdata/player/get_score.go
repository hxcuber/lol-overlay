package player

import (
	"encoding/json"
	"fmt"
	"github.com/hxcuber/lol-overlay/model/liveclientdata/allgamedata/player/score"
	"io"
	"net/http"
)

func (i impl) GetScore(summonerName string) (score.Score, error) {
	resp, err := http.Get(fmt.Sprintf("%s/playerscores?summonerName=%s", i.address, summonerName))
	if err != nil {
		return score.Score{}, err
	}
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return score.Score{}, err
	}
	var s score.Score
	err = json.Unmarshal(resBody, &s)
	if err != nil {
		return score.Score{}, err
	}
	return s, nil
}
