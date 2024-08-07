package statuses

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	fmt.Fprint(w, id)
}