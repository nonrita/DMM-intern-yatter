package timelines

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type handler struct {
}

func test (w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "time")
}

func NewRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/", test)

	return r
}
