package models

import "time"

type Brand struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Transaction struct {
	ID             int    `json:"id"`
	CustomerName   string `json:"customer_name"`
	RedemptionDate string `json:"redemption_date"`
	CreatedAt	  time.Time `json:"created_at"`
}

type Voucher struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	Cost          int       `json:"cost"`
	BrandID       int       `json:"brand_id"`
	Code          string    `json:"code"`
	Description   string    `json:"description"`
	Value         float64       `json:"value"`
	CreatedAt     time.Time `json:"created_at"`
}

type Redemption struct {
	ID           int    `json:"id"`
	TransactionID int    `json:"transaction_id"`
	VoucherID    int    `json:"voucher_id"`
	CreatedAt	time.Time `json:"created_at"`
	Voucher       Voucher   `json:"voucher"`
}

const (
	TABLE_NAME_BRAND       = "brands"
	TABLE_NAME_TRANSACTION = "transactions"
	TABLE_NAME_VOUCHER     = "vouchers"
	TABLE_NAME_REDEMPTION  = "redemptions"

	CreateBrandSQLQuery    = "INSERT INTO " + TABLE_NAME_BRAND + "(name) VALUES ($1) RETURNING *"
	CreateVoucherSQLQuery  = "INSERT INTO " + TABLE_NAME_VOUCHER + "(name, cost, brand_id, code, description, value) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *"
	GetVoucherSQLQuery     = "SELECT * FROM " + TABLE_NAME_VOUCHER + " WHERE id = $1"
	GetAllVoucherByBrand   = "SELECT * FROM " + TABLE_NAME_VOUCHER + " WHERE brand_id = $1"
	CreateTransactionQuery = "INSERT INTO " + TABLE_NAME_TRANSACTION + "(customer_name) VALUES ($1) RETURNING *"
	GetTransactionQuery    = "SELECT * FROM " + TABLE_NAME_TRANSACTION + " WHERE id = $1"
	CreateRedemptionQuery  = "INSERT INTO " + TABLE_NAME_REDEMPTION + "(transaction_id, voucher_id) VALUES ($1, $2) RETURNING *"
	GetRedemptionQuery     = "SELECT * FROM " + TABLE_NAME_REDEMPTION + " WHERE transaction_id = $1"
)