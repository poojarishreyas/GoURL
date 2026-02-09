package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("server is starting...")
	http.HandleFunc("/Health", writee)
	http.ListenAndServe(":8000", nil)
	

}
	func writee(w http.ResponseWriter , r *http.Request){
		w.Write([]byte("ok"))
	}

