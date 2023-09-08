package activeplayer

import (
	"encoding/json"
	"fmt"
	"github.com/hxcuber/lol-overlay/model/liveclientdata/allgamedata/activeplayer/fullrunes"
	"github.com/pkg/errors"
	"io"
	"net/http"
)

func (i impl) GetRunes() (fullrunes.FullRunes, error) {
	resp, err := http.Get(fmt.Sprintf("%s/activeplayerrunes", i.address))
	if err != nil {
		return fullrunes.FullRunes{}, errors.WithStack(err)
	}
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fullrunes.FullRunes{}, errors.WithStack(err)
	}
	var fr fullrunes.FullRunes
	err = json.Unmarshal(resBody, &fr)
	if err != nil {
		return fullrunes.FullRunes{}, errors.WithStack(err)
	}
	return fr, errors.WithStack(err)
}
