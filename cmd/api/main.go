package main

import "github.com/bulhoes1998/irpf-go-api/cmd/api/internal/app"

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	app.BuildApplication()

	return nil
}
