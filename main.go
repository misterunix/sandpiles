package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"strconv"
	"time"
)

//#define gsize uint32_t // fall back is uint16_t
//#define pilehalf 2147483648 // fall back is 32768

// globals

var bMinX int // bMinX : The bounding box in lower the X plane.
var bMinY int // bMinY : The bounding box in lower the Y plane.
var bMaxX int // bMaxX : The bounding box in the upper X plane.
var bMaxY int // bMaxY : The bounding box in the upper Y plane.

var grid_X int
var grid_Y int
var grid_size int

var shiftb int

var grid1 []uint32

func PrintPNG() {

	bMinX -= 10
	bMinY -= 10
	bMaxX += 10
	bMaxY += 10

	width := bMaxX - bMinX
	height := bMaxY - bMinY

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	fmt.Println("x,y:", bMaxX, bMaxY, bMinX, bMinY)

	for y := bMinY; y < bMaxY; y++ {
		for x := bMinX; x < bMaxX; x++ {
			xx := x - bMinX
			yy := y - bMinY
			img.Set(xx, yy, color.White)
		}
	}

	fmt.Println(bMinX, bMaxX)
	fmt.Println(bMinY, bMaxY)
	fmt.Println(width, height)

	for y := bMinY; y < bMaxY; y++ {
		for x := bMinX; x < bMaxX; x++ {
			xx := x - bMinX
			yy := y - bMinY

			index := y*grid_X + x
			num := grid1[index]

			switch {
			case num == 0:
				img.Set(xx, yy, color.RGBA{18, 72, 249, 255})
			case num == 1:

				img.Set(xx, yy, color.RGBA{115, 170, 249, 255})
			case num == 2:

				img.Set(xx, yy, color.RGBA{255, 192, 0, 255})
			case num == 3:

				img.Set(xx, yy, color.RGBA{124, 0, 0, 255})
			default:
				img.Set(xx, yy, color.Black)
			}

		}

	}

	fn := fmt.Sprintf("images/%v-%d.png", time.Now().Local().Format("20060102150405"), shiftb)
	f, err := os.Create(fn)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	err = png.Encode(f, img)
	if err != nil {
		log.Fatalln(err)
	}

}

func topple() {

	bail := false

	for !bail {
		wMinX := bMinX
		wMaxX := bMaxX
		wMinY := bMinY
		wMaxY := bMaxY

		bail = true

		for y := wMinY; y <= wMaxY; y++ {
			for x := wMinX; x <= wMaxX; x++ {
				index := y*grid_X + x

				if grid1[index] >= 4 {
					bail = false

					grid1[index] -= 4

					tyn := y - 1
					if tyn >= 0 {
						t_index := tyn*grid_X + x
						grid1[t_index]++
						if tyn < bMinY {
							bMinY = tyn
						}
					}
					tys := y + 1
					if tys <= grid_Y-1 {
						t_index := tys*grid_X + x
						grid1[t_index]++
						if tys > bMaxY {
							bMaxY = tys
						}
					}

					txw := x - 1
					if txw >= 0 {
						t_index := y*grid_X + txw
						grid1[t_index]++
						if txw < bMinX {
							bMinX = txw
						}
					}
					txe := x + 1
					if txe <= grid_X-1 {
						t_index := y*grid_X + txe
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

func main() {

	var grains uint32

	shift, _ := strconv.Atoi(os.Args[1])

	shiftb = shift
	grains = 1 << shift // grains : Total number of grains to place on the grid.

	grid_X = 12000              // grid_X : Maximum size of the grid/image in the X
	grid_Y = 12000              // grid_Y : Maximum size of the grid/image in the Y
	grid_size = grid_X * grid_Y // grid_size : total number of units for the grid array.

	fmt.Println("shift", shift)
	fmt.Println("grains", grains)
	fmt.Println("grid_X", grid_X)
	fmt.Println("grid_Y", grid_Y)
	fmt.Println("grid_size", grid_size)

	// grid1 : Array where the grains are stored. Changed to gsize for speed
	// reasons - and memory.
	grid1 = make([]uint32, grid_size) //  new gsize[grid_size];

	// Init the grid array to all 0s
	// shouldnt need to be done in go

	bMinX = grid_X // Make the min high enough so the true lower bound can be found.
	bMaxX = 0      // Make the max low enough so that true upper bound can be found.
	bMinY = grid_Y // make the min high enough so the true lower bound can be found.
	bMaxY = 0      // bMaxY : The bounding box in the upper Y plane.

	ip1x := grid_X / 2
	ip1y := grid_Y / 2
	pos1 := ip1y*grid_X + ip1x

	fmt.Println("ip1x", ip1x)
	fmt.Println("ip1y", ip1y)
	fmt.Println("pos1", pos1)

	if grains < 2147483648 {
		grid1[pos1] = grains
	} else {
		grid1[pos1] = 2147483648
	}

	// find area of grid with starting grains
	for y := 0; y < grid_Y; y++ {
		for x := 0; x < grid_X; x++ {
			index := y*grid_X + x

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

	start := time.Now()

	var grains_put uint32
	grains_put = 2147483648
	if grains < 2147483648 {
		grains_put = grains
	}

	for {
		grains -= grains_put
		topple()
		ratio := float64(grains) / 2147483648.0

		fmt.Println("granis", grains)
		fmt.Println("ratio", ratio)
		fmt.Println("pilehalf", 2147483648)

		if grains > 0 && grains >= 2147483648 {
			grains_put = 2147483648
			grid1[pos1] += grains_put
		} else if grains > 0 && grains < 2147483648 {
			grains_put = grains
			grid1[pos1] += grains_put
		} else {
			break
		}
	}

	fmt.Println("2^", shift, "grains placed")

	timetook := time.Since(start)

	fmt.Println("Time: ", timetook.Seconds())

	PrintPNG()

}
