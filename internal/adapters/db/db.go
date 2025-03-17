package db

import (
	"fmt"

	"github.com/josepablocastro/poc_remittance/internal/application/core/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	Sender      string
	Beneficiary string
	Amount      float32
	Currency    string
	Number      string
	State       string
}

type Adapter struct {
	db *gorm.DB
}

func NewAdapter(dataSourceUrl string) (*Adapter, error) {
	db, openErr := gorm.Open(postgres.Open(dataSourceUrl), &gorm.Config{})

	if openErr != nil {
		return nil, fmt.Errorf("db connection error: %v", openErr)
	}

	migrationErr := db.AutoMigrate(&Payment{})

	if migrationErr != nil {
		return nil, fmt.Errorf("db migration error: %v", migrationErr)
	}

	return &Adapter{db: db}, nil
}

func (a Adapter) GetByNumber(number string) (domain.Payment, error) {
	var paymentEntity Payment
	res := a.db.Where("number = ?", number).First(&paymentEntity)

	payment := domain.Payment{
		ID:          int64(paymentEntity.ID),
		Sender:      paymentEntity.Sender,
		Beneficiary: paymentEntity.Beneficiary,
		Amount:      paymentEntity.Amount,
		Currency:    paymentEntity.Currency,
		Number:      paymentEntity.Number,
		State:       paymentEntity.State,
		CreatedAt:   paymentEntity.CreatedAt.Unix(),
	}

	return payment, res.Error
}

func (a Adapter) Save(payment *domain.Payment) error {
	paymentModel := Payment{
		Sender:      payment.Sender,
		Beneficiary: payment.Beneficiary,
		Amount:      payment.Amount,
		Currency:    payment.Currency,
		Number:      payment.Number,
		State:       payment.State,
	}
	res := a.db.Create(&paymentModel)
	if res.Error != nil {
		payment.ID = int64(paymentModel.ID)
	}
	return res.Error
}

func (a Adapter) UpdateState(payment *domain.Payment, newState string) error {
	res := a.db.Model(&Payment{}).Where("number = ?", payment.Number).Update("State", newState)
	if res.Error != nil {
		payment.State = newState
	}
	return res.Error
}

func (a Adapter) AcceptPayment(number string) (domain.Payment, error) {
	payment, err := a.GetByNumber(number)
	if err != nil {
		return domain.Payment{}, err
	}
	err = a.UpdateState(&payment, "accepted")
	return payment, err
}

func (a Adapter) RejectPayment(number string) (domain.Payment, error) {
	payment, err := a.GetByNumber(number)
	if err != nil {
		return domain.Payment{}, err
	}
	err = a.UpdateState(&payment, "rejected")
	return payment, err
}
