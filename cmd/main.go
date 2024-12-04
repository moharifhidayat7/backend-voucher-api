package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"voucher-api/pkg/handler"
	"voucher-api/pkg/repository"
	"voucher-api/pkg/usecase"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {
	// Load environment variables from the .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	// Get the database URL from the environment variables
	dbURL := os.Getenv("DATABASE_URL")

	// Connect to the database
	db, err := sqlx.Connect("postgres", dbURL)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer db.Close()

	// Initialize repositories, use cases, and handlers
	repo := repository.NewRepository(db)
	uc := usecase.NewUseCase(repo)
	h := handler.NewHandler(uc)

	// Set up the router
	r := mux.NewRouter()

	// Define API routes
	r.HandleFunc("/brand", h.CreateBrandHandler).Methods("POST")
	r.HandleFunc("/voucher", h.CreateVoucherHandler).Methods("POST")
	r.HandleFunc("/voucher", h.GetVoucherHandler).Methods("GET")
	r.HandleFunc("/voucher/brand", h.GetVouchersByBrandHandler).Methods("GET")
	r.HandleFunc("/transaction/redemption", h.MakeRedemptionHandler).Methods("POST")
	r.HandleFunc("/transaction/redemption", h.GetTransactionDetailHandler).Methods("GET")

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}
	fmt.Println("Server started on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
