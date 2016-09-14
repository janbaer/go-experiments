package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"
)

func main() {
	var port int
	var path string

	flag.IntVar(&port, "port", 8000, "You can provide the port (default 8000)")
	flag.StringVar(&path, "path", ".", "You can define the path, default is the current directory")

	flag.Parse()

	log.Printf("Start fileserver on localhost port %d for directory %s", port, path)
	http.ListenAndServe(":"+strconv.Itoa(port), http.FileServer(http.Dir(path)))
}
