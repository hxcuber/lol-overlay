package controller

func (i impl) GetActivePlayerGPM() (float64, error) {
	ap, err := i.liveClientDataReg.ActivePlayer().Get()
	if err != nil {
		return 0, err
	}
	is, err := i.liveClientDataReg.Player().GetItems(ap.SummonerName)
	if err != nil {
		return 0, err
	}
	gd, err := i.liveClientDataReg.Game().GetStats()
	if err != nil {
		return 0, err
	}

	total := 0
	for _, item := range is {
		price, err := i.dataDragonReg.Item().GetPrice(item.ItemID)
		if err != nil {
			return 0, err
		}
		total += price
	}
	return (float64(total) + ap.CurrentGold) / (gd.GameTime / 60.0), nil
}
