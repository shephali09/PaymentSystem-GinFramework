package controller

import (
	"fmt"
	"net/http"
	"paymentapi/entity"
	"paymentapi/service"

	"github.com/gin-gonic/gin"
)

type PaymentController struct {
	Service service.PaymentService
}

func (pc PaymentController) AddPayment(c *gin.Context) {

	var newPayment entity.Payment
	err := c.BindJSON(&newPayment)
	if err != nil {
		fmt.Println("Error Occured", err)
		return
	}
	//newPayment.CreatedAt = time.Now().Format("02-01-2006T15:04:05")
	//newPayment.UpdatedAt = time.Now().Format("02-01-2006T15:04:05")
	newPayment = pc.Service.AddPayment(newPayment)

	//PaymentDetails = append(PaymentDetails, newPayment)
	c.JSON(http.StatusCreated, newPayment)

}

func (pc PaymentController) GetPayment(c *gin.Context) {
	paymentDetails := pc.Service.GetPayment()
	c.JSON(http.StatusOK, paymentDetails)

}

func (pc PaymentController) GetSinglePayment(c *gin.Context) {
	id := c.Params.ByName("id")

	payment, err := pc.Service.GetSinglePayment(id)
	if err != nil {
		c.String(http.StatusBadRequest, "Payment not found")
	}
	c.JSON(http.StatusOK, payment)
}

func (pc PaymentController) DeletePayment(c *gin.Context) {
	fmt.Println("Inside delete function")
	fmt.Println(c.Params)
	id := c.Params.ByName("id")

	fmt.Println(id)

	err := pc.Service.DeletePayment(id)
	if err != nil {
		c.String(http.StatusBadRequest, "Payment Deleted")
	}
	/*if isFound {
		err := pc.Service.DeletePayment(id)
		if err != nil {
			c.String(http.StatusBadRequest, "Id not Found")
		}
	} else {
		c.String(http.StatusBadRequest, "Id not found, Please add the id.")
	}*/

}
