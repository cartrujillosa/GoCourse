package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/pabloos/http/user"
)

const (
	URL = "https://localhost:8080"
)

func main() {

	withCache := flag.Bool("cache", false, "cache")
	var userURL string

	if *withCache {
		userURL = fmt.Sprintf("%s/%s", URL, "users_with_cache")
	} else {
		userURL = fmt.Sprintf("%s/%s", URL, "users_with_no_cache")
	}

	user1 := user.User{
		ID:      "001",
		Lessons: []string{"Concurrencia", "HTTP"},
	}

	user2 := user.User{
		ID:      "001",
		Lessons: []string{"the next one"},
	}

	users := []user.User{user1, user2}

	client := newClient()
	for _, u := range users {
		buf := new(bytes.Buffer)
		json.NewEncoder(buf).Encode(u)

		resp, err := client.Post(userURL, "application/json; charset=utf-8", buf)
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()

		io.Copy(os.Stdout, resp.Body)

		fmt.Fprintf(os.Stdout, "Status code: %d\n", resp.StatusCode)
	}
}
