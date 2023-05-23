package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rohitq50/go-lang-project/config"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	BaseRoute = "api/v1"
)

func InitializeRoutes(router *mux.Router, dbConnection *mongo.Client, config *config.Configuration) {
	userRepository := UserRepository.New(dbConnection, conf)
	userService := UserService.New(userRepository)
	userAPI := api.NewUserAPI(userService)

	// Routes

	// -------------------------- User APIs ------------------------------------
	router.HandleFunc(BaseRoute+"/users/me", userAPI.Get).Methods(http.MethodGet)
	router.HandleFunc(BaseRoute+"/users", userAPI.Update).Methods(http.MethodPut)
}
