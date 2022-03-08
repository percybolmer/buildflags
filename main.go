package main

import (
	"log"

	"programmingpercy.tech/buildflags/version"
)

var (
	runner = "client-1.0.0"
)

func main() {

	log.Println("Starting runner", runner, "version", version.Version)
}
