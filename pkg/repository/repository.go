package repository

import (
	"voucher-api/pkg/model"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	CreateBrand(name, description string) (*model.Brand, error)
	CreateVoucher(brandID int, costInPoints float64, voucherCode string) (*model.Voucher, error)
	GetVoucherByID(id string) (*model.Voucher, error)
	GetVouchersByBrand(brandID string) ([]model.Voucher, error)
	MakeRedemption(customerID int, voucherIDs []int, totalCostInPoints float64) (*model.Transaction, error)
	GetTransactionDetail(transactionID string) (*model.Transaction, error)
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateBrand(name, description string) (*model.Brand, error) {
	query := `INSERT INTO brands (name, description) VALUES ($1, $2) RETURNING id`
	var brand model.Brand
	err := r.db.Get(&brand, query, name, description)
	if err != nil {
		return nil, err
	}
	brand.Name = name
	brand.Description = description
	return &brand, nil
}

func (r *repository) CreateVoucher(brandID int, costInPoints float64, voucherCode string) (*model.Voucher, error) {
	query := `INSERT INTO vouchers (brand_id, cost_in_points, voucher_code) VALUES ($1, $2, $3) RETURNING id`
	var voucher model.Voucher
	err := r.db.Get(&voucher, query, brandID, costInPoints, voucherCode)
	if err != nil {
		return nil, err
	}
	return &voucher, nil
}

func (r *repository) GetVoucherByID(id string) (*model.Voucher, error) {
	query := `SELECT * FROM vouchers WHERE id = $1`
	var voucher model.Voucher
	err := r.db.Get(&voucher, query, id)
	if err != nil {
		return nil, err
	}
	return &voucher, nil
}

func (r *repository) GetVouchersByBrand(brandID string) ([]model.Voucher, error) {
	query := `SELECT * FROM vouchers WHERE brand_id = $1`
	var vouchers []model.Voucher
	err := r.db.Select(&vouchers, query, brandID)
	if err != nil {
		return nil, err
	}
	return vouchers, nil
}

func (r *repository) MakeRedemption(customerID int, voucherIDs []int, totalCostInPoints float64) (*model.Transaction, error) {
	tx := r.db.MustBegin()

	var transaction model.Transaction
	err := tx.Get(&transaction, `INSERT INTO transactions (customer_id, total_cost_in_points) VALUES ($1, $2) RETURNING id`, customerID, totalCostInPoints)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	for _, voucherID := range voucherIDs {
		_, err := tx.Exec(`INSERT INTO transaction_vouchers (transaction_id, voucher_id) VALUES ($1, $2)`, transaction.ID, voucherID)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	tx.Commit()
	return &transaction, nil
}

func (r *repository) GetTransactionDetail(transactionID string) (*model.Transaction, error) {
	query := `SELECT * FROM transactions WHERE id = $1`
	var transaction model.Transaction
	err := r.db.Get(&transaction, query, transactionID)
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}
