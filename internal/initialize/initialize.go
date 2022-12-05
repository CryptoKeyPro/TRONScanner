package initialize

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

const packageTitle = "initialize: "

func Server() error {
	const funcTitle = packageTitle + "Server"

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	if err := godotenv.Load(); err != nil {
		return errors.Wrap(err, funcTitle)
	}

	return nil
}
