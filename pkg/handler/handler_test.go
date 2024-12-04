package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"voucher-api/pkg/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUseCase struct {
	mock.Mock
}

func (m *MockUseCase) CreateBrand(name, description string) (*model.Brand, error) {
	args := m.Called(name, description)
	return args.Get(0).(*model.Brand), args.Error(1)
}

func (m *MockUseCase) CreateVoucher(brandID int, costInPoints float64, voucherCode string) (*model.Voucher, error) {
	args := m.Called(brandID, costInPoints, voucherCode)
	return args.Get(0).(*model.Voucher), args.Error(1)
}

func (m *MockUseCase) GetVoucherByID(id string) (*model.Voucher, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Voucher), args.Error(1)
}

func (m *MockUseCase) GetVouchersByBrand(brandID string) ([]model.Voucher, error) {
	args := m.Called(brandID)
	return args.Get(0).([]model.Voucher), args.Error(1)
}

func (m *MockUseCase) MakeRedemption(customerID int, voucherIDs []int) (*model.Transaction, error) {
	args := m.Called(customerID, voucherIDs)
	return args.Get(0).(*model.Transaction), args.Error(1)
}

func (m *MockUseCase) GetTransactionDetail(transactionID string) (*model.Transaction, error) {
	args := m.Called(transactionID)
	return args.Get(0).(*model.Transaction), args.Error(1)
}

func TestCreateBrandHandler(t *testing.T) {
	mockUseCase := new(MockUseCase)
	h := NewHandler(mockUseCase)

	mockUseCase.On("CreateBrand", "Test Brand", "Test Description").Return(&model.Brand{Name: "Test Brand", Description: "Test Description"}, nil)

	brand := map[string]string{
		"name":        "Test Brand",
		"description": "Test Description",
	}
	body, _ := json.Marshal(brand)
	req, _ := http.NewRequest("POST", "/brand", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()

	h.CreateBrandHandler(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	mockUseCase.AssertExpectations(t)
}

func TestCreateVoucherHandler(t *testing.T) {
	mockUseCase := new(MockUseCase)
	h := NewHandler(mockUseCase)

	mockUseCase.On("CreateVoucher", 1, 100.0, "TESTCODE").Return(&model.Voucher{VoucherCode: "TESTCODE"}, nil)

	voucher := map[string]interface{}{
		"brand_id":       1,
		"cost_in_points": 100.0,
		"voucher_code":   "TESTCODE",
	}
	body, _ := json.Marshal(voucher)
	req, _ := http.NewRequest("POST", "/voucher", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()

	h.CreateVoucherHandler(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	mockUseCase.AssertExpectations(t)
}

func TestGetVoucherHandler(t *testing.T) {
	mockUseCase := new(MockUseCase)
	h := NewHandler(mockUseCase)

	mockUseCase.On("GetVoucherByID", "1").Return(&model.Voucher{ID: 1}, nil)

	req, _ := http.NewRequest("GET", "/voucher?id=1", nil)
	rr := httptest.NewRecorder()

	h.GetVoucherHandler(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	mockUseCase.AssertExpectations(t)
}

func TestGetVouchersByBrandHandler(t *testing.T) {
	mockUseCase := new(MockUseCase)
	h := NewHandler(mockUseCase)

	mockUseCase.On("GetVouchersByBrand", "1").Return([]model.Voucher{{ID: 1}}, nil)

	req, _ := http.NewRequest("GET", "/vouchers/brand?id=1", nil)
	rr := httptest.NewRecorder()

	h.GetVouchersByBrandHandler(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	mockUseCase.AssertExpectations(t)
}

func TestMakeRedemptionHandler(t *testing.T) {
	mockUseCase := new(MockUseCase)
	h := NewHandler(mockUseCase)

	mockUseCase.On("MakeRedemption", 1, []int{1, 2}).Return(&model.Transaction{ID: 1}, nil)

	redemption := map[string]interface{}{
		"customer_id": 1,
		"voucher_ids": []int{1, 2},
	}
	body, _ := json.Marshal(redemption)
	req, _ := http.NewRequest("POST", "/redemption", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()

	h.MakeRedemptionHandler(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	mockUseCase.AssertExpectations(t)
}

func TestGetTransactionDetailHandler(t *testing.T) {
	mockUseCase := new(MockUseCase)
	h := NewHandler(mockUseCase)

	mockUseCase.On("GetTransactionDetail", "1").Return(&model.Transaction{ID: 1}, nil)

	req, _ := http.NewRequest("GET", "/transaction/redemption?transactionId=1", nil)
	rr := httptest.NewRecorder()

	h.GetTransactionDetailHandler(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	mockUseCase.AssertExpectations(t)
}
