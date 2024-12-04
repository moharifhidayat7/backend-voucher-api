# Backend Voucher API

## Introduction
This is a backend service for managing vouchers. The service is built with Go and uses PostgreSQL as the database. It provides endpoints for creating brands, vouchers, and handling voucher redemptions.

## Prerequisites
- Go 1.19 or later
- PostgreSQL
- Make

## Installation

1. **Clone the repository**
    ```sh
    git clone https://github.com/moharifhidayat7/backend-voucher-api.git
    cd backend-voucher-api
    ```

2. **Set up environment variables**

    Create a `.env` file in the root directory and add the following variables:
    ```
    DATABASE_URL=your_postgresql_database_url
    PORT=your_preferred_port
    ```

3. **Install dependencies**
    ```sh
    make deps
    ```
3. **Run Database Migration**
    ```sh
    make migrate
    ```

## Running the Application

1. **Run the application**
    ```sh
    make run
    ```

## Makefile Commands

- To install dependencies:
    ```sh
    make deps
    ```

- To run the application:
    ```sh
    make run
    ```

- To clean up:
    ```sh
    make clean
    ```
- To rollback database:
    ```sh
    make rollback
    ```

## API Documentation

### Create a Brand
- **Endpoint**: `/brand`
- **Method**: `POST`
- **Request**:
    ```sh
    curl -X POST http://localhost:8080/brand \
    -H "Content-Type: application/json" \
    -d '{
        "name": "BrandName",
        "description": "Brand Description"
    }'
    ```

### Create a Voucher
- **Endpoint**: `/voucher`
- **Method**: `POST`
- **Request**:
    ```sh
    curl -X POST http://localhost:8080/voucher \
    -H "Content-Type: application/json" \
    -d '{
        "brand_id": 1,
        "voucher_code": "VOUCHERCODE",
        "cost_in_points": 10000
    }'
    ```

### Get Voucher by Id
- **Endpoint**: `/voucher`
- **Method**: `GET`
- **Request**:
    ```sh
    curl -X GET http://localhost:8080/voucher?id=1
    ```

### Get Vouchers by Brand
- **Endpoint**: `/voucher/brand`
- **Method**: `GET`
- **Request**:
    ```sh
    curl -X GET http://localhost:8080/voucher/brand \
    -H "Content-Type: application/json" \
    -d '{
        "id": 1
    }'
    ```

### Make a Redemption
- **Endpoint**: `/transaction/redemption`
- **Method**: `POST`
- **Request**:
    ```sh
    curl -X POST http://localhost:8080/transaction/redemption \
    -H "Content-Type: application/json" \
    -d '{
        "voucher_code": "VOUCHER_CODE",
        "customer_id": 1
    }'
    ```

### Get Transaction Detail
- **Endpoint**: `/transaction/redemption`
- **Method**: `GET`
- **Request**:
    ```sh
    curl -X GET http://localhost:8080/transaction/redemption \
    -H "Content-Type: application/json" \
    -d '{
        "transaction_id": 1
    }'
    ```
