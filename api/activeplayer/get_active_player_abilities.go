package activeplayer

import (
	"encoding/json"
	"github.com/hxcuber/lol-overlay/response/activeplayer/abilities"
	"io"
	"net/http"
)

func GetActivePlayerAbilities() (abilities.Abilities, error) {
	resp, err := http.Get("https://127.0.0.1:2999/liveclientdata/activeplayerabilities")
	if err != nil {
		return abilities.Abilities{}, err
	}
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return abilities.Abilities{}, err
	}
	var as abilities.Abilities
	err = json.Unmarshal(resBody, &as)
	if err != nil {
		return abilities.Abilities{}, err
	}
	return as, nil
}
