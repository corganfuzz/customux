package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func main() {
	fmt.Println("Server is running in port 8000 ...")

	newMux := http.NewServeMux()

	newMux.HandleFunc("/randomFloat", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, rand.Float64())
	})

	newMux.HandleFunc("/randomInt", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, rand.Intn(100))
	})

	http.ListenAndServe(":8000", newMux)
}
