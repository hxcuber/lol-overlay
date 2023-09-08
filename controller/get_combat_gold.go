package controller

func (i impl) GetCombatGold() (map[string]int, error) {
	ps, err := i.liveClientDataReg.Player().GetList()
	if err != nil {
		return nil, err
	}
	cg := map[string]int{
		"blue": 0,
		"red":  0,
	}
	for _, p := range ps {
		for _, item := range p.Items {
			price, err := i.dataDragonReg.Item().GetPrice(item.ItemID)
			if err != nil {
				return nil, err
			}
			switch p.Team {
			case "ORDER":
				{
					cg["blue"] += price
				}
			case "CHAOS":
				{
					cg["red"] += price
				}
			default:
				{
				}
			}
		}
	}
	return cg, nil
}
