package main

import (
	"fmt"
	"net/http"
)

func main() {
	//fmt.Println("vim-go")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello, world")
		if err != nil {
			fmt.Println("Error occurred:", err)
		}
		fmt.Println(fmt.Sprintf("Number of bytes written was: %d", n))
	})

	_ = http.ListenAndServe(":8080", nil)
}
