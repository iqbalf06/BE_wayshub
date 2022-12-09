package routes

import (
	"wayshub/handlers"
	"wayshub/pkg/middleware"
	"wayshub/pkg/mysql"
	"wayshub/repositories"

	"github.com/gorilla/mux"
)

func SubscriptionRoutes(r *mux.Router) {
	subscriptionRepository := repositories.RepositorySubscription(mysql.DB)
	h := handlers.HandlerSubscription(subscriptionRepository)

	r.HandleFunc("/subscribe", middleware.Auth(h.AddSubscription)).Methods("POST")
	r.HandleFunc("/subscribe", middleware.Auth(h.GetSubscription)).Methods("GET")
	r.HandleFunc("/channel/{id}/subscribe/{id}", middleware.Auth(h.Unsubscribe)).Methods("DELETE")
}
