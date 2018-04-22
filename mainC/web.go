package main

import (
	"gweb/router"
	"net/http"
)

var mux *http.ServeMux

func startServer() {
	router.RegisterHandler()

	server := &http.Server{
		Addr:    ":5050",
		Handler: router.ApiHdl,
	}
	server.ListenAndServe()
}

// func InitServeMux() {
// 	mux.Handle("/api")
// 	// mux.Handle("/file", http.FileServer(http.Dir("/Users/yeqiang")))
// }
