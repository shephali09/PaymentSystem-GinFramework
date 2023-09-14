package main

import (
	"net/http"
	"paymentapi/config"
	"paymentapi/controller"
	"paymentapi/service"
	database "paymentapi/storage"

	"github.com/gin-gonic/gin"
)

func main() {

	config := config.Init()

	portNumber := config.String("portNumber")

	paymentObjectController := controller.PaymentController{
		Service: service.PaymentService{
			Database: database.PaymentDatabase{},
		},
	}

	router := gin.Default()
	router.GET("/hi", helloHandler)
	router.POST("/addPayment", paymentObjectController.AddPayment)
	router.GET("/getPayment", paymentObjectController.GetPayment)
	router.GET("/getSinglePayment/:id", paymentObjectController.GetSinglePayment)
	router.DELETE("/deletePayment/:id", paymentObjectController.DeletePayment)

	router.Run(":" + portNumber)

}

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello",
	})
}
