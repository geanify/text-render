package main

import (
	"fmt"
)

// e29688 at https://www.fileformat.info/info/charset/UTF-8/list.htm?start=8192

func main() {
	cube := splitCube()
	arr := GenerateVerticesStrArray(cube)
	fmt.Println(arr)

	fmt.Println("ello" + "\u258F \u2571")
}
