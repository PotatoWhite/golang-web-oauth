package main

import (
	"github.com/potatowhite/web/study-oauth/router"
	"net/http"
)

func main() {
	http.ListenAndServe(":3000", router.Router())
}
