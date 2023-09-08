package runes

import "github.com/hxcuber/lol-overlay/model/liveclientdata/allgamedata/activeplayer/fullrunes/generalrune"

type Runes struct {
	Keystone          generalrune.GeneralRune `json:"keystone"`
	PrimaryRuneTree   generalrune.GeneralRune `json:"primaryRuneTree"`
	SecondaryRuneTree generalrune.GeneralRune `json:"secondaryRuneTree"`
}
