package domain

import "time"

type Payment struct {
	ID          int64   `json:"id"`
	Sender      string  `json:"sender"`
	Beneficiary string  `json:"beneficiary"`
	Amount      float32 `json:"amount"`
	Currency    string  `json:"currency"`
	Number      string  `json:"number"`
	State       string  `json:"state"`
	CreatedAt   int64   `json:"created_at"`
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
