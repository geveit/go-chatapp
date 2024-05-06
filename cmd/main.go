package main

import (
	"fmt"
	"io/fs"
	"log"
	"net/http"

	"github.com/geveit/go-chatapp/frontend"
	"github.com/geveit/go-chatapp/server"
)

func main() {
	s := http.NewServeMux()

	// static files
	s.HandleFunc("/", staticHandler)
	staticFS, _ := fs.Sub(frontend.StaticFiles, "dist")
	httpFS := http.FileServer(http.FS(staticFS))
	s.Handle("/assets/", httpFS)

	// websocket
	hub := server.NewHub()
	go hub.Run()

	s.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		log.Println("WS connection")
		server.HandleWS(hub, w, r)
	})

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", s)
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/favicon.ico" {
		rawFile, _ := frontend.StaticFiles.ReadFile("dist/favicon.ico")
		w.Write(rawFile)
		return
	}
	rawFile, _ := frontend.StaticFiles.ReadFile("dist/index.html")
	w.Write(rawFile)
}
