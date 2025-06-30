package routes

import (
	"github.com/gorilla/mux"
	"github.comn/Moballigh5225/Bookstore-Management-APIS/pkg/controllers"

)

RegisterBookStoreRoutes := func(router *mux.Router){
	router.HandlerFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandlerFunc("/book/", controllers.GetBook).Methods("GET")
}
