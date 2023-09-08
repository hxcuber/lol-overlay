package events

import (
	"encoding/json"
	"fmt"
	"github.com/hxcuber/lol-overlay/model/liveclientdata/allgamedata/events"
	"github.com/hxcuber/lol-overlay/model/liveclientdata/allgamedata/events/event"
	"github.com/pkg/errors"
	"io"
	"net/http"
)

func (i impl) Get() ([]event.Event, error) {
	resp, err := http.Get(fmt.Sprintf("%s/eventdata", i.address))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var e events.Events
	err = json.Unmarshal(resBody, &e)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return e.Events, nil
}
