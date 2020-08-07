package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"urlshort"
)

func main() {
	exeType := flag.String("type", "map", "a string")
	flag.Parse()

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	fmt.Println("Starting the server on :8080")
	if *exeType == "map" {
		http.ListenAndServe(":8080", mapHandler)
	} else if *exeType == "yaml" {
		file, err := readFile("./file.yaml")
		if err != nil {
			fmt.Println(err)
		}
		yamlHandler, err := urlshort.YAMLHandler(file, mapHandler)
		if err != nil {
			panic(err)
		}
		http.ListenAndServe(":8080", yamlHandler)
	} else {
		file, err := readFile("./file.json")
		if err != nil {
			fmt.Println(err)
		}
		jsonHandler, err := urlshort.JSONHandler(file, mapHandler)
		if err != nil {
			panic(err)
		}
		http.ListenAndServe(":8080", jsonHandler)
	}

}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func readFile(fp string) ([]byte, error) {
	data, err := ioutil.ReadFile(fp)
	if err != nil {
		fmt.Println("File reading error", err)
		return data, err
	}
	return data, nil
}
