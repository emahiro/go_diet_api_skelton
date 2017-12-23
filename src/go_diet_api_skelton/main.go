package main

import (
	"fmt"
	"go_diet_api_skelton/web"
	"net/http"
)

var port = 8080

func main() {
	router := web.Router()
	http.ListenAndServe(fmt.Sprintf(":%v", port), router)
}
