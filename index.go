package handler

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/goreleaser/func/count"
	"github.com/patrickmn/go-cache"
)

const key = "count"

var c = cache.New(30*time.Minute, 40*time.Minute)

func H(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	if r.Method != http.MethodGet {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	cached, found := c.Get(key)
	if found {
		log.Println("from cache")
		fmt.Fprint(w, cached)
		return
	}

	live, err := count.Count(r.Context())
	if err != nil {
		log.Println("failed", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, live)
	log.Println("live")
	c.Set(key, live, cache.DefaultExpiration)
}
