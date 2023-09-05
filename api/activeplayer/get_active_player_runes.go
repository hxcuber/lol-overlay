package activeplayer

import (
	"encoding/json"
	"github.com/hxcuber/lol-overlay/response/activeplayer/fullrunes"
	"io"
	"net/http"
)

func GetActivePlayerRunes() (fullrunes.FullRunes, error) {
	resp, err := http.Get("https://127.0.0.1:2999/liveclientdata/activeplayerrunes")
	if err != nil {
		return fullrunes.FullRunes{}, err
	}
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fullrunes.FullRunes{}, err
	}
	var fr fullrunes.FullRunes
	err = json.Unmarshal(resBody, &fr)
	if err != nil {
		return fullrunes.FullRunes{}, err
	}
	return fr, err
}
