package main

import (
	"fmt"
)

type tile string

const (
	Void  tile = " "
	Ship  tile = "#"
	Havoc tile = "X"
	FSize int  = 10
)

type sizeCnt struct {
	size  int
	count int
}

var SHIPS = [...]sizeCnt{
	sizeCnt{1, 4},
	sizeCnt{2, 3},
	sizeCnt{3, 2},
	sizeCnt{4, 1},
}

type dir struct {
	x    int
	y    int
	name string
}

var DIRECTIONS = [...]dir{
	dir{1, 0, "right"}, // right
	dir{0, 1, "down"},  // down
	dir{-1, 0, "left"}, // left
	dir{0, -1, "up"},   // up
}

func main() {
	defer fmt.Println("Goodbye!")

	mainScreen := makeScreen()
	botScreen := makeScreen()
	fmt.Println(mainScreen)
	fmt.Println(botScreen)

}
