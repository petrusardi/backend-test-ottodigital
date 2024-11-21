package services

import (
	"backend-test/db"
	domain "backend-test/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBrand(c *gin.Context) {
	var brand domain.Brand
	if err := c.ShouldBindJSON(&brand); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if db.DB == nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Database not connected"})
        return
    }

	err := db.DB.QueryRow(domain.CreateBrandSQLQuery, brand.Name).Scan(&brand.ID, &brand.Name)
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create brand"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": brand.ID, "name": brand.Name})
}

func CreateVoucher(c *gin.Context) {
	var voucher domain.Voucher
	if err := c.ShouldBindJSON(&voucher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := db.DB.QueryRow(domain.CreateVoucherSQLQuery, voucher.Name, voucher.Cost, voucher.BrandID, voucher.Code, voucher.Description, voucher.Value).Scan(&voucher.ID, &voucher.Name, &voucher.Cost, &voucher.BrandID, &voucher.Code, &voucher.Description, &voucher.Value, &voucher.CreatedAt)
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create voucher"})
		return
	}

	c.JSON(http.StatusOK, voucher)
}

func GetVoucher(c *gin.Context) {
	id := c.Query("id")

	var voucher domain.Voucher
	err := db.DB.QueryRow(domain.GetVoucherSQLQuery, id).Scan(&voucher.ID, &voucher.Name, &voucher.Cost, &voucher.BrandID, &voucher.Code, &voucher.Description, &voucher.Value, &voucher.CreatedAt)	
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Voucher not found"})
		return
	}

	c.JSON(http.StatusOK, voucher)
}

func GetVouchersByBrand(c *gin.Context) {
	brandID := c.Query("id")
	rows, err := db.DB.Query(domain.GetAllVoucherByBrand, brandID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch vouchers"})
		return
	}
	defer rows.Close()

	var vouchers []domain.Voucher
	for rows.Next() {
		var voucher domain.Voucher
		err := rows.Scan(&voucher.ID, &voucher.Name, &voucher.Cost, &voucher.BrandID, &voucher.Code, &voucher.Description, &voucher.Value, &voucher.CreatedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing vouchers"})
			return
		}
		vouchers = append(vouchers, voucher)
	}

	c.JSON(http.StatusOK, vouchers)
}

func MakeTransaction(c *gin.Context) {
	var transaction domain.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := db.DB.QueryRow(domain.CreateTransactionQuery, transaction.CustomerName).Scan(&transaction.ID, &transaction.CustomerName, &transaction.RedemptionDate, &transaction.CreatedAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to make redemption"})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

func GetTransactionDetail(c *gin.Context) {
	id := c.Query("id")

	var transaction domain.Transaction
	err := db.DB.QueryRow(domain.GetTransactionQuery, id).Scan(&transaction.ID, &transaction.CustomerName, &transaction.RedemptionDate, &transaction.CreatedAt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

func CreateRedemption(c *gin.Context) {
	var redemption domain.Redemption
	if err := c.ShouldBindJSON(&redemption); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := db.DB.QueryRow(domain.CreateRedemptionQuery, redemption.TransactionID, redemption.VoucherID).Scan(&redemption.ID, &redemption.TransactionID, &redemption.VoucherID, &redemption.CreatedAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to make redemption"})
		return
	}

	c.JSON(http.StatusOK, redemption)
}

func GetRedemptionDetail(c *gin.Context) {
	transactionID := c.Query("transactionId")
	rows, err := db.DB.Query(domain.GetRedemptionQuery, transactionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch vouchers"})
		return
	}
	defer rows.Close()

	var redemptions []domain.Redemption
	var totalCost int = 0 
	var totalValue float64 = 0 
	for rows.Next() {
		var redemption domain.Redemption
		err := rows.Scan(&redemption.ID, &redemption.TransactionID, &redemption.VoucherID, &redemption.CreatedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing vouchers"})
			return
		}

		var voucher domain.Voucher
		err = db.DB.QueryRow("SELECT * FROM vouchers WHERE id = $1", redemption.VoucherID).
			Scan(&voucher.ID, &voucher.Name, &voucher.Cost, &voucher.BrandID, &voucher.Code, &voucher.Description, &voucher.Value, &voucher.CreatedAt)
			fmt.Println(err)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching voucher details"})
			return
		}

		totalCost += voucher.Cost
		totalValue += voucher.Value

		redemption.Voucher = voucher

		redemptions = append(redemptions, redemption)
	}

	c.JSON(http.StatusOK, gin.H{
		"redemptions": redemptions,
		"totalCost":   totalCost,
		"totalValue":  totalValue,
	})
}