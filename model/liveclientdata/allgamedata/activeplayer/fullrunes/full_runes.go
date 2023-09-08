package fullrunes

import (
	"github.com/hxcuber/lol-overlay/model/liveclientdata/allgamedata/activeplayer/fullrunes/generalrune"
	"github.com/hxcuber/lol-overlay/model/liveclientdata/allgamedata/activeplayer/fullrunes/statrune"
)

type FullRunes struct {
	GeneralRunes      []generalrune.GeneralRune `json:"generalRunes"`
	Keystone          generalrune.GeneralRune   `json:"keystone"`
	PrimaryRuneTree   generalrune.GeneralRune   `json:"primaryRuneTree"`
	SecondaryRuneTree generalrune.GeneralRune   `json:"secondaryRuneTree"`
	StatRunes         []statrune.StatRune       `json:"statRunes"`
}
