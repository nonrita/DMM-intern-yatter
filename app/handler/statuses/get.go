package statuses

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	entity, err := h.sr.FindByID(r.Context(), id)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	m, _ := json.MarshalIndent(entity, "", "  ")
	fmt.Fprint(w, string(m))
}