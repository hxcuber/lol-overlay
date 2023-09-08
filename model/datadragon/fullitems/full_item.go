package fullitems

type FullItem struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Colloq      string   `json:"colloq"`
	Plaintext   string   `json:"plaintext"`
	Into        []string `json:"into"`
	Image       struct {
		Full   string `json:"full"`
		Sprite string `json:"sprite"`
		Group  string `json:"group"`
		X      int    `json:"x"`
		Y      int    `json:"y"`
		W      int    `json:"w"`
		H      int    `json:"h"`
	} `json:"image"`
	Gold struct {
		Base        int  `json:"base"`
		Purchasable bool `json:"purchasable"`
		Total       int  `json:"total"`
		Sell        int  `json:"sell"`
	} `json:"gold"`
	Tags []string `json:"tags"`
	Maps struct {
		Field1 bool `json:"11"`
		Field2 bool `json:"12"`
		Field3 bool `json:"21"`
		Field4 bool `json:"22"`
		Field5 bool `json:"30"`
	} `json:"maps"`
	Stats map[string]float64 `json:"stats"`
}
