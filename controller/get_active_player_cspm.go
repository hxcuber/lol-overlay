package controller

func (i impl) GetActivePlayerCSPM() (float64, error) {
	ap, err := i.liveClientDataReg.ActivePlayer().Get()
	if err != nil {
		return 0, err
	}
	s, err := i.liveClientDataReg.Player().GetScore(ap.SummonerName)
	if err != nil {
		return 0, err
	}
	gd, err := i.liveClientDataReg.Game().GetStats()
	if err != nil {
		return 0, err
	}
	return float64(s.CreepScore) / (gd.GameTime / 60.0), nil
}
