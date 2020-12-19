package main

import (
	"fmt"
	"log"
	"net/http"

	handler "github.com/rkobayashi/project-euler-profile/api"
)

func main() {
	const port = 3000
	addr := fmt.Sprintf("localhost:%d", port)

	fmt.Println("listening on", addr)
	http.HandleFunc("/api", handler.Handler)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal("http listen error:", err)
	}
}
