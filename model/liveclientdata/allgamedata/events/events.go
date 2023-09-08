package events

import (
	"github.com/hxcuber/lol-overlay/model/liveclientdata/allgamedata/events/event"
)

type Events struct {
	Events []event.Event `json:"Events"`
}
