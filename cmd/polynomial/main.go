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

	// mod := polynomial.New(0b110001)
	// p := polynomial.New(0b1000)

	// fmt.Println(polynomial.Inv(p, mod))
}
