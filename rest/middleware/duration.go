package middleware

import (
	"log"
	"net/http"
	"time"
)

func DurationHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("before handler; middleware start")
		start := time.Now()
		handler.ServeHTTP(w, r)
		log.Println("after handler; middleware stops; time elapsed: ", time.Since(start), r.URL.Path)
	})
}
