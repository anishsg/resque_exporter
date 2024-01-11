package main

import (
	"os"
	"github.com/anishsg/resque_exporter"
)

func main() {
	resqueExporter.Run(os.Args[1:])
}
