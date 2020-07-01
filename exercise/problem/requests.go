package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {
	sites := []string{
		"https://www.google.com",
		"https://drive.google.com",
		"https://maps.google.com",
		"https://hangouts.google.com",
	}

	var wg sync.WaitGroup
	wg.Add(len(sites))
	ctx, cancel := context.WithCancel(context.Background())
	for _, site := range sites {
		go func(ctx context.Context, site string) {
			defer wg.Done()
			res, err := http.Get(site)
			if err != nil {
				cancel()
			}
			select {
			case <-ctx.Done():
				return
			case <-time.After(time.Nanosecond):
				io.WriteString(os.Stdout, res.Status+"\n")
			}
		}(ctx, site)
	}
	wg.Wait()
}
