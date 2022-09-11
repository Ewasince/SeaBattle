package main

import (
	"fmt"
	"strconv"
)

const (
	weightSc int = 2
	heightSc int = 1
)

type Field [FSize][FSize]tile

type Screen struct {
	ownField   Field
	enemyField Field
	helper     Field
	ownAlive   int
	enemyAlive int
}

type Ship [][2]int

func makeScreen() Screen {
	sc := Screen{}
	sc.ownField = makeField()
	sc.enemyField = makeField()
	sc.helper = makeField()
	alive := 0
	for _, item := range SHIPS {
		alive += item.size * item.count
	}
	sc.ownAlive = alive
	sc.enemyAlive = alive

	return sc
}

func makeField() Field {
	field := Field{}
	for i := 0; i < FSize; i++ {
		for j := 0; j < FSize; j++ {
			field[i][j] = VoidTile
		}
	}
	return field
}

func (sc *Screen) setShip(tiles Ship) {
	var x int
	var y int
	for _, tile_ := range tiles {
		x = tile_[0]
		y = tile_[1]
		(*sc).ownField[y][x] = ShipTile
		(*sc).helper[y][x] = ShipTile
	}
	(*sc).makeSaveZone(tiles)
}

func (sc *Screen) shoot(x int, y int) bool {
	if (*sc).ownField[y][x] == ShipTile {
		(*sc).ownField[y][x] = HavocTile
		(*sc).ownAlive -= 1
		return true
	} else {
		(*sc).ownField[y][x] = MissTile
		return false
	}
}

func (sc *Screen) generateShips() {
	flagCreate := false
	for _, ship := range SHIPS {
		for i := 0; i < ship.count; i++ {
			flagCreate = false
			for flagCreate == false {
				flagCreate = (*sc).generateShip(ship.size)
				//fmt.Println(helperField)
			}
			//(*sc).enemyField = *helperField
			//(*MainScreen).showScreen()
		}
	}
}

func (sc *Screen) generateShip(size int) bool {
	x := rnd.Intn(FSize)
	y := rnd.Intn(FSize)
	helper := &(*sc).helper

	if !(checkCoord(x, y, helper)) {
		return false
	}

	dirNum := rnd.Intn(4)
	gipoTiles := make(Ship, 0, size)
	var flag bool
	for d := 0; d < 4; d++ {
		direct := DIRECTIONS[(d+dirNum)%4]
		flag = true

		gipoTiles, flag = checkCap(x, y, size, direct, helper)

		if flag {
			break
		}
	}

	if flag {
		(*sc).setShip(gipoTiles)
	}
	return flag
}

func checkCap(x int, y int, size int, direct dir, helper *Field) (gipoTiles Ship, flag bool) {
	for j := 0; j < size; j++ {
		newX := x + direct.x*j
		newY := y + direct.y*j

		if !(checkCoord(newX, newY, helper)) {
			gipoTiles = gipoTiles[:0]
			flag = false
			return
		}
		gipoTiles = append(gipoTiles, [2]int{newX, newY})
	}
	flag = true
	return
}

func checkCoord(x int, y int, field *Field) bool {
	if x >= FSize || x < 0 || y >= FSize || y < 0 {
		return false
	}
	if (*field)[y][x] != VoidTile {
		return false
	}
	return true
}

var oMatrix = [8][2]int{
	{-1, -1}, {0, -1}, {1, -1},
	{1, 0}, {1, 1},
	{0, 1}, {-1, 1}, {-1, 0},
}

func (sc *Screen) makeSaveZone(tiles Ship) {
	for _, tilePr := range tiles {
		for _, coeff := range oMatrix {
			x := tilePr[0] + coeff[0]
			y := tilePr[1] + coeff[1]
			if x >= FSize || x < 0 || y >= FSize || y < 0 {
				continue
			}
			tileLi := &((*sc).helper[y][x])
			if *tileLi == VoidTile {
				*tileLi = MissTile
			}
			//(*MainScreen).showScreen()
		}

	}
}

func (sc *Screen) showScreen() {
	for k := 0; k < 2; k++ {
		for j := 0; j < weightSc; j++ {
			fmt.Print(" ")
		}
		for i := 0; i < FSize; i++ {
			value := condValue(strconv.Itoa(i))

			fmt.Print(value)
		}
		fmt.Print("\t")
	}
	fmt.Print("\n")

	for i := 0; i < FSize; i++ {
		for h := 0; h < heightSc; h++ {
			if h == 0 {
				value := condValue(strconv.Itoa(i))
				fmt.Print(value)
			} else {
				for t := 0; t < weightSc; t++ {
					fmt.Print(" ")
				}
			}
			for j := 0; j < FSize; j++ {
				for w := 0; w < weightSc; w++ {
					fmt.Print(sc.ownField[i][j])
				}
			}
			fmt.Print("\t")

			if h == 0 {
				value := condValue(strconv.Itoa(i))
				fmt.Print(value)
			} else {
				for t := 0; t < weightSc; t++ {
					fmt.Print(" ")
				}
			}
			for j := 0; j < FSize; j++ {
				for w := 0; w < weightSc; w++ {
					fmt.Print(sc.enemyField[i][j])
				}
			}
			fmt.Println()
		}
	}
	textAlive := "Your lives: " + strconv.Itoa(sc.ownAlive)
	indent := FSize*weightSc - len(textAlive)
	for i := 0; i < indent; i++ {
		textAlive += " "
	}
	textAlive += "\tEnemy lives: " + strconv.Itoa(sc.enemyAlive)
	fmt.Println(textAlive)
	fmt.Println()
}

func condValue(value string) string {
	if len(value) > weightSc {
		value = value[:weightSc]
	} else {
		for {
			if len(value) >= weightSc {
				break
			}
			value = " " + value
		}
	}
	return value
}
