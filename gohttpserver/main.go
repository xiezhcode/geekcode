package main

import (
	"net/http"

	"github.com/xiezhcode/geekcode/gohttpserver/httphandler"
	"github.com/xiezhcode/geekcode/gohttpserver/logger"
)

func main() {

	http.HandleFunc("/server/healthz", httphandler.Healthz)
	http.HandleFunc("/server/stop", httphandler.Stop)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		logger.Logger.Fatal(err)
	}
}
