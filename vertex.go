package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Vertex struct {
	X float64
	Y float64
	Z float64
}

func VertexFromString(s string) *Vertex {

	vertStrings := strings.Split(s, " ")

	fmt.Println(vertStrings)

	if len(vertStrings) != 4 {
		return nil
	}

	x, _ := strconv.ParseFloat(vertStrings[1], 10)
	y, _ := strconv.ParseFloat(vertStrings[2], 10)
	z, _ := strconv.ParseFloat(vertStrings[3], 10)

	return &Vertex{X: x, Y: y, Z: z}

}

func GenerateVerticesStrArray(s []string) []*Vertex {
	vertexArr := make([]*Vertex, 0)
	for i := 0; i < len(s); i++ {
		vert := VertexFromString(s[i])
		if vert == nil {
			continue
		}
		vertexArr = append(vertexArr, VertexFromString(s[i]))
	}

	return vertexArr
}
