package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	var err error
	port := 8080
	dir := "."
	mux := http.NewServeMux()

	portS := os.Getenv("PORT")
	if len(portS) > 0 {
		if port, err = strconv.Atoi(portS); err != nil {
			log.Fatal(err)
		}
	}
	dirS := os.Getenv("DIR")
	if len(dirS) > 0 {
		if _, err := os.Stat(dirS); os.IsNotExist(err) {
			log.Fatal(err)
		} else {
			dir = dirS
		}
	}
	if dir, err = filepath.Abs(dir); err != nil {
		log.Fatal(err)
	}

	mux.Handle("/", http.FileServer(http.Dir(dir)))
	fmt.Printf("serving the directory '%s' at :%d\n", dir, port)
	if err = http.ListenAndServe(fmt.Sprintf(":%d", port), mux); err != nil {
		log.Fatal(err)
	}
}
