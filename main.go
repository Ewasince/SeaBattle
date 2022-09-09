package main

import (
	"fmt"
	"math/rand"
)

type tile string
const (
	Void  tile = " "
	Ship  tile = "#"
	Havoc tile = "X"
	FSize int = 10
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

type dir struct{
	x int
	y int
}
var DIRECTIONS = [...]dir{
	dir{1,0}, // left
	dir{0,1}, // down
	dir{-1,0}, // right
	dir{0,-1}, // up
}

func main() {
	defer fmt.Println("Goodbye!")

	mainScreen := makeScreen()
	botScreen := makeScreen()
	fmt.Println(mainScreen)
	fmt.Println(botScreen)

}

type Screen struct {
	field [10][10]tile
	alive int
}

func makeScreen() Screen {
	sc := Screen{}
	sc.field = [FSize][FSize]tile{}
	for i, j := range sc.field{ // TODO: проверить как работает
		sc.field[i][j] = Void
	}
	sc.alive = 0
	for _, item := range SHIPS {
		sc.alive += item.size * item.count
	}
	return sc
}

func (sc Screen) setShips() {
	helperScreen := Screen{}
	for _, ship := range SHIPS{
		for i := 0; i< ship.count; i++{
			x := rand.Intn(FSize)
			y := rand.Intn(FSize)
			if helperScreen[x][y] != Void
		}
	}
	sc.
}
