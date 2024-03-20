package main

import (
	"fmt"
	"os"
	"strings"
)

func readCube() string {
	content, err := os.ReadFile("cube.obj")

	if err != nil {
		fmt.Println("Err")
	}
	return string(content)
}

func splitCube() []string {
	cube := readCube()

	newCube := strings.Split(cube, "\n")
	return newCube
}
