package ports

import "github.com/josepablocastro/poc_remittance/internal/application/core/domain"

type DBPort interface {
	GetByNumber(number string) (domain.Payment, error)
	Save(payment *domain.Payment) error
	UpdateState(payment *domain.Payment, newState string) error
	AcceptPayment(number string) (domain.Payment, error)
	RejectPayment(number string) (domain.Payment, error)
}
