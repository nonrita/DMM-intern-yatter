package accounts

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	user, err := h.accountRepository.FindByUsername(r.Context(), username)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	if user == nil {
		fmt.Fprint(w, "no user")
		return
	}
	m, _ := json.MarshalIndent(user, "", "  ")
	fmt.Fprint(w, string(m))
}
