package main

import "fmt"
import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
    pixels := make([][]uint8, dy)
    for y := range pixels {
        pixels[y] = make([]uint8, dx)
        for x := range pixels[y] {
            pixels[y][x] = uint8(x ^ y)
        }
    }
    fmt.Println(pixels)
    return pixels
}

func main() {
	pic.Show(Pic)
}
