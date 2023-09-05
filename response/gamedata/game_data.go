package gamedata

type GameData struct {
	GameMode   string  `json:"gameMode"`
	GameTime   float64 `json:"gameTime"`
	MapName    string  `json:"mapName"`
	MapNumber  int     `json:"mapNumber"`
	MapTerrain string  `json:"mapTerrain"`
}
