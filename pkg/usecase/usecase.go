package usecase

import (
	"voucher-api/pkg/model"
	"voucher-api/pkg/repository"
)

type UseCase interface {
	CreateBrand(name, description string) (*model.Brand, error)
	CreateVoucher(brandID int, costInPoints float64, voucherCode string) (*model.Voucher, error)
	GetVoucherByID(id string) (*model.Voucher, error)
	GetVouchersByBrand(brandID string) ([]model.Voucher, error)
	MakeRedemption(customerID int, voucherIDs []int) (*model.Transaction, error)
	GetTransactionDetail(transactionID string) (*model.Transaction, error)
}

type useCase struct {
	repo repository.Repository
}

func NewUseCase(repo repository.Repository) UseCase {
	return &useCase{repo: repo}
}

func (uc *useCase) CreateBrand(name, description string) (*model.Brand, error) {
	return uc.repo.CreateBrand(name, description)
}

func (uc *useCase) CreateVoucher(brandID int, costInPoints float64, voucherCode string) (*model.Voucher, error) {
	return uc.repo.CreateVoucher(brandID, costInPoints, voucherCode)
}

func (uc *useCase) GetVoucherByID(id string) (*model.Voucher, error) {
	return uc.repo.GetVoucherByID(id)
}

func (uc *useCase) GetVouchersByBrand(brandID string) ([]model.Voucher, error) {
	return uc.repo.GetVouchersByBrand(brandID)
}

func (uc *useCase) MakeRedemption(customerID int, voucherIDs []int) (*model.Transaction, error) {
	return uc.repo.MakeRedemption(customerID, voucherIDs)
}

func (uc *useCase) GetTransactionDetail(transactionID string) (*model.Transaction, error) {
	return uc.repo.GetTransactionDetail(transactionID)
}
