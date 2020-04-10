package mongo

import (
	"context"
	"time"

	"github.com/percypham/go-rest-template/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// CustomerStorage is implementation of Customer Repository
type CustomerStorage struct {
	collection *mongo.Collection
}

// MakeCustomerStorage abc
func MakeCustomerStorage(collection *mongo.Collection) CustomerStorage {
	return CustomerStorage{collection}
}

// Customer asb
type Customer struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Email     string             `bson:"email,omitempty"`
	MobileNo  string             `bson:"mobile_no,omitempty"`
	Password  string             `bson:"password,omitempty"`
	FirstName string             `bson:"first_name,omitempty"`
	LastName  string             `bson:"last_name,omitempty"`
	CreatedAt time.Time          `bson:"created_at,omitempty"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty"`
}

// Create abc
func (cs CustomerStorage) Create(cus models.Customer) (models.Customer, error) {
	currentTime := time.Now()
	customer := Customer{
		Email:     cus.Email,
		MobileNo:  cus.MobileNo,
		Password:  cus.Password,
		FirstName: cus.FirstName,
		LastName:  cus.LastName,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}

	insertedOne, err := cs.collection.InsertOne(context.TODO(), customer)
	if err != nil {
		return models.Customer{}, err
	}

	objID, _ := insertedOne.InsertedID.(primitive.ObjectID)
	cus.ID = objID.Hex()
	cus.CreatedAt = currentTime
	cus.UpdatedAt = currentTime

	return cus, nil
}

// FindByEmail abcases
func (cs CustomerStorage) FindByEmail(email string) models.Customer {
	filter := bson.D{{"email", email}}
	return cs.findByFilter(filter)
}

// FindByMobileNo abcases
func (cs CustomerStorage) FindByMobileNo(mobileNo string) models.Customer {
	filter := bson.D{{"mobile_no", mobileNo}}
	return cs.findByFilter(filter)
}

func (cs CustomerStorage) findByFilter(filter interface{}) models.Customer {
	var foundCustomer Customer

	err := cs.collection.FindOne(context.TODO(), filter).Decode(&foundCustomer)
	if err != nil {
		return models.Customer{}
	}

	result := models.Customer{
		ID:        foundCustomer.ID.Hex(),
		Email:     foundCustomer.Email,
		MobileNo:  foundCustomer.MobileNo,
		Password:  foundCustomer.Password,
		FirstName: foundCustomer.FirstName,
		LastName:  foundCustomer.LastName,
		CreatedAt: foundCustomer.CreatedAt,
		UpdatedAt: foundCustomer.UpdatedAt,
	}

	return result
}
