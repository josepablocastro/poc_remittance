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

func (a Application) ReceivePayment(payment domain.Payment) (domain.Payment, error) {
	err := a.db.Save(&payment)
	if err != nil {
		return domain.Payment{}, err
	}
	return payment, nil
}
