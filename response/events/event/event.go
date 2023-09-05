package event

type Event struct {
	EventID    int      `json:"EventID"`
	EventName  string   `json:"EventName"`
	EventTime  float64  `json:"EventTime"`
	Assisters  []string `json:"Assisters,omitempty"`
	KillerName string   `json:"KillerName,omitempty"`
	VictimName string   `json:"VictimName,omitempty"`
	Recipient  string   `json:"Recipient,omitempty"`
}
