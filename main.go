package main

import (
	"fmt"
	"github.com/aasumitro/go-learn/res"
	"net/http"
)

func main() {
	http.HandleFunc("/", res.GetExampleData)
	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
