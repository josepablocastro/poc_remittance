package ports

import "github.com/josepablocastro/poc_remittance/internal/application/core/domain"

type APIPort interface {
	ReceivePayment(payment domain.Payment) (domain.Payment, error)
}
