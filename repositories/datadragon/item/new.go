package item

type Repository interface {
	GetPrice(id int) (int, error)
}

type impl struct {
	address string
}

func New(address string) Repository {
	return impl{address: address}
}
