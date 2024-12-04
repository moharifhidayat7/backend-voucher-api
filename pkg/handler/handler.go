package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"voucher-api/pkg/usecase"
)

type Handler struct {
	UseCase usecase.UseCase
}

func NewHandler(uc usecase.UseCase) *Handler {
	return &Handler{UseCase: uc}
}

func (h *Handler) CreateBrandHandler(w http.ResponseWriter, r *http.Request) {
	var brand struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	if err := json.NewDecoder(r.Body).Decode(&brand); err != nil {
		http.Error(w, fmt.Sprintf("Invalid request body: %v", err), http.StatusBadRequest)
		return
	}

	createdBrand, err := h.UseCase.CreateBrand(brand.Name, brand.Description)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating brand: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdBrand)
}

func (h *Handler) CreateVoucherHandler(w http.ResponseWriter, r *http.Request) {
	var voucher struct {
		BrandID      int     `json:"brand_id"`
		CostInPoints float64 `json:"cost_in_points"`
		VoucherCode  string  `json:"voucher_code"`
	}

	if err := json.NewDecoder(r.Body).Decode(&voucher); err != nil {
		http.Error(w, fmt.Sprintf("Invalid request body: %v", err), http.StatusBadRequest)
		return
	}

	createdVoucher, err := h.UseCase.CreateVoucher(voucher.BrandID, voucher.CostInPoints, voucher.VoucherCode)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating voucher: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdVoucher)
}

func (h *Handler) GetVoucherHandler(w http.ResponseWriter, r *http.Request) {
	voucherID := r.URL.Query().Get("id")
	if voucherID == "" {
		http.Error(w, "Voucher ID is required", http.StatusBadRequest)
		return
	}

	voucher, err := h.UseCase.GetVoucherByID(voucherID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching voucher: %v", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(voucher)
}

func (h *Handler) GetVouchersByBrandHandler(w http.ResponseWriter, r *http.Request) {
	brandID := r.URL.Query().Get("id")
	if brandID == "" {
		http.Error(w, "Brand ID is required", http.StatusBadRequest)
		return
	}

	vouchers, err := h.UseCase.GetVouchersByBrand(brandID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching vouchers: %v", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(vouchers)
}

func (h *Handler) MakeRedemptionHandler(w http.ResponseWriter, r *http.Request) {
	var redemptionRequest struct {
		CustomerID        int     `json:"customer_id"`
		VoucherIDs        []int   `json:"voucher_ids"`
		TotalCostInPoints float64 `json:"total_cost_in_points"`
	}

	if err := json.NewDecoder(r.Body).Decode(&redemptionRequest); err != nil {
		http.Error(w, fmt.Sprintf("Invalid request body: %v", err), http.StatusBadRequest)
		return
	}

	transaction, err := h.UseCase.MakeRedemption(redemptionRequest.CustomerID, redemptionRequest.VoucherIDs, redemptionRequest.TotalCostInPoints)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error processing redemption: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(transaction)
}

func (h *Handler) GetTransactionDetailHandler(w http.ResponseWriter, r *http.Request) {
	transactionID := r.URL.Query().Get("transactionId")
	if transactionID == "" {
		http.Error(w, "Transaction ID is required", http.StatusBadRequest)
		return
	}

	transaction, err := h.UseCase.GetTransactionDetail(transactionID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching transaction detail: %v", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(transaction)
}
