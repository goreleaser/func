package handler

import (
	"fmt"
	"image-to-ascii/count"
	"net/http"
)

func H(w http.ResponseWriter, r *http.Request) {
	c, err := count.Count(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, c)
}
