package main

import (
	"hex_ddd_app/db"
	"hex_ddd_app/internal/client"
	"hex_ddd_app/internal/handlers/rest"
	"hex_ddd_app/internal/repositories"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// db := repositories.NewInMemoryDB()
	pg, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		log.Fatalf(err.Error())
	}

	pg.AutoMigrate(&db.Customer{}, &db.Appointment{}, &db.Service{})

	customerRepo := repositories.NewPostgresDB(pg)

	customerService := client.NewCustomerService(customerRepo)

	custHandler := rest.NewHTTPCustomerHandler(customerService)

	r := mux.NewRouter()

	r.HandleFunc("/customers/{id}", custHandler.Get).Methods(http.MethodGet)
	r.HandleFunc("/customers", custHandler.Create).Methods(http.MethodPost)

	log.Println("listening at port 4200...")
	panic(http.ListenAndServe(":4200", r))
}
