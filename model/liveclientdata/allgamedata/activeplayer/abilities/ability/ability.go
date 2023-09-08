package ability

type Ability struct {
	AbilityLevel   int    `json:"abilityLevel"`
	DisplayName    string `json:"displayName"`
	Id             string `json:"id"`
	RawDescription string `json:"rawDescription"`
	RawDisplayName string `json:"rawDisplayName"`
}
