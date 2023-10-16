package main

import (
	"log"
	"os"

	"github.com/Athunlal/Website-Health-Checker-GO.git/app"
)

func main() {

	app := app.Start()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
