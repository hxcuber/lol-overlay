package main

import (
	"crypto/tls"
	"github.com/hxcuber/lol-overlay/controller"
	"github.com/hxcuber/lol-overlay/repositories/datadragon"
	"github.com/hxcuber/lol-overlay/repositories/liveclientdata"
	"log"
	"net/http"
	"time"
)

func main() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	log.Println("Starting up")
	dataDragonReg := datadragon.New("https://ddragon.leagueoflegends.com/cdn/13.17.1/data/en_US")
	liveClientDataReg := liveclientdata.New("https://127.0.0.1:2999/liveclientdata")
	ctrl := controller.New(dataDragonReg, liveClientDataReg)
	//goland:noinspection GoInfiniteFor
	for {
		time.Sleep(time.Second * 5)
		cspm, err := ctrl.GetActivePlayerCSPM()
		if err != nil {
			log.Printf("%+v\n", err)
			continue
		}
		gpm, err := ctrl.GetActivePlayerGPM()
		if err != nil {
			log.Printf("%+v\n", err)
			continue
		}
		cg, err := ctrl.GetCombatGold()
		if err != nil {
			log.Printf("%+v\n", err)
			continue
		}
	}
}
