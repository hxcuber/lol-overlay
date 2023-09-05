package abilities

import "github.com/hxcuber/lol-overlay/response/activeplayer/abilities/ability"

type Abilities struct {
	E       ability.Ability `json:"E"`
	Passive ability.Ability `json:"Passive"`
	Q       ability.Ability `json:"Q"`
	R       ability.Ability `json:"R"`
	W       ability.Ability `json:"W"`
}
