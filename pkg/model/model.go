package model

type Brand struct {
    ID          int    `db:"id"`
    Name        string `db:"name"`
    Description string `db:"description"`
}

type Voucher struct {
    ID            int     `db:"id"`
    BrandID       int     `db:"brand_id"`
    CostInPoints  float64 `db:"cost_in_points"`
    VoucherCode   string  `db:"voucher_code"`
}

type Transaction struct {
    ID               int     `db:"id"`
    CustomerID       int     `db:"customer_id"`
    TotalCostInPoints float64 `db:"total_cost_in_points"`
}

type TransactionVoucher struct {
    TransactionID int `db:"transaction_id"`
    VoucherID     int `db:"voucher_id"`
}
