package models

// Payload structure
type Payment struct {
	UserID        int `json:"user_id"`
	PaymentID     int `json:"payment_id"`
	DepositAmount int `json:"deposit_amount"`
}

func GetMockPayments() []Payment {
	return []Payment{
		{UserID: 1, PaymentID: 1, DepositAmount: 10},
		{UserID: 1, PaymentID: 2, DepositAmount: 20},
		{UserID: 2, PaymentID: 3, DepositAmount: 20},
	}
}

func GetMockDuplicatePayment() Payment {
	return Payment{UserID: 1, PaymentID: 1, DepositAmount: 10}

}
