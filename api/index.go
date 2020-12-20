package handler

import (
	"log"
	"net/http"

	"github.com/rkobayashi/project-euler-profile/profsvg"
)

// Handler is api entry point.
func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	if err := profsvg.Write(w, r.URL.Query()); err != nil {
		log.Println("generate svg error: ", err)
	}
}
