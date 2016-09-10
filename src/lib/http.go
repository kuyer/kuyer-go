package lib

import (
	"fmt"
	"log"
	"net/http"
)

func Http() {
	http.HandleFunc("/hello", handleHello)
}

func handleHello(w http.ResponseWriter, req *http.Request) {
	log.Println("serving", req.URL)
	fmt.Fprintln(w, "hello, kuyer!")
}
