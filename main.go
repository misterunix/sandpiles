package main

import (
	"flag"
	"fmt"
)

var bMinX int      //The bounding box in lower the X plane.
var bMinY int      //The bounding box in lower the Y plane.
var bMaxX int      //The bounding box in upper the X plane.
var bMaxY int      //The bounding box in upper the Y plane.
var grid_x int     //The grid size in the X plane.
var grid_y int     //The grid size in the Y plane.
var grid_size int  // The total number of cells in the grid.
var shift int      // number of bits to shift for grain count.
var grid1 []uint32 // The working grid.
// var shift uint32 // The number of bits to shift for grain count.
var pilehalf uint64 = 2147483648

func main() {

	flag.IntVar(&shift, "s", 8, "Number of bits to shift for grain count.")
	flag.Parse()

	var grains uint64

	grid_x = 1920
	grid_y = 1080
	grid_size = grid_x * grid_y

	// Initialize the working grid.
	grid1 = make([]uint32, grid_size) // Allocate the working grid.

	bMinX = grid_x // Make the min high enough so the true lower bound can be found.
	bMinY = grid_y // Make the min high enough so the true lower bound can be found.
	bMaxX = 0      // Make the max low enough so the true upper bound can be found.
	bMaxY = 0      // Make the max low enough so the true upper bound can be found.

	ip1x := grid_x / 2
	ip1y := grid_y / 2
	pos1 := ip1y*grid_x + ip1x

	if grains < pilehalf {
		grid1[pos1] = uint32(grains)
	} else {
		grid1[pos1] = uint32(pilehalf)
	}

	for y := 0; y < grid_y; y++ {
		for x := 0; x < grid_x; x++ {
			index := y*grid_x + x

			if grid1[index] != 0 {
				if x < bMinX {
					bMinX = x
				}
				if x > bMaxX {
					bMaxX = x
				}
				if y < bMinY {
					bMinY = y
				}
				if y > bMaxY {
					bMaxY = y
				}
			}
		}
	}

	bMaxX++
	bMaxY++
	bMinX--
	bMinY--

	grains_put := pilehalf
	if grains < pilehalf {
		grains_put = grains
	}

	for {
		grains -= grains_put
		topple()
		ratio := float64(grains) / float64(pilehalf)
		fmt.Println("grains: ", grains, " pilehalf: ", pilehalf, " ratio: ", ratio)
		if grains > 0 && grains >= pilehalf {
			grains_put = pilehalf
			grid1[pos1] += uint32(grains_put)
		} else if grains > 0 && grains < pilehalf {
			grains_put = grains
			grid1[pos1] += uint32(grains_put)
		} else {
			break
		}
	}

}

func topple() {

	var bail bool
	bail = false
	for !bail {
		wMinX := bMinX
		wMaxX := bMaxX
		wMinY := bMinY
		wMaxY := bMaxY

		bail = true

		for y := wMinY; y <= wMaxY; y++ {
			for x := wMinX; x <= wMaxX; x++ {
				index := y*grid_x + x
				if grid1[index] >= 4 {
					bail = false

					grid1[index] -= 4

					tyn := y - 1
					if tyn >= 0 {
						t_index := tyn*grid_x + x
						grid1[t_index]++
						if tyn < bMinY {
							bMinY = tyn
						}
					}
					tys := y + 1
					if tys <= grid_y-1 {
						t_index := tys*grid_x + x
						grid1[t_index]++
						if tys > bMaxY {
							bMaxY = tys
						}
					}

					txw := x - 1
					if txw >= 0 {
						t_index := y*grid_x + txw
						grid1[t_index]++
						if txw < bMinX {
							bMinX = txw
						}
					}
					txe := x + 1
					if txe <= grid_x-1 {
						t_index := y*grid_x + txe
						grid1[t_index]++
						if txe > bMaxX {
							bMaxX = txe
						}
					}
				}
			}
		}
	}
}
