package datadragon

import "github.com/hxcuber/lol-overlay/repositories/datadragon/item"

type Registry interface {
	Item() item.Repository
}

type impl struct {
	itemRepo item.Repository
}

func (i impl) Item() item.Repository {
	return i.itemRepo
}

func New(address string) Registry {
	return impl{
		itemRepo: item.New(address),
	}
}
