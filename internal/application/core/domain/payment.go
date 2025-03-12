package domain

import "time"

type Payment struct {
	ID          int64
	Sender      string
	Beneficiary string
	Amount      float32
	Currency    string
	Number      string
	State       string
	CreatedAt   int64
}

func ReceivePayment(number string, sender string, beneficiary string, amount float32, currency string) Payment {
	return Payment{
		Number:      number,
		Sender:      sender,
		Beneficiary: beneficiary,
		Amount:      amount,
		Currency:    currency,
		State:       "received",
		CreatedAt:   time.Now().Unix(),
	}
}
