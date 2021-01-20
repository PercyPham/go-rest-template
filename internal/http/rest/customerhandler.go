package rest

import (
	"encoding/json"
	"net/http"

	"github.com/percypham/go-rest-template/internal/models"
	"github.com/percypham/go-rest-template/internal/services/customer"
)

func genSignUpHandler(cs customer.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload struct {
			Email     string `json:"email"`
			MobileNo  string `json:"mobile_no"`
			Password  string `json:"password"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
		}
		json.NewDecoder(r.Body).Decode(&payload)

		customer := models.Customer{
			Email:     payload.Email,
			MobileNo:  payload.MobileNo,
			Password:  payload.Password,
			FirstName: payload.FirstName,
			LastName:  payload.LastName,
		}

		createdCustomer, err := cs.SignUp(customer)
		if err != nil {
			responseError(w, http.StatusBadRequest, err)
			return
		}

		responseSuccess(w, createdCustomer)
	}
}
