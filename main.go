package main

import (
	"fmt"
	"prac_go/surface"

	// run "go get github.com/google/uuid" to install this package
	// if you want to store the source code with your project, you
	// can run "go mod vendor".
	"github.com/google/uuid"
)

func main() {
	// package 'fmt' is probably the most widely used Go package. It
	// deals with formatting strings. But don't take my word for it...
	// read the docs! They're very friendly :) https://pkg.go.dev/fmt
	fmt.Println("Hello World")

	// external dependency
	id := uuid.NewString()
	fmt.Println(id)

	// internal dependency
	rectSurface := surface.Rect(4, 8)
	fmt.Println(rectSurface)
}
