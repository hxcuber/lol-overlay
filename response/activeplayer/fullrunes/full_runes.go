package fullrunes

import (
	"github.com/hxcuber/lol-overlay/response/activePlayer/fullRunes/generalRune"
	"github.com/hxcuber/lol-overlay/response/activePlayer/fullRunes/statRune"
)

type FullRunes struct {
	GeneralRunes      []generalrune.GeneralRune `json:"generalRunes"`
	Keystone          generalrune.GeneralRune   `json:"keystone"`
	PrimaryRuneTree   generalrune.GeneralRune   `json:"primaryRuneTree"`
	SecondaryRuneTree generalrune.GeneralRune   `json:"secondaryRuneTree"`
	StatRunes         []statrune.StatRune       `json:"statRunes"`
}
