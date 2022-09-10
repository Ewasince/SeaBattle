package main

import (
	"fmt"
)

type tile string

const (
	Void   tile = "."
	Ship   tile = "#"
	Havoc  tile = "X"
	FSize  int  = 10
	Helper tile = "*"
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
	x     int
	y     int
	label string
}

var DIRECTIONS = [...]dir{
	dir{1, 0, "r"},  // right
	dir{0, 1, "d"},  // down
	dir{-1, 0, "l"}, // left
	dir{0, -1, "г"}, // up
}

var MainScreen *Screen

func main() {
	defer defrFunc()

	userScreen := makeScreen()

	botScreen := makeScreen()
	MainScreen = &botScreen
	botScreen.setShips()

	userScreen.showScreen()
	//fmt.Println("")
	var (
		x         int
		y         int
		d         string
		direction dir
		dirFlag   bool
	)
	fmt.Println("Please set your ships!")
	for _, ship := range SHIPS {
		for i := ship.count; i > 0; i-- {
			fmt.Printf("Remain %v %v-deck ships\n", i, ship.size)
			var tiles [][2]int
			var res bool
			for {
				dirFlag = false
				fmt.Print("Type ccordinates and direction litteral \"{x} {y} {d}\": ")
				fmt.Scanf("%d %d %s", &x, &y, &d)
				for _, dir_ := range DIRECTIONS {
					if dir_.label == d {
						direction = dir_
						dirFlag = true
						break
					}
				}
				if dirFlag {
					break
				}
				tiles, res = checkCap(x, y, ship.size, direction, &userScreen.ownField) // TODO: доедлать ввод координат
				if res == true {
					break
				}
			}
		}
	}
}

func defrFunc() {
	fmt.Print("Push return to quit the program")
	test, _ := fmt.Scanln()
	fmt.Println(test)
}
