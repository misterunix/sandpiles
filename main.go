package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
	"time"
)

// Rec : Simple struct for holding the bounding box.
type Rec struct {
	MinX int
	MinY int
	MaxX int
	MaxY int
}

var maxwidth int     // maxwidth : The starting width of the working grid.
var maxheight int    // maxheight : The starting height of the working grid.
var maxsize int      // maxsize : maxwidth * maxheight
var centerwidth int  // centerwidth : Centerpoint for X
var centerheight int // centerheight : Centerpoint for Y

var grid1 []uint8 // grid1 : Main working grid.

var wrec Rec // wrec : Main working bounding box.

// topple : Searches the bounding box for cells that contain grains. If a cell
// has less than 4 grains then just copy the number over to the tempory grid. If
// a cell has 4 or more then topple by removing 4 and adding 1 to each of the
// cardinal directions. (Nasty gotha, dont overwrite the tempgrid. Add the grains
// then subtract 4.) If one of the grains lands outside the bounding box,
// increase the main bounding box. Continue looping through all the cells until
// there are no more cells with 4 or more grains left.
func topple() bool {

	// bail : If any cell has more than 4 grains then bail will be false. If no
	// cell has 4 or more (no changes) then set to true. Returned to allow the
	// calling routine to know status.
	var bail bool

	// grid2 : Temporary grid for holding results of the topples.
	var grid2 = make([]uint8, maxsize)

	bail = true // default to signaling there were no topples.

	// loop through the grid and if any cell has less than 4 then just copy the
	// contents to the tempory grid.
	for y := wrec.MinY; y <= wrec.MaxY; y++ {
		for x := wrec.MinX; x <= wrec.MaxX; x++ {
			index := y*maxwidth + x
			if grid1[index] < 4 {
				grid2[index] = grid1[index]
			}
		}
	}

	w := wrec // w : Temporary bounding box.

	// Check all the bounding box edges and fix to maxwidth and height. Grains
	// will be lost if the exceed the hard limits.
	if wrec.MinX < 0 {
		wrec.MinX = 0
	}
	if wrec.MinY < 0 {
		wrec.MinY = 0
	}
	if wrec.MaxX > maxwidth {
		wrec.MaxX = maxwidth
	}
	if wrec.MaxY > maxheight {
		wrec.MaxY = maxheight
	}

	// Loop through the main grid. If any cell has 4 or more then topple them.
	for y := w.MinY; y <= w.MaxY; y++ {
		for x := w.MinX; x <= w.MaxX; x++ {

			index := y*maxwidth + x // Standard x,y to position
			num := grid1[index]     // num : holds the value of the current position.

			// if num has 4 or more, do it to it!
			if num >= 4 {
				grid2[index] += (num - 4) // Nasty gotha here. Add then subtract!
				if grid2[index] >= 4 {
					bail = false // A cell had 4 or more.
				}

				// west
				if x-1 >= 0 {
					tx := x - 1
					grid2[y*maxwidth+tx]++
					if tx < wrec.MinX {
						wrec.MinX = tx // modifiy the working bounding box because of a spill over.
					}
					if grid2[y*maxwidth+tx] >= 4 {
						bail = false // A cell has 4 or more.
					}
				}

				// north
				if y-1 >= 0 {
					ty := y - 1
					grid2[ty*maxwidth+x]++
					if ty < wrec.MinY {
						wrec.MinY = ty // modifiy the working bounding box because of a spill over.
					}
					if grid2[ty*maxwidth+x] >= 4 {
						bail = false // A cell has 4 or more.
					}
				}

				// east
				if x+1 <= maxwidth-1 {
					tx := x + 1
					grid2[y*maxwidth+tx]++
					if tx > wrec.MaxX {
						wrec.MaxX = tx // modifiy the working bounding box because of a spill over.
					}
					if grid2[y*maxwidth+tx] >= 4 {
						bail = false // A cell has 4 or more.
					}
				}

				//south
				if y+1 <= maxheight-1 {
					ty := y + 1
					grid2[ty*maxwidth+x]++
					if ty > wrec.MaxY {
						wrec.MaxY = ty // modifiy the working bounding box because of a spill over.
					}
					if grid2[ty*maxwidth+x] >= 4 {
						bail = false // A cell has 4 or more.
					}
				}
			}
		}
	}

	grid1 = grid2 // Copy the temporary grid to the working grid.

	// return status of grid to calling routine.
	return bail
}

func main() {

	fmt.Println("Program started.")
	start := time.Now()

	//seed := time.Now().UnixNano()
	//randomSource := rand.NewSource(seed)
	//rnd := rand.New(randomSource)

	maxwidth = 1920  // originally for image size, now just grid size.
	maxheight = 1920 // originally for image size, now just grid size.
	maxsize = maxheight * maxwidth
	grid1 = make([]uint8, maxsize) // allocate memory for main grid.

	// set center variables.
	centerwidth = maxwidth / 2
	centerheight = maxheight / 2

	// shift : Easy place to set the number of bits to shift.
	shift := 18

	// grains : The number of grains that will be feed to the starting locations.
	grains := 1 << shift

	// setup the working bounding box.
	wrec.MaxX = 0
	wrec.MaxY = 0
	wrec.MinX = maxwidth
	wrec.MinY = maxheight

	// cindex : Standard routine to conver x,y to a position. Currently used as
	// base for feeding the grid.
	cindex := centerheight*maxwidth + centerwidth

	// prime the grid
	grid1[cindex-100] = 128
	grid1[cindex+100] = 128

	// Scan grid to find working rectangle
	for y := 0; y < maxheight; y++ {
		for x := 0; x < maxwidth; x++ {
			index := y*maxwidth + x
			if grid1[index] != 0 && x < wrec.MinX {
				wrec.MinX = x
			}
			if grid1[index] != 0 && y < wrec.MinY {
				wrec.MinY = y
			}
			if grid1[index] != 0 && x > wrec.MaxX {
				wrec.MaxX = x
			}
			if grid1[index] != 0 && y > wrec.MaxY {
				wrec.MaxY = y
			}
		}
	}

	// Give a little wiggle room for the working bounding box.
	wrec.MinX--
	wrec.MinY--
	wrec.MaxX++
	wrec.MaxY++

	fmt.Printf("Min X:Y %d:%d Max X:Y %d:%d\n", wrec.MinX, wrec.MinY, wrec.MaxX, wrec.MaxY)

	frame := 0 // frame : the number of times, topple() has been run. Why frame... because...
	ty := 0    // ty : used for the spinner.

	// Main loop for feeding the grid. The grid is uint8 so feed the grid 128 grains per loop.
	for outerloop := 0; outerloop < grains; outerloop += 128 {

		for {
			t := topple() // do it
			// if topple returned true then exit the for loop, but not the outer loop.
			if t {
				break
			}

			frame++ // topple was run, inc counter.

			// Check frame modulo number to update the spinner.
			if frame%100 == 0 {

				var char string

				switch ty {
				case 0:
					char = "|"
				case 1:
					char = "/"
				case 2:
					char = "-"
				case 3:
					char = "\\"
				case 4:
					char = "|"
				case 5:
					char = "/"
				case 6:
					char = "-"
				case 7:
					char = "\\"
				}

				fmt.Printf("%s\r", char)
				ty++
				if ty == 7 {
					ty = 0
				}
			}
		}

		// topple exited out of the inner loop. Add 129 to the feed points on
		// the grid. Allow the outer loop to check if enough grains have been
		// set.
		grid1[cindex-100] += 128
		grid1[cindex+100] += 128
	}

	fmt.Printf("Min X:Y %d:%d Max X:Y %d:%d\n", wrec.MinX, wrec.MinY, wrec.MaxX, wrec.MaxY)

	// make bounding box a little larger so there is padding for the image.
	wrec.MinX -= 10
	wrec.MinY -= 10
	wrec.MaxX += 10
	wrec.MaxY += 10

	// make sure the bounding box doesnt exceed the image size limit.
	if wrec.MinX < 0 {
		wrec.MinX = 0
	}
	if wrec.MinY < 0 {
		wrec.MinY = 0
	}
	if wrec.MaxX > maxwidth {
		wrec.MaxX = maxwidth
	}
	if wrec.MaxY > maxheight {
		wrec.MaxY = maxheight
	}

	// calculate the image width and height
	iwidth := wrec.MaxX - wrec.MinX  // iwidth : The image's width
	iheight := wrec.MaxY - wrec.MinY // iheight : The image's height

	img := image.NewRGBA(image.Rect(0, 0, iwidth, iheight))
	bgcolor := color.RGBA{R: 0, G: 0, B: 0, A: 0xFF}
	draw.Draw(img, img.Bounds(), &image.Uniform{bgcolor}, image.Point{}, draw.Src)

	// dump some good stuff in markdaown format for readme.
	fmt.Printf("- %d 1<<%d\n", grains, shift)
	fmt.Printf("  - Time %s\n", time.Since(start))
	fmt.Println("  - Frames:", frame)
	fmt.Printf("  - w:%d h:%d\n", iwidth, iheight)
	fmt.Println()

	x := 0 // x : image's x location.
	y := 0 // y : image's y location.

	// Loop through the bounding box, setting the pixels.
	for wy := wrec.MinY; wy < wrec.MaxY; wy++ {
		for wx := wrec.MinX; wx < wrec.MaxX; wx++ {

			num := grid1[wy*maxwidth+wx]
			switch num {
			case 0:
				c := color.RGBA{R: 0x47, G: 0x2e, B: 0x74, A: 0xff}
				img.Set(x, y, c)
			case 1:
				c := color.RGBA{R: 0x31, G: 0x3a, B: 0x75, A: 0xff}
				img.Set(x, y, c)
			case 2:
				c := color.RGBA{R: 0xaa, G: 0x8a, B: 0x39, A: 0xff}
				img.Set(x, y, c)
			case 3:
				c := color.RGBA{R: 0xaa, G: 0x9c, B: 0x39, A: 0xff}
				img.Set(x, y, c)
			case 4:
				c := color.RGBA{R: 0x00, G: 0xff, B: 0x00, A: 0xff}
				img.Set(x, y, c)
			}

			x++
		}
		y++
		x = 0
	}

	pid := os.Getpid()
	//
	fn := fmt.Sprintf("images/%011d-%06d.png", grains, pid)
	//fn := fmt.Sprintf("images/test-%d.png", pid)
	f, err := os.Create(fn)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	// Encode to `PNG` with `DefaultCompression` level
	// then save to file
	err = png.Encode(f, img)
	if err != nil {
		log.Fatalln(err)
	}

}
