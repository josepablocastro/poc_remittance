package poc_remittance

import (
	"github.com/josepablocastro/poc_remittance/internal/adapters/db"
	"github.com/josepablocastro/poc_remittance/internal/application/core/domain"
	"github.com/josepablocastro/poc_remittance/internal/ports"
)

type Application struct {
	db ports.DBPort
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{
		db: db,
	}
}

func NewDBAdapter(dataSourceUrl string) (*db.Adapter, error) {
	return db.NewAdapter(dataSourceUrl)
}

// func (a Application) ReceivePayment(payment domain.Payment) (domain.Payment, error) {
func (a Application) ReceivePayment(number string, sender string, beneficiary string, amount float32, currency string) (domain.Payment, error) {
	payment := domain.ReceivePayment(number, sender, beneficiary, amount, currency)

	err := a.db.Save(&payment)
	if err != nil {
		return domain.Payment{}, err
	}
	return payment, nil
}

func (a Application) AcceptPayment(number string, reject bool) (domain.Payment, error) {
	if reject {
		return a.db.RejectPayment(number)
	} else {
		return a.db.AcceptPayment(number)
	}
}
