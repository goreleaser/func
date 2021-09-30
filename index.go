package handler

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/goreleaser/func/count"
)

var previous = 0
var lock sync.RWMutex

func H(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	if r.Method != http.MethodGet {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	lock.RLock()
	defer lock.RUnlock()
	if previous > 0 {
		fmt.Fprint(w, previous)
		return
	}

	lock.Lock()
	defer lock.Unlock()
	live, err := count.Count(r.Context())
	if err != nil {
		log.Println("failed", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("live")
	previous = live
	fmt.Fprint(w, live)
	go func() {
		time.Sleep(time.Hour * 6)

		lock.Lock()
		defer lock.Unlock()
		previous = 0
	}()
}
