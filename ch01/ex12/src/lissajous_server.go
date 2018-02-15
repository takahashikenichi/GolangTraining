package main

import (
	"log"
	"net/http"
	"time"
	"math/rand"
	"strconv"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {

	var cycles int
	param_cycles := r.Form["cycles"]
	if 0 == len(param_cycles) || "" == param_cycles[0] {
		cycles = 5
	} else {
		cycles, _ = strconv.Atoi(param_cycles[0])
	}	
	lissajous(w, cycles)
}

