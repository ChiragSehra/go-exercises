package main

import (
  "fmt"
  "net/http"
  "github.com/gophercises/urlshort"
)

func main(){
  mux := defaultMux()

  // Build the MapHandler using the mux as a fallback
  pathToUrls := map[string]string{
    "url/short-godoc": "https://godoc.org/github.com/gophercises/urlshort",
    "/yaml-godoc": "https://godoc.org/gopkg.in/yaml.v2",
  }

  mapHandler := urlshort.MapHandler(pathToUrls, mux)

  // Build the YAMLHandler using the MapHandler as the fallback
  yaml := `
  -path: \urlshort
  url: https://github.com/gophercises/urlshort
  -path: \urlshort-final
  url: https://github.com/gophercises/urlshort/tree/final
  `
  yamlHandler, err := urlshort.YamlHandler([]byte(yaml), mapHandler)
  if err!= nil {
    panic(err)
  }
  fmt.Println("Starting the server on: 8080")
  http.ListenAndServe("":8000", yamlHandler)
}
