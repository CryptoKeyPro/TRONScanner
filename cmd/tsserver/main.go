package main

import (
	"log"

	"github.com/CryptoKeyPro/TRONScanner/internal/initialize"
)

var version = "v0.0.0"

func main() {
	if err := initialize.Server(); err != nil {
		log.Fatal(err)
	}
	log.Println("TRON Scanner Server", version)
}
