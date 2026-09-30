package racer

import (
	"net/http"
)

func Racer(a, b string) (winner string) {
	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		resp, err := http.Get(url)
		if err == nil {
			resp.Body.Close()
		}
		close(ch)
	}()
	return ch
}

// func measureResponseTime(url string) time.Duration {
// 	start := time.Now()
// 	resp, err := http.Get(url)
// 	if err == nil {
// 		resp.Body.Close()
// 	}
// 	return time.Since(start)
// }
