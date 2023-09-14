package service

import (
	"fmt"
	"paymentapi/entity"
	database "paymentapi/storage"
	"time"
)

type PaymentService struct {
	Database database.PaymentDatabase
}

func (ps PaymentService) AddPayment(newPayment entity.Payment) entity.Payment {
	newPayment.CreatedAt = time.Now().Format("02-01-2006T15:04:05")
	newPayment.UpdatedAt = time.Now().Format("02-01-2006T15:04:05")

	return ps.Database.AddPayment(newPayment)

}

func (ps PaymentService) GetPayment() []entity.Payment {
	return ps.Database.GetPayment()

}

func (ps PaymentService) GetSinglePayment(id string) (entity.Payment, error) {
	paymentDetails := ps.Database.GetSinglePayment(id)
	for _, payment := range paymentDetails {
		if payment.PaymentId == id {
			return payment, nil

		}
	}
	return entity.Payment{}, fmt.Errorf("Payment not found")

}

func (ps PaymentService) DeletePayment(id string) error {
	err := ps.Database.DeletePayment(id)

	if err != nil {
		return err
	}

	return nil

}
