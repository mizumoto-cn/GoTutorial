package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //parse arguments, you have to call this manually
	fmt.Println(r.Form) //Form contains the parsed form data, including both the URL field's query parameters and the PATCH, POST, or PUT form data. This field is only available after ParseForm is called. The HTTP client ignores Form and uses Body instead.
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello Miz!") //send data w to client side
}

func main() {
	http.HandleFunc("/", sayhelloName)       // set router
	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// try http://localhost:9090/
// try http://localhost:9090/?url_long=111&url_long=222
