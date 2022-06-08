package middleware

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func AuthorizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req, err := http.NewRequest("GET", "http://localhost:8081/sessions", nil)
		if err != nil {
			log.Printf("Error: %s", err.Error())
		}

		req.Header = http.Header{
			"Authorization": r.Header["Authorization"],
		}

		client := http.Client{}
		res, err := client.Do(req)
		if err != nil || res.StatusCode != 200 {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Invalid token")
			return
		}

		bytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Invalid token")
			return
		}
		res.Body.Close()

		ctx := context.WithValue(r.Context(), "account_id", string(bytes))

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
