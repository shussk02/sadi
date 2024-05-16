package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	var port = os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}
	fmt.Println("Hello, World!")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	})

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println(err)
	}

}
