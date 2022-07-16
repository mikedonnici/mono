package api

import (
	"io"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>Hello from Go!</h1>")
}
