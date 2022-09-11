package main

import (
	"fmt"
	"math/rand"
	"time"
)

type tile string

const (
	VoidTile  tile = "."
	ShipTile  tile = "#"
	HavocTile tile = "X"
	FSize     int  = 10
	MissTile  tile = "*"
)

type sizeCnt struct {
	size  int
	count int
}

var SHIPS = [...]sizeCnt{
	{1, 4},
	{2, 3},
	{3, 2},
	{4, 1},
}

type dir struct {
	x     int
	y     int
	label string
}

var DIRECTIONS = [...]dir{
	{1, 0, "r"},  // right
	{0, 1, "d"},  // down
	{-1, 0, "l"}, // left
	{0, -1, "u"}, // up
}

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

//var UserScreen *Screen

func main() {
	defer deferFunc()

	userScreen := makeUserScreen()

	botScreen := makeBotScreen()
	botScreen.generateShips()

	userScreen.showScreen()
	var (
		x int
		y int
	)
	fmt.Println("Please set your ships!")
	//setting user ships cycle
	for {
		fmt.Println("Would you like auto set ships? (\"y\" or \"n\"): ")
		ans := ""
		n, err := fmt.Scanln(&ans)
		if n != 1 || err != nil {
			fmt.Printf("Error in input: %v\n", err)
		}
		switch ans {
		case "y":
			userScreen.generateShips()
			userScreen.showScreen()
		case "n":
			userScreen.setUserShips()
		}
		break
	}

	//game cycle
	for {
		if userScreen.ownAlive <= 0 {
			fmt.Println("You lose!")
			break
		} else if userScreen.enemyAlive <= 0 {
			fmt.Println("You win!")
			break
		}
		// player shoot cycle
		for {
			fmt.Println("Choose coordinates to shoot \"{x} {y}\": ")
			n, err := fmt.Scanln(&x, &y)
			if n != 2 || err != nil {
				fmt.Printf("Error in input: %v\n", err)
				continue
			}

			res := botScreen.shoot(x, y)
			if res {
				userScreen.enemyField[y][x] = HavocTile
				userScreen.enemyAlive -= 1
				fmt.Println("You hit!")
				userScreen.showScreen()
				continue
			} else {
				userScreen.enemyField[y][x] = MissTile
				fmt.Println("You miss!")
			}
			break
		}
		userScreen.showScreen()
		// bot shoot cycle
		for {
			x, y = botScreen.nextShoot()

			res := userScreen.shoot(x, y)
			if res {
				botScreen.enemyField[y][x] = HavocTile
				botScreen.enemyAlive -= 1
				continue
			} else {
				botScreen.enemyField[y][x] = MissTile
			}
			break
		}

	}
}

func deferFunc() {
	fmt.Print("Push return to quit the program")
	test, _ := fmt.Scanln()
	fmt.Println(test)
}
