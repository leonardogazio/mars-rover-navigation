package main

import (
	"flag"
	"fmt"

	"github.com/leonardogazio/mars_rover_navigation/rover"
	"github.com/leonardogazio/mars_rover_navigation/server"
)

func main() {
	rover.PlateauInstance = rover.NewPlateau(700, 1000)

	runAsAPI := flag.Bool("api", false, "Defines if will start as gRPC api. Usage: api=true|false")
	inputFile := flag.String("input", "", "Input parameter file path, use when on CLI mode(api=false).")
	flag.Parse()

	if *runAsAPI {
		server.StartGRPC()
	}

	out, err := rover.ParseFile(*inputFile)
	if err != nil {
		panic(err)
	}

	fmt.Println(out)
}
