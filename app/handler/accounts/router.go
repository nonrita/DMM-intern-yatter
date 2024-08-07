package accounts

import (
	"net/http"
	"yatter-backend-go/app/domain/repository"
	"yatter-backend-go/app/usecase"

	"github.com/go-chi/chi/v5"
)

// Implementation of handler
type handler struct {
	accountUsecase usecase.Account
	accountRepository repository.Account
}

// Create Handler for `/v1/accounts/`
func NewRouter(au usecase.Account, ar repository.Account) http.Handler {
	r := chi.NewRouter()
	h := &handler{
		accountUsecase: au,
		accountRepository: ar,
	}
	r.Get("/{username}", h.Get)
	r.Post("/", h.Create)

	return r
}
