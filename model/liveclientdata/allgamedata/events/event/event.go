package event

type Event struct {
	EventID      int      `json:"EventID"`
	EventName    string   `json:"EventName"`
	EventTime    float64  `json:"EventTime"`
	KillerName   string   `json:"KillerName,omitempty"`
	TurretKilled string   `json:"TurretKilled,omitempty"`
	Assisters    []string `json:"Assisters,omitempty"`
	InhibKilled  string   `json:"InhibKilled,omitempty"`
	DragonType   string   `json:"DragonType,omitempty"`
	Stolen       string   `json:"Stolen,omitempty"`
	VictimName   string   `json:"VictimName,omitempty"`
	KillStreak   int      `json:"KillStreak,omitempty"`
	Acer         string   `json:"Acer,omitempty"`
	AcingTeam    string   `json:"AcingTeam,omitempty"`
}
