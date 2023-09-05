package activeplayer

import (
	"encoding/json"
	"github.com/hxcuber/lol-overlay/response/activeplayer"
	"io"
	"net/http"
)

func GetActivePlayer() (activeplayer.ActivePlayer, error) {
	resp, err := http.Get("https://127.0.0.1:2999/liveclientdata/activeplayer")
	if err != nil {
		return activeplayer.ActivePlayer{}, err
	}
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return activeplayer.ActivePlayer{}, err
	}
	var ap activeplayer.ActivePlayer
	err = json.Unmarshal(resBody, &ap)
	if err != nil {
		return activeplayer.ActivePlayer{}, err
	}
	return ap, nil
}
