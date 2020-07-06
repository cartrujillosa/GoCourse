package main

import (
	"net/http"
)

func newMux() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/users_with_cache", Cache(POST(userLessonsHandler)))
	mux.HandleFunc("/users_with_no_cache", POST(userLessonsHandler))

	return mux
}
