package middleware

import (
	"fmt"
	"log"
	"net/http"
)

func AuthorizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req, err := http.NewRequest("GET", "http://localhost:8081/sessions/", nil)
		if err != nil {
			log.Printf("Error: %s", err.Error())
		}

		req.Header = http.Header{
			"Authorization": r.Header["Authorization"],
		}

		client := http.Client{}
		res, err := client.Do(req)
		if err != nil || res.StatusCode != 204 {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Invalid token")
			return
		}

		next.ServeHTTP(w, r)
	})
}
