package main

import (
	"log"
	"net/http"
	controller "proc-parse/Controller"
)

func Middleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func init() {

}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/api/meminfo", Middleware(http.HandlerFunc(controller.ReadMemInfo)))
	mux.Handle("/api/cpuinfo", Middleware(http.HandlerFunc(controller.ReadCpuInfo)))
	log.Println("Listening on :3000...")
	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)

}
