package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func factorial_recursive(n int) int {
	if n < 0 {
		// Invalid
		return -1
	}

	if n <= 1 {
		return 1
	}

	return n * factorial(n-1)
}

func factorial_iter(n int) int {
	r := 1

	for i := 1; i <= n; i++ {
		r = r * i
	}

	return r
}

func factorial(n int) int {
	return factorial_iter(n)
}

func main() {
	http.HandleFunc("/fact", func(w http.ResponseWriter, r *http.Request) {
		ns := r.URL.Query().Get("n")
		if ns == "" {
			http.Error(w, "please enter n", 400)
			return
		}

		n, err := strconv.Atoi(ns)
		if err != nil {
			http.Error(w, "not a valid number", 400)
		}

		if n < 0 {
			http.Error(w, "please enter a positive n", 400)
			return
		}

		res := factorial(n)

		json.NewEncoder(w).Encode(map[string]int{
			"result": res,
		})
	})

	http.ListenAndServe(":8080", nil)
}
