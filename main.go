package main

import (
	"fmt"
	"math/rand"
	"time"

	gd "github.com/misterunix/cgo-gd"
)

var ibuf0 *gd.Image
var grid1 []int
var grid2 []int
var grid3 []int
var width int
var height int
var size int
var center int

func gridadd() {
	for i := 0; i < size; i++ {
		grid3[i] = grid1[i] + grid2[i]
		grid1[i] = grid3[i]
		grid2[i] = 0
		grid3[i] = 0
	}
}

func topple() bool {

	didround := false

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {

			index := y*width + x

			if grid1[index] > 3 {

				grid2[index] += grid1[index] - 4
				didround = true
				var tx, ty int

				tx = x - 1
				if tx >= 0 {
					grid2[y*width+tx] += grid1[y*width+tx] + 1 // grid1[y*width+tx] + grid2[y*width+tx] + 1
				}
				tx = x + 1
				if tx < width {
					grid2[y*width+tx] += grid1[y*width+tx] + 1 //grid1[y*width+tx] + grid2[y*width+tx] + 1
				}

				ty = y - 1
				if ty >= 0 {
					grid2[ty*width+x] += grid1[ty*width+x] + 1 //grid1[ty*width+x] + grid2[ty*width+x] + 1
				}
				ty = y + 1
				if ty < height {
					grid2[ty*width+x] += grid1[ty*width+x] + 1 // grid1[ty*width+x] + grid2[ty*width+x] + 1
				}

			}

		}
	}

	if didround {
		for i := 0; i < size; i++ {
			grid1[i] = grid2[i]
			grid2[i] = 0
		}
		//fmt.Println("In topple")
		//printgrid()
	}

	return didround
}

func printgrid() {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			fmt.Printf("%03d ", grid1[y*width+x])
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {

	fmt.Println("Program started.")
	seed := time.Now().UnixNano()
	randomSource := rand.NewSource(seed)
	rnd := rand.New(randomSource)

	width = 1025
	height = 1025

	center = (width - 1) / 2

	size = width * height
	fmt.Printf("width:%d height:%d center:%d size:%d\n", width, height, center, size)

	grid1 = make([]int, size)
	grid2 = make([]int, size)
	grid3 = make([]int, size)

	ibuf0 = gd.CreateTrueColor(width, height)
	var c [5]gd.Color
	c[0] = ibuf0.ColorAllocateAlpha(0x02, 0x01, 0x22, 0)
	c[1] = ibuf0.ColorAllocateAlpha(0xF2, 0xF3, 0xAE, 0)
	c[2] = ibuf0.ColorAllocateAlpha(0xED, 0xD3, 0x82, 0)
	c[3] = ibuf0.ColorAllocateAlpha(0xFC, 0x9E, 0x4F, 0)
	c[4] = ibuf0.ColorAllocateAlpha(0xFF, 0x52, 0x1B, 0)
	ccc := ibuf0.ColorAllocateAlpha(0xFF, 0xff, 0xff, 0)
	//for sn := 1000; sn < 100000; sn = sn + 1000 {
	for i := 0; i < size; i++ {
		grid1[i] = 0
		grid2[i] = 0
		grid3[i] = 0
	}

	//for y := 1025/2 - 100; y < 1025/2+100; y++ {
	//		for x := 1025/2 - 100; x < 1025/2+100; x++ {
	//			index := y*width + x
	//			grid1[index] = rnd.Intn(4096)
	//			grid2[index] = rnd.Intn(4)
	//		}
	//	}

	grid1[center*width+center] = 25 + rnd.Intn(1)
	//grid2[center*width+center] = 0

	//fmt.Println("Start grid")
	//printgrid()
	var frame int
	for {
		//gridadd()
		j := topple()
		//fmt.Println("in loop")
		//printgrid()
		if !j {
			break
		}
		if frame%1000 == 0 {
			fmt.Println(frame)
		}
		frame++
	}

	//fmt.Println("Final")
	//printgrid()

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			index := y*width + x
			if grid1[index] > 4 {
				ibuf0.SetPixel(y, x, ccc)
			} else {
				ibuf0.SetPixel(y, x, c[grid1[index]])
			}
		}
	}
	fn := "images/final.png"
	fmt.Println(fn)
	ibuf0.Png(fn)
	//}
}
