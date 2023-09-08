package events

import (
	"github.com/hxcuber/lol-overlay/model/liveclientdata/allgamedata/events/event"
)

type Repository interface {
	Get() ([]event.Event, error)
}

type impl struct {
	address string
}

func New(address string) Repository {
	return impl{address: address}
}
