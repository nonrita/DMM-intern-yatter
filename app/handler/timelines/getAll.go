package timelines

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (h *handler) GetAll(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	onlyMedia, err := strconv.ParseBool(params.Get("only_media"))
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	maxID, err := parseParamStoI(params.Get("max_id"), 1)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	sinceID, err := parseParamStoI(params.Get("since_id"), 0)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	limit, err := parseParamStoI(params.Get("limit"), 40)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	data, err := h.tr.FindAll(r.Context(), onlyMedia, *maxID, *sinceID, *limit)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	m, _ := json.MarshalIndent(data, "", "  ")
	fmt.Fprint(w, string(m))
}

func parseParamStoI(s string, defalt int) (*int, error) {
	if s == "" {
		s = strconv.Itoa(defalt)
	}
	rNum, err := strconv.Atoi(s)
	if err != nil {
		return nil, err
	}
	return &rNum, nil
}