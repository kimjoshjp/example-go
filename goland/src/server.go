package main

import (
  "os"
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("content-type", "text/html")
    fmt.Fprintf(w, "&#60;h2&#62;Hello World from Distelli! You have a working Go application Deployment!&#60;/h2&#62;")
}

func main() {
  port := os.Args[1]
  fmt.Println(port)

    http.HandleFunc("/", handler)
    http.ListenAndServe(":"+port, nil)
}
