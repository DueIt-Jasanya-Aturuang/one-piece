package repositories

type Repository interface {
	PaymentRepository
}

type RepositoryImpl struct {
	*PaymentRepositoryImpl
}

func NewRepositoryImpl() Repository {
	return &RepositoryImpl{
		PaymentRepositoryImpl: &PaymentRepositoryImpl{},
	}
}
