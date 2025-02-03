package device

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}
