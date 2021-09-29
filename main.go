package main

import (
	"log"
	"net/http"
	controller "proc-parse/Controller"
)

type EndPoints struct {
	Url string
	Fun func(w http.ResponseWriter, r *http.Request)
}

func Middleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

var Urls []EndPoints

func init() {
	Urls = make([]EndPoints, 2)
	Urls[0] = EndPoints{Url: "/api/meminfo", Fun: controller.ReadMemInfo}
	Urls[1] = EndPoints{Url: "/api/cpuinfo", Fun: controller.ReadCpuInfo}
}

func main() {
	mux := http.NewServeMux()
	for _, url := range Urls {
		mux.Handle(url.Url, Middleware(http.HandlerFunc(url.Fun)))
	}
	log.Println("Listening on :3000...")
	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)

}
