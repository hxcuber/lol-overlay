package player

import (
	"encoding/json"
	"fmt"
	"github.com/hxcuber/lol-overlay/response/player/score"
	"io"
	"net/http"
)

func GetPlayerScore(summonerName string) (score.Score, error) {
	resp, err := http.Get(fmt.Sprintf("https://127.0.0.1:2999/liveclientdata/playerscores?summonerName=%s", summonerName))
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
