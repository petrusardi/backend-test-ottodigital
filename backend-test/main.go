package main

import (
	"backend-test/db"
	"backend-test/services"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()

	router := gin.Default()

	router.POST("/brand", services.CreateBrand)
	router.POST("/voucher", services.CreateVoucher)
	router.GET("/voucher", services.GetVoucher)
	router.GET("/voucher/brand", services.GetVouchersByBrand)
	router.POST("/transaction", services.MakeTransaction)
	router.GET("/transaction", services.GetTransactionDetail)
	router.POST("transaction/redemption", services.CreateRedemption)
	router.GET("transaction/redemption", services.GetRedemptionDetail)

	router.Run(":8080")
}
