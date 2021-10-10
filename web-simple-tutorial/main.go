package main

import (
	"fmt"
	"net/http"
	"web-simple-tutorial/handler"
)

const basePath = "/web-tutorial"

func main() {
	server := http.Server{
		Addr:    ":8080",
		Handler: nil,
	}
	http.HandleFunc(basePath+"/", handler.Index)
	http.HandleFunc(basePath+"/api", handler.Json)
	http.HandleFunc(basePath+"/yaml", handler.MysqlConf)
	fmt.Printf("Starting server at port 8080\n")
	server.ListenAndServe()
}
