package api

import (
	"net/http"
	"time"

	"github.com/event-management/src/lib/msgqueue"
	"github.com/event-management/src/lib/store"
	"github.com/gorilla/mux"
)

func ServeAPI(listenAddr string, database store.DatabaseHandler, eventEmitter msgqueue.EventEmitter) {
	r := mux.NewRouter()
	r.Methods("post").Path("/events/{eventID}/bookings").Handler(&CreateBookingHandler{eventEmitter, database})

	srv := http.Server{
		Handler:      r,
		Addr:         listenAddr,
		WriteTimeout: 2 * time.Second,
		ReadTimeout:  1 * time.Second,
	}

	srv.ListenAndServe()
}
