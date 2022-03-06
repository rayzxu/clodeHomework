package main

import (
	"net/http"
	"log"
	"os"
	"strings"
	"net"
	"fmt"
)

func index(w http.ResponseWriter, r *http.Request) {
	os.Setenv("VERSION", "1.0")
	version := os.Getenv("VERSION")

	w.Header().Set("VERSION", version)
	fmt.Println("VERSION:", version)
	for k,v := range r.Header{
		for _, vv := range v{
			fmt.Println(k, ":", vv)
			w.Header().Set(k, vv)
		}
	}
	clientIp := getCurrentIP(r)
	w.Header().Set("clientIp", clientIp)
	fmt.Println("clientIp:", clientIp)
	w.Write([]byte("<h1>Hello World</h1>" + "<h4>"+ clientIp +"</h4>"))
}

func healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Println(200)
} 

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/index", index)
	mux.HandleFunc("/healthz", healthz)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("start http fauld, err: %s\n", err.Error())
	}
	fmt.Println("server started at 8080")
}

func getCurrentIP(r *http.Request) string {
	/* ip := r.Header.Get("X-Real-IP")
	if ip == "" {
		ip = strings.Split(r.RemoteAddr, ":")[0]
	}
	return ip */
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}
	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}