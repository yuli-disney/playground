package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
				v := r.URL.Query()
				for _, vs := range v {
						message := "Hello, " + vs[0] + "!"
						fmt.Fprintf(w, message)
				}
    })

    server := &http.Server{Addr: ":8080"}
    if err := server.ListenAndServe(); err != nil {
        log.Println(err)
    }
}