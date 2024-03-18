package main

import (
	"fmt"
	"log/slog"
	"net/http"
)

func main() {
	svr := &http.Server{
		Addr: ":7000",
	}

	err := svr.ListenAndServe()
	if err != nil {
		slog.With(slog.Any("error", err)).Error("Failed to serve")
	}

	fmt.Println("After")
}
