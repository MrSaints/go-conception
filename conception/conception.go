package conception

import (
	"io"
)

type Options struct {
	Image   string
	Command []string
	Stdout  io.Writer
	Stderr  io.Writer
}

type Conceiver interface {
	Run(Options) error
}
