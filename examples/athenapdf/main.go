package main

import (
	"bytes"
	"fmt"
	"github.com/mrsaints/go-conception/conception"
	"os"
)

func main() {
	// See https://github.com/arachnys/athenapdf/

	client, err := conception.NewDockerConceiver("")
	if err != nil {
		panic(err)
	}

	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	opts := conception.Options{
		Image:   "arachnysdocker/athenapdf:latest",
		Command: []string{"athenapdf", "-S", "-Z2", "https://www.google.com"},
		Stdout:  stdout,
		Stderr:  stderr,
	}
	err = client.Run(opts)
	if err != nil {
		panic(err)
	}

	if stderr.Len() != 0 {
		fmt.Fprintf(os.Stderr, "Error: %+v", stderr)
	}

	if stdout.Len() != 0 {
		fmt.Print(stdout)
	}
}
