package routes

import (
	"dumbmerch/handlers"
	"dumbmerch/pkg/mysql"
	"dumbmerch/repositories"

	"github.com/gorilla/mux"
)

// Create ProfileRoutes function and add the route here ...
func ProfileRoutes(r *mux.Router) {
	ProfileRepository := repositories.RepositoryProfile(mysql.DB)
	h := handlers.HandlerProfile(ProfileRepository)

	r.HandleFunc("/profile/{id}", h.GetProfile).Methods("GET")
}
