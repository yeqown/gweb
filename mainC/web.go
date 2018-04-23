package main

import (
	. "gweb/logger"
	"gweb/router"
	"gweb/utils"

	"net/http"
)

var mux *http.ServeMux

func startServer() {
	router.RegisterHandler()

	server := &http.Server{
		Addr:    utils.Fstring(":%d", _instance.ServerC.Port),
		Handler: router.ApiHdl,
	}
	if err := server.ListenAndServe(); err != nil {
		AppL.Fatal(err.Error())
	}
}

// func InitServeMux() {
// 	mux.Handle("/api")
// 	// mux.Handle("/file", http.FileServer(http.Dir("/Users/yeqiang")))
// }
