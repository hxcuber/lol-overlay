package controller

import (
	"github.com/hxcuber/lol-overlay/repositories/datadragon"
	"github.com/hxcuber/lol-overlay/repositories/liveclientdata"
)

type Controller interface {
	GetActivePlayerCSPM() (float64, error)
	GetActivePlayerGPM() (float64, error)
	GetCombatGold() (map[string]int, error)
}

type impl struct {
	dataDragonReg     datadragon.Registry
	liveClientDataReg liveclientdata.Registry
}

func New(dataDragonReg datadragon.Registry, liveClientDataReg liveclientdata.Registry) Controller {
	return impl{
		dataDragonReg:     dataDragonReg,
		liveClientDataReg: liveClientDataReg,
	}
}
