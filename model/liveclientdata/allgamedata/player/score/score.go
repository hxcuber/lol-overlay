package score

type Score struct {
	Assists    int     `json:"assists"`
	CreepScore int     `json:"creepScore"`
	Deaths     int     `json:"deaths"`
	Kills      int     `json:"kills"`
	WardScore  float64 `json:"wardScore"`
}
