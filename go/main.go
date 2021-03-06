package main

import (
    "fmt"
    "log"
    "net/http"
)

func getTable(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("username")
	if err != nil {
		fmt.Fprintln(w, "Cannot get cookie")
	}

	// TODO: load data from database

	fmt.Fprintln(w, c)
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("username")
	if err != nil {
		fmt.Fprintln(w, "Cannot get cookie")
	}

	fmt.Fprintln(w, c)
}

func setCookie(w http.ResponseWriter, r *http.Request) {
    log.Println("/cookie/set")
	c := http.Cookie{
		Name:     "username",
		Value:    "bater",
		HttpOnly: true,
        Domain: "localhost",
        Path: "/",
	}
	http.SetCookie(w, &c)
}

func main() {
    http.HandleFunc("/private/table", getTable)
    http.HandleFunc("/cookie/get", getCookie)
    http.HandleFunc("/cookie/set", setCookie)
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        log.Println("/")
        fmt.Fprintf(w, "Hello go!")
    })

    log.Fatal(http.ListenAndServe(":80", nil))
}
