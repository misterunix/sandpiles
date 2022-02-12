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

type Rec struct {
	MinX int
	MinY int
	MaxX int
	MaxY int
}

var maxwidth int
var maxheight int
var maxsize int
var centerwidth int
var centerheight int

var grid1 []int

var wrec Rec

func topple() bool {
	var bail bool
	var grid2 = make([]int, maxsize)

	bail = true

	for y := wrec.MinY; y <= wrec.MaxY; y++ {
		for x := wrec.MinX; x <= wrec.MaxX; x++ {
			index := y*maxwidth + x
			if grid1[index] < 4 {
				grid2[index] = grid1[index]
			}
		}
	}

	w := wrec
	//for y := 0; y < maxheight; y++ {
	//	for x := 0; x < maxwidth; x++ {
	for y := w.MinY; y <= w.MaxY; y++ {
		for x := w.MinX; x <= w.MaxX; x++ {

			index := y*maxwidth + x
			num := grid1[index]
			if num >= 4 {
				grid2[index] += (num - 4)
				if grid2[index] >= 4 {
					bail = false
				}

				if x-1 >= 0 {
					tx := x - 1
					grid2[y*maxwidth+tx]++
					if tx < wrec.MinX {
						wrec.MinX = tx
					}
					if grid2[y*maxwidth+tx] >= 4 {
						bail = false
					}
				}
				if y-1 >= 0 {
					ty := y - 1
					grid2[ty*maxwidth+x]++
					if ty < wrec.MinY {
						wrec.MinY = ty
					}
					if grid2[ty*maxwidth+x] >= 4 {
						bail = false
					}
				}

				if x+1 <= maxwidth-1 {
					tx := x + 1
					grid2[y*maxwidth+tx]++
					if tx > wrec.MaxX {
						wrec.MaxX = tx
					}
					if grid2[y*maxwidth+tx] >= 4 {
						bail = false
					}
				}
				if y+1 <= maxheight-1 {
					ty := y + 1
					grid2[ty*maxwidth+x]++
					if ty > wrec.MaxY {
						wrec.MaxY = ty
					}
					if grid2[ty*maxwidth+x] >= 4 {
						bail = false
					}
				}
			}
		}
	}

	grid1 = grid2
	/*
		for i := 0; i < maxsize; i++ {
			grid1[i] = grid2[i]
			if grid1[i] >= 4 {
				bail = false
			}
			//	grid2[i] = 0
		}
	*/
	return bail
}

func printboard() {
	for y := 0; y < maxheight; y++ {
		for x := 0; x < maxwidth; x++ {
			index := y*maxwidth + x
			fmt.Printf("%02X ", grid1[index])
		}
		fmt.Println()
	}

	fmt.Println()
}

func main() {

	fmt.Println("Program started.")
	start := time.Now()
	//seed := time.Now().UnixNano()
	//randomSource := rand.NewSource(seed)
	//rnd := rand.New(randomSource)

	maxwidth = 1920
	maxheight = 1920
	centerwidth = maxwidth / 2
	centerheight = maxheight / 2
	maxsize = maxheight * maxwidth

	wrec.MaxX = 0
	wrec.MaxY = 0
	wrec.MinX = maxwidth
	wrec.MinY = maxheight

	grid1 = make([]int, maxsize)
	//	grid2 = make([]int, maxsize)

	index := centerheight*maxwidth + centerwidth
	grid1[index] = 5000
	//printboard()

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

	wrec.MinX--
	wrec.MinY--
	wrec.MaxX++
	wrec.MaxY++

	fmt.Printf("Min X:Y %d:%d Max X:Y %d:%d\n", wrec.MinX, wrec.MinY, wrec.MaxX, wrec.MaxY)

	frame := 0
	ty := 0
	for {
		t := topple()
		//printboard()
		if t {
			break
		}
		frame++
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

	/*
		minx := maxwidth
		maxx := 0
		miny := maxheight
		maxy := 0
		for y := 0; y < maxheight; y++ {
			for x := 0; x < maxwidth; x++ {
				if grid1[y+maxwidth+x] != 0 {
					if x > maxx {
						maxx = x
					}
					if x < minx {
						minx = x
					}
					if y > maxy {
						maxy = y
					}
					if y < miny {
						miny = y
					}
				}
			}
		}

		fmt.Printf("minx %d miny %d\nmaxx %d maxy %d\n", minx, miny, maxx, maxy)
	*/

	/*
	   F2F3AE
	   EDD382
	   FC9E4F
	   FF521B
	   020122
	*/

	fmt.Printf("Min X:Y %d:%d Max X:Y %d:%d\n", wrec.MinX, wrec.MinY, wrec.MaxX, wrec.MaxY)
	fmt.Println("Frames:", frame)

	//iwidth := wrec.MaxX - wrec.MinX
	//iheight := wrec.MaxY - wrec.MinY

	img := image.NewRGBA(image.Rect(0, 0, maxwidth, maxheight))
	bgcolor := color.RGBA{R: 0, G: 0, B: 0, A: 0xFF}
	draw.Draw(img, img.Bounds(), &image.Uniform{bgcolor}, image.Point{}, draw.Src)
	for y := 0; y < maxheight; y++ {
		for x := 0; x < maxwidth; x++ {
			num := grid1[y*maxwidth+x]
			switch num {
			case 0:
				c := color.RGBA{R: 0x3d, G: 0x31, B: 0x5b, A: 0xff}
				img.Set(x, y, c)
			case 1:
				c := color.RGBA{R: 0x44, G: 0x4b, B: 0x6e, A: 0xff}
				img.Set(x, y, c)
			case 2:
				c := color.RGBA{R: 0x70, G: 0x8b, B: 0x75, A: 0xff}
				img.Set(x, y, c)
			case 3:
				c := color.RGBA{R: 0x9a, G: 0xb8, B: 0x7a, A: 0xff}
				img.Set(x, y, c)
			}
		}
	}

	fmt.Println()
	fmt.Println()

	pid := os.Getpid()
	fn := fmt.Sprintf("images/test-%d.png", pid)
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

	fmt.Printf("Time %s\n\n", time.Since(start))

}
