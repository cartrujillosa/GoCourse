package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/pabloos/http/user"
)

func POST(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			h.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func Debug(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer h.ServeHTTP(w, r)

		dump, err := httputil.DumpRequest(r, true)
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
			return
		}

		w.Write([]byte(dump))
	}
}

func Delay(delay time.Duration, h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer h.ServeHTTP(w, r)

		time.Sleep(delay)
	}
}

func Cache(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u user.User
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		if i := user.UserExists(u.ID); i != nil {
			user.UpdateUser(*i, u.Lessons)

			ctx := context.WithValue(r.Context(), "cached", true)
			h.ServeHTTP(w, r.WithContext(ctx))
			return
		} else {
			h.ServeHTTP(w, r)
			return
		}
	}
}
