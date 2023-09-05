package main

import (
	"crypto/tls"
	"fmt"
	"github.com/hxcuber/lol-overlay/api/activeplayer"
	"github.com/hxcuber/lol-overlay/api/game"
	"github.com/hxcuber/lol-overlay/api/player"
	"net/http"
	"time"
)

const httpVer = "http"

func main() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	//goland:noinspection GoInfiniteFor
	for {
		ap, err := activeplayer.GetActivePlayer()
		if err != nil {
			continue
		}

		s, err := player.GetPlayerScore(ap.SummonerName)
		if err != nil {
			continue
		}

		cg, err := game.GetGameStats()
		if err != nil {
			continue
		}

		fmt.Printf("CS/min: %.1f, G/min: %.0f\n", float64(s.CreepScore)/(cg.GameTime/60.0), ap.CurrentGold/(cg.GameTime/60.0))
		time.Sleep(time.Second * 5)
	}
}
