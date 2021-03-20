package main

import (
	"log"
	"os"

	"h12.io/pbspec"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	protocArgs := os.Args[1:]
	set, err := pbspec.Load(protocArgs)
	if err != nil {
		return err
	}
	jsonBytes, err := pbspec.ToJSON(set)
	if err != nil {
		return err
	}
	_, err = os.Stdout.Write(jsonBytes)
	return err
}
