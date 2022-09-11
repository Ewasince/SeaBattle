package main

import "fmt"

type userScreen struct {
	Screen
}

func (sc *userScreen) setUserShips() {
	var (
		x         int
		y         int
		d         string
		direction dir
		dirFlag   bool
	)
	for _, ship := range SHIPS {
		for i := ship.count; i > 0; i-- {
			fmt.Printf("Remain %v %v-deck ships\n", i, ship.size)
			var tiles [][2]int
			var res bool
			for {
				dirFlag = false
				fmt.Print("Type coordinates and direction literal \"{x} {y} {d}\": ")
				//n, err := fmt.Scanf("%d %d %s", &x, &y, &d)
				n, err := fmt.Scanln(&x, &y, &d)
				if n != 3 || err != nil {
					fmt.Printf("Error in input: %v\n", err)
				}
				for _, dir_ := range DIRECTIONS {
					if dir_.label == d {
						direction = dir_
						dirFlag = true
						break
					}
				}
				if !(dirFlag) {
					fmt.Println("Wrong direction literal!")
					continue
				}
				tiles, res = checkCap(x, y, ship.size, direction, &(*sc).helper)
				if res {
					break
				}
				fmt.Println("The ship doesn't fit!")
			}
			(*sc).setShip(tiles)

			(*sc).showScreen()
		}
	}
}

func makeUserScreen() userScreen {
	var screen userScreen
	screen.Screen = makeScreen()
	return screen
}
