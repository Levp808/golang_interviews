package main

import (
	"net/http"
	"sync/atomic"
	"time"
)

type Throttler struct {
	maxRequestsPerSecond int
	currRequsts          atomic.Uint32
}

func NewThrottler(maxRequestsPerSecond int) *Throttler {
	trotler := &Throttler{
		maxRequestsPerSecond: maxRequestsPerSecond,
		currRequsts:          atomic.Uint32{},
	}
	trotler.currRequsts.Store(0)

	go func() {
		timer := time.Tick(1 * time.Second)

		for range timer {
			trotler.currRequsts.Store(0)
		}
	}()

	return trotler
}

func (t *Throttler) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if t.currRequsts.Load() > uint32(t.maxRequestsPerSecond) {
			w.WriteHeader(http.StatusTooManyRequests)
			return
		}
		t.currRequsts.Add(1)

		next.ServeHTTP(w, r)
	})
}

func main() {
	throttler := NewThrottler(5)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	handler := throttler.Middleware(mux)

	http.ListenAndServe(":8080", handler)
}

/*
   Создайте Throttler middleware, который принимает параметр maxRequestsPerSecond

   Если лимит запросов превышен, middleware должен возвращать HTTP 429 (Too Many Requests)

   Middleware должен быть потокобезопасным
*/
