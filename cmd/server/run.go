package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/percypham/go-rest-template/configs"
	"github.com/percypham/go-rest-template/internal/http/rest"
	"github.com/percypham/go-rest-template/internal/storage/mongo"
)

// Run server
func Run() {
	client := mongo.Connect(configs.MongoURI)

	database := client.Database(configs.MongoDBName)
	customerCollection := database.Collection("customers")
	cs := mongo.MakeCustomerStorage(customerCollection)

	routers := rest.Handler(cs)
	http.Handle("/", routers)

	fmt.Println("Server is listening on port", configs.Port)
	portToServe := ":" + strconv.Itoa(configs.Port)
	http.ListenAndServe(portToServe, nil)
}
