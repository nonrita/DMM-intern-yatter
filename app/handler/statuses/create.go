package statuses

import (
	"encoding/json"
	"fmt"
	"net/http"
	"yatter-backend-go/app/domain/auth"
)

// Request body for `POST /v1/statuses`
type AddRequest struct {
	Content string
	Url string
}

// Handle request for `POST /v1/statuses`
func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	var req AddRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	account_info := auth.AccountOf(r.Context()) // 認証情報を取得する

	dto, err := h.su.Create(r.Context(), int(account_info.ID) , req.Content, &req.Url)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(dto.Status); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
