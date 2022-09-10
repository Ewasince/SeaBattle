package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

//	type Field struct {
//		field [10][10]tile
//		alive int
//	}
const (
	weightSc int = 2
	heightSc int = 1
)

type Field [FSize][FSize]tile

type Screen struct {
	ownField   Field
	enemyField Field
	ownAlive   int
	enemyAlive int
}

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func makeScreen() Screen {
	sc := Screen{}
	sc.ownField = makeField()
	sc.enemyField = makeField()
	//for i := 0; i < FSize; i++ {
	//	for j := 0; j < FSize; j++ {
	//		sc.ownField[i][j] = Void
	//		sc.enemyField[i][j] = Void
	//	}
	//}
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
			field[i][j] = Void
		}
	}
	return field
}

func (sc *Screen) setShips() {
	ownField := &sc.ownField
	helperField := makeField()
	//helperField := &sc.enemyField
	flagCreate := false
	for _, ship := range SHIPS {
		for i := 0; i < ship.count; i++ {
			flagCreate = false
			for flagCreate == false {
				flagCreate = generateShip(ship.size, ownField, &helperField)
				//fmt.Println(helperField)
			}
			//(*sc).enemyField = *helperField
			//(*MainScreen).showScreen()
		}
	}
}

func generateShip(size int, field *Field, helper *Field) bool {
	x := rnd.Intn(FSize)
	y := rnd.Intn(FSize)

	if !(checkCoord(x, y, helper)) {
		return false
	}

	dirNum := rnd.Intn(4)
	gipoTiles := make([][2]int, 0, size)
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
		for _, tile_ := range gipoTiles {
			(*field)[tile_[1]][tile_[0]] = Ship
			(*helper)[tile_[1]][tile_[0]] = Ship
		}
		//(*MainScreen).showScreen()
		makeSaveZone(helper, gipoTiles)
	}
	return flag
}

func checkCap(x int, y int, size int, direct dir, helper *Field) (gipoTiles [][2]int, flag bool) {
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
	if (*field)[y][x] != Void {
		return false
	}
	return true
}

var oMatrix = [8][2]int{
	{-1, -1}, {0, -1}, {1, -1},
	{1, 0}, {1, 1},
	{0, 1}, {-1, 1}, {-1, 0},
}

func makeSaveZone(field *Field, tiles [][2]int) {
	for _, tilePr := range tiles {
		for _, coeff := range oMatrix {
			x := tilePr[0] + coeff[0]
			y := tilePr[1] + coeff[1]
			if x >= FSize || x < 0 || y >= FSize || y < 0 {
				continue
			}
			tileLi := &((*field)[y][x])
			if *tileLi == Void {
				*tileLi = Helper
			}
			//(*MainScreen).showScreen()
		}

	}
}

func (sc Screen) showScreen() {
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
