package main

import (
	"net/http"
)

type Throttler struct {
	maxRequestsPerSecond int
	//TODO
}

func NewThrottler(maxRequestsPerSecond int) *Throttler {
	// TODO
	return nil
}

func (t *Throttler) Middleware(next http.Handler) http.Handler {
	// TODO
	return nil
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

   Реализуйте точное ограничение (не "примерно N запросов", а строго не более N)
*/
