package main

import (
	"fmt"
)

var maxwidth int
var maxheight int
var maxsize int
var centerwidth int
var centerheight int

var grid1 []int

func topple() bool {

	var grid2 = make([]int, maxsize)

	for y := 0; y < maxheight; y++ {
		for x := 0; x < maxwidth; x++ {
			index := y*maxwidth + x
			if grid1[index] < 4 {
				grid2[index] = grid1[index]
			}
		}
	}

	for y := 0; y < maxheight; y++ {
		for x := 0; x < maxwidth; x++ {

			index := y*maxwidth + x
			num := grid1[index]
			if num >= 4 {
				grid2[index] += (num - 4)

				if x-1 >= 0 {
					tx := x - 1
					grid2[y*maxwidth+tx]++
				}
				if y-1 >= 0 {
					ty := y - 1
					grid2[ty*maxwidth+x]++
				}

				if x+1 <= maxwidth-1 {
					tx := x + 1
					grid2[y*maxwidth+tx]++
				}
				if y+1 <= maxheight-1 {
					ty := y + 1
					grid2[ty*maxwidth+x]++
				}
			}
		}
	}
	bail := true
	for i := 0; i < maxsize; i++ {
		grid1[i] = grid2[i]
		if grid1[i] >= 4 {
			bail = false
		}
		//	grid2[i] = 0
	}
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
	//seed := time.Now().UnixNano()
	//randomSource := rand.NewSource(seed)
	//rnd := rand.New(randomSource)

	maxwidth = 17
	maxheight = 17
	centerwidth = maxwidth / 2
	centerheight = maxheight / 2
	maxsize = maxheight * maxwidth

	grid1 = make([]int, maxsize)
	//	grid2 = make([]int, maxsize)

	index := centerheight*maxwidth + centerwidth
	grid1[index] = 333
	printboard()

	for {
		t := topple()
		printboard()
		if t {
			break
		}
	}

}
