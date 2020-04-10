package customer

import "github.com/percypham/go-rest-template/pkg/models"

// Repository for customer use cases
type Repository interface {
	Create(models.Customer) (models.Customer, error)
	FindByEmail(email string) models.Customer
	FindByMobileNo(mobileNo string) models.Customer
	// FindByID(id string) (*models.Customer, error)
	// Update(models.Customer) error
}
