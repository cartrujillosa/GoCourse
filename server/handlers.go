package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/pabloos/http/user"
)

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "You are on the index page\n")
}

func userLessonsHandler(w http.ResponseWriter, r *http.Request) {
	cached, ok := r.Context().Value("cached").(bool)
	if !ok || !cached {
		var u user.User
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		user.AddUser(u.ID, u.Lessons)
	}
	response, err := json.Marshal(user.GetUsers())
	if err != nil {
		http.Error(w, err.Error(), 422)
		return
	}
	w.Write([]byte(string(response)))
	return
}
