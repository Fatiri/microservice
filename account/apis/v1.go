package apis

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/microservice/account/driver"
	"github.com/sirupsen/logrus"

	accountHandler "github.com/microservice/account/delivery"
	_accountRepo "github.com/microservice/account/repository"
	_accountService "github.com/microservice/account/service"
)

//API implement to controll the routing
func API() {
	db := driver.Config()

	defer func() {
		err := db.Close()
		if err != nil {
			logrus.Error(err)
			return
		}
	}()

	r := mux.NewRouter().StrictSlash(true)

	r.Handle("/", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logrus.Error("Welcome to Internal API")
		return
	}))).Methods(http.MethodGet)

	accountRepo := _accountRepo.NewAccountRepository(db)
	accountService := _accountService.NewAccountService(accountRepo)
	accountHandler.NewAccountHandler(r, accountService)

	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		logrus.Error("Error path")
		return
	})

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"})
	serverAddress := os.Getenv("SERVER_ADDRESS")

	log.Println("Server start at http://localhost" + serverAddress)
	logrus.Fatalln(http.ListenAndServe(serverAddress, handlers.CORS(headersOk, originsOk, methodsOk)(r)))
}
