package delivery

import (
	"net/http"
	"os"

	"github.com/Fatiri/common/message"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/microservice/account/models"
	"github.com/microservice/account/service"
	"github.com/sirupsen/logrus"
)

//AccountHandler struct
type AccountHandler struct {
	AccountService service.Service
}

//NewAccountHandler route endpoint with mux
func NewAccountHandler(r *mux.Router, accountService service.Service) {
	claimHandler := &AccountHandler{
		AccountService: accountService,
	}

	v1 := r.PathPrefix("/v1/account").Subrouter()

	v1.Handle("", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(claimHandler.createDataAccount))).Methods(http.MethodPost)
}

//createDataAccount function create data account
func (c *AccountHandler) createDataAccount(w http.ResponseWriter, r *http.Request) {
	account := new(models.Account)

	fullName := r.FormValue("fullname")
	place := r.FormValue("price")

	account.Fullname = fullName
	account.Place = place

	response, err := c.AccountService.CreateNewAccount(account)
	if err != nil {
		logrus.Error(err)
		return
	}

	message.JSON(w, http.StatusOK, response)
	return
}
