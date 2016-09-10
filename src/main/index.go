/**
 * go run index.go
 * go build index.go
 */
package main

import (
	"fmt"
	"lib"
	"log"
	"net/http"
)

func main() {
	lib.Http()
	fmt.Println("server start on http://localhost:8888/hello")
	log.Fatal(http.ListenAndServe("localhost:8888", nil))
}
