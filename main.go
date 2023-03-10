package main

import (
	"fmt"
	"net/http"

	"github.com/xltank/tom"
)

func main() {
	tom := tom.New()

	tom.Get("/test", mid1, func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("visit /test")
		w.Write([]byte("This is a response from tom" + r.RequestURI))
	})

	tom.Listen(":3002")
}

func mid1(w http.ResponseWriter, r *http.Request) {
	fmt.Println("mid1:", r.URL.Path)
}
