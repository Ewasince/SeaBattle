package main

import (
	"fmt"
	"math/rand"
)

type Screen struct {
	field [10][10]tile
	alive int
}

func makeScreen() Screen {
	sc := Screen{}
	sc.field = [FSize][FSize]tile{}
	field := [FSize][FSize]tile{}
	for i, arr := range field {
		for j := range arr {
			field[i][j] = Void
		}
	}
	sc.alive = 0
	for _, item := range SHIPS {
		sc.alive += item.size * item.count
	}
	return sc
}

func (sc Screen) setShips() {
	helperScreen := Screen{}
	for _, ship := range SHIPS {
		for i := 0; i < ship.count; i++ {
			fmt.Println(helperScreen)
		}
	}
}

func generateShip(size int, screen *Screen, helper *Screen) {
	x := rand.Intn(FSize)
	y := rand.Intn(FSize)
	direct := DIRECTIONS[rand.Intn(4)]
	gipoTiles := make([][2]int, 0, 4)
	for j := 0; j < ship.size; j++ {
		newX := x + direct.x*j
		newY := y + direct.y*j

		if (*helper).field[newX][newX] != Void {
			flag := false

		}
		gipoTiles = append(gipoTiles, [2]int{newX, newY})
	}

	for _, tile_ := range gipoTiles {

	}

}
