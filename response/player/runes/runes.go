package runes

import "github.com/hxcuber/lol-overlay/response/activePlayer/fullRunes/generalRune"

type Runes struct {
	Keystone          generalrune.GeneralRune `json:"keystone"`
	PrimaryRuneTree   generalrune.GeneralRune `json:"primaryRuneTree"`
	SecondaryRuneTree generalrune.GeneralRune `json:"secondaryRuneTree"`
}
