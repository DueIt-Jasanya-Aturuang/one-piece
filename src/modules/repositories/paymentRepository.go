package repositories

type PaymentRepository interface{}

type PaymentRepositoryImpl struct{}

func NewPaymentRepositoryImpl() PaymentRepository {
	return &PaymentRepositoryImpl{}
}
