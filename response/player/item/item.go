package item

type Item struct {
	CanUse         bool   `json:"canUse"`
	Consumable     bool   `json:"consumable"`
	Count          int    `json:"count"`
	DisplayName    string `json:"displayName"`
	ItemID         int    `json:"itemID"`
	Price          int    `json:"price"`
	RawDescription string `json:"rawDescription"`
	RawDisplayName string `json:"rawDisplayName"`
	Slot           int    `json:"slot"`
}
