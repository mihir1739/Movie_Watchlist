package main

import (
	"fmt"
	"net/http"

	"githum.com/mihir1739/mongoapi/router"
)

func main() {
	fmt.Println("Hello!")
	r := router.Router()
	fmt.Println("server starting...")
	http.ListenAndServe(":4000", r)
	fmt.Println("Listening at Port 4000 ...")
}
