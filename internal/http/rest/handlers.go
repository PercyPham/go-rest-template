package rest

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	CustomerService "github.com/percypham/go-rest-template/internal/services/customer"
	"github.com/percypham/go-rest-template/internal/storage/mongo"
)

// Handler is to handle all http requests. Will return Mux router.
func Handler(cs mongo.CustomerStorage) http.Handler {
	r := mux.NewRouter()

	customerService := CustomerService.New(cs)

	r.HandleFunc("/", indexHandler).Methods("GET")
	r.HandleFunc("/signup", genSignUpHandler(customerService)).Methods("POST")

	return r
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Hello World")
}
