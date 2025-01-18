package main

import (
	"flag"

	"github.com/joho/godotenv"
	"github.com/nyybl/dendenmushi/lib"
)

var devFlag = flag.Bool("dev", false, "Environment type")
var logger = lib.NewLogger("main")

func main() {
	flag.Parse()

	if *devFlag {
		logger.Print("Envrionment mode set to DEVELOPER, loading variables from .env.local")

		if err := godotenv.Load(".env.local"); err != nil {
			logger.Error("Failed to load environmental variables", err)

		} else {
			logger.Print("Environmental variables loaded from .env.local")
		}
	}
}