package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/LLIEPJIOK/polynomial/internal/application/polynomial"
)

func main() {
	if err := polynomial.Start(); err != nil {
		slog.Error(fmt.Sprintf("polynomial.Start(): %s", err))
		os.Exit(1)
	}
}
