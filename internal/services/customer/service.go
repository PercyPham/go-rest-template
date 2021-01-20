package customer

import "github.com/percypham/go-rest-template/internal/models"

// Service of customer
type Service interface {
	SignUp(models.Customer) (models.Customer, error)
}

type service struct {
	repo Repository
}

// New creates a new customer service
func New(repo Repository) Service {
	return service{repo}
}
