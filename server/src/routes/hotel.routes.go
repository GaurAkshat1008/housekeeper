package routes

import (
	"net/http"
	"src/index/src/utils"

	"github.com/gorilla/mux"
)

func InitHotelRoutes() *mux.Router {

	utils.Logger.Info("Initializing hotel routes")

	router := mux.NewRouter()
	router.HandleFunc("/hotels", getHotels).Methods("GET")

	return router
}

func getHotels(w http.ResponseWriter, r *http.Request) {
	// Get all hotels
}
