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
			x := rand.Intn(FSize)
			y := rand.Intn(FSize)
			if helperScreen.field[x][y] == Void {
				dir := DIRECTIONS[rand.Intn(4)]
				gipoTiles := make([][2]int, 0, 4)
				for j := 0; j < ship.size; j++ {
					newX := x + dir.x*j
					newY := y + dir.y*j
					gipoTiles = append(gipoTiles, [2]int{newX, newY})

					fmt.Println(gipoTiles)
				}
			}
		}
	}
}
