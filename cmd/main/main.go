package main

import (
	"cache_module/cmd/server"
	"cache_module/internal/cache"
)

func main() {
	c := cache.New("inmemory")
	server.New(c).Listen()
}
