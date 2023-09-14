package database

import (
	"fmt"
	"paymentapi/entity"
)

type PaymentDatabase struct {
}

var PaymentDetails = make([]entity.Payment, 0)

func (pdb PaymentDatabase) AddPayment(newPayment entity.Payment) entity.Payment {
	PaymentDetails = append(PaymentDetails, newPayment)
	return newPayment
}

func (pdb PaymentDatabase) GetPayment() []entity.Payment {
	return PaymentDetails
}

func (pdb PaymentDatabase) GetSinglePayment(id string) []entity.Payment {
	return PaymentDetails

}

func (pdb PaymentDatabase) DeletePayment(id string) error {

	fmt.Println("Slice of Paymment details", PaymentDetails)

	for i, payment := range PaymentDetails {
		if payment.PaymentId == id {
			PaymentDetails = append(PaymentDetails[:i], PaymentDetails[i+1:]...)
		}
	}
	return fmt.Errorf("Payment not found")
}
