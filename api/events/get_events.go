package events

import (
	"encoding/json"
	"github.com/hxcuber/lol-overlay/response/events"
	"io"
	"net/http"
)

func GetEvents() (events.Events, error) {
	resp, err := http.Get("https://127.0.0.1:2999/liveclientdata/eventdata")
	if err != nil {
		return events.Events{}, err
	}
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return events.Events{}, err
	}
	var e events.Events
	err = json.Unmarshal(resBody, &e)
	if err != nil {
		return events.Events{}, err
	}
	return e, nil
}
