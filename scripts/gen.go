//go:build ignore

package main

import (
	"flag"
	"github/sarumaj/go-kakasi/pkg/codegen"
	"log"
	"os"
)

var logger = log.New(os.Stderr, "codegen: ", 0)
var buildDir = flag.String("buildDir", "build", "build directory")
var indent = flag.String("indent", "", "indentation string")

func main() {
	flag.Parse()
	logger.Printf("Generating code in %s\n", *buildDir)

	if err := codegen.Generate(*buildDir, *indent); err != nil {
		logger.Fatalln(err)
	}

	logger.Println("Code generation complete")
}
