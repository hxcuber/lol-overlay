package activeplayer

import (
	"encoding/json"
	"fmt"
	"github.com/hxcuber/lol-overlay/model/liveclientdata/allgamedata/activeplayer/abilities"
	"github.com/pkg/errors"
	"io"
	"net/http"
)

func (i impl) GetAbilities() (abilities.Abilities, error) {
	resp, err := http.Get(fmt.Sprintf("%s/activeplayerabilities", i.address))
	if err != nil {
		return abilities.Abilities{}, errors.WithStack(err)
	}
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return abilities.Abilities{}, errors.WithStack(err)
	}
	var as abilities.Abilities
	err = json.Unmarshal(resBody, &as)
	if err != nil {
		return abilities.Abilities{}, errors.WithStack(err)
	}
	return as, nil
}
