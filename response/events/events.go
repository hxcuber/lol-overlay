package events

import "github.com/hxcuber/lol-overlay/response/events/event"

type Events struct {
	Events []event.Event `json:"Events"`
}
