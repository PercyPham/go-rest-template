package customer

import (
	"errors"

	"github.com/percypham/go-rest-template/internal/models"
	"github.com/percypham/go-rest-template/internal/services/utils"
)

func (s service) SignUp(cus models.Customer) (models.Customer, error) {
	// TODO: implement validation
	if cus.Email == "" {
		return models.Customer{}, errors.New("Email must be provided")
	}

	if !utils.IsEmail(cus.Email) {
		return models.Customer{}, errors.New("Has wrong email format")
	}

	if cus.MobileNo != "" && !utils.IsMobileNo(cus.MobileNo) {
		return models.Customer{}, errors.New("Has wrong mobile no format")
	}

	if len(cus.Password) < 6 {
		return models.Customer{}, errors.New("Password must have at least 6 charactes")
	}

	foundCustomer := s.repo.FindByEmail(cus.Email)
	if !foundCustomer.IsEmpty() {
		return models.Customer{}, errors.New("Email is already used")
	}

	if cus.MobileNo != "" {
		foundCustomer = s.repo.FindByMobileNo(cus.MobileNo)
		if !foundCustomer.IsEmpty() {
			return models.Customer{}, errors.New("MobileNo is already used")
		}
	}

	cus.Password = utils.HashSHA256(cus.Password)

	createdCustomer, err := s.repo.Create(cus)

	createdCustomer.Password = "------"

	return createdCustomer, err
}
