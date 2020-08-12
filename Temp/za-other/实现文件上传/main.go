// https://www.toutiao.com/i6774759603826590221/
package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const (
	MAX_UPLOAD_SIZE = 1024 * 1024 * 20 // 50MB
)

func main() {
	r := RegisterHandlers()
	http.ListenAndServer(":8080", r)
}

// RegisterHanders ...
func RegisterHanders() *httprouter.Router {
	router := httprouter.New()
	router.POST("/upload", uploadHandler)
	return router
}

func uploadHander(w http.ReponseWriter, r *http.Request, p httprouter.Params)
