package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		d := rand.Intn(500) + 500
		time.Sleep(time.Duration(d) * time.Millisecond)
		fmt.Fprintln(w, d)
	})
	fmt.Println("listening on :9090")
	log.Fatalln(http.ListenAndServe(":9090", nil))
}
