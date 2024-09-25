package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	fmt.Println("---> GO -- X")

	fruits := []string{"apple", "banana", "cherry", "blueberry"}

	for fruit_index := 0; fruit_index < len(fruits); fruit_index++ {
		println("Fruit: ", fruits[fruit_index])
	}

	start_time := time.Now()
	sum := 0
	//  how to determin the cpu usage
	for i := 0; i < 10000000; i++ {
		sum += i
	}

	elapsed_time := time.Since(start_time)

	fmt.Println("Elapsed time: ", elapsed_time)

	println("Sum: ", sum)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Whoos! a sd aklsdm asklmdklamdklm aklmsda sdas daksmd klasmd Go!"))
	})

	r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello moto!"))
	})

	http.ListenAndServe(":3000", r)
}
