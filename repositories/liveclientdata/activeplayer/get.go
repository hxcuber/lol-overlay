package activeplayer

import (
	"encoding/json"
	"fmt"
	"github.com/hxcuber/lol-overlay/model/liveclientdata/allgamedata/activeplayer"
	"github.com/pkg/errors"
	"io"
	"net/http"
)

func (i impl) Get() (activeplayer.ActivePlayer, error) {
	resp, err := http.Get(fmt.Sprintf("%s/activeplayer", i.address))
	if err != nil {
		return activeplayer.ActivePlayer{}, errors.WithStack(err)
	}
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return activeplayer.ActivePlayer{}, errors.WithStack(err)
	}
	var ap activeplayer.ActivePlayer
	err = json.Unmarshal(resBody, &ap)
	if err != nil {
		return activeplayer.ActivePlayer{}, errors.WithStack(err)
	}
	return ap, nil
}
