package main
import "github.com/ginwan/ecowatch/apps/api/handlers"

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println("EcoWatch API is starting...")

	http.HandleFunc("/health", handlers.GetHealth)

	err := http.ListenAndServe(":8080", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
