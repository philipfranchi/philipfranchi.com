package main

import (
    "log"
    "net/http"
    "os"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
	    port = "5000"
    }

    f, _ := os.Create("/var/log/golang/golang-server.log")
    defer f.Close()
    log.SetOutput(f)

    fs := http.FileServer(http.Dir("./public")) 
    http.Handle("/", fs)

    log.Printf("Listening on port %s\n\n", port)
    for ;; {
        http.ListenAndServe(":"+port, nil)
    }
}
