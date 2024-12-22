package main

import (
  "fmt"
  "net/http"
)

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello from Go Backend!")
  })
  http.ListenAndServe(":8080", nil)

  http.HandleFunc("/tony", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
  
    html := `<!DOCTYPE html>
    <html lang="en">
    <head>
      <meta charset="UTF-8" />
      <meta name="viewport" content="width=device-width, initial-scale=1.0" />
      <title>Tony's Video</title>
    </head>
    <body>
      <h1>This is Tony's special route!</h1>
      <iframe width="560" height="315"
        src="https://www.youtube.com/embed/aAkMkVFwAoo"
        title="YouTube video player"
        frameborder="0"
        allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
        allowfullscreen>
      </iframe>
    </body>
    </html>`
  
    fmt.Fprint(w, html)
  })  
}