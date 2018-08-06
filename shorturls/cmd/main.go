package main

import (
	"log"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/sethetter/go-practice/shorturls"
)

var (
	addr = ":3927"
)

func main() {
	store := shorturls.NewRedisStore(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       1,
	})
	shortener := shorturls.NewShortener(store)

	http.HandleFunc("/", shorturls.RootHandler(shortener))

	log.Printf("Listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
