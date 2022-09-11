package main

type botScreen struct {
	Screen
}

func (sc *botScreen) nextShoot() (x int, y int) {
	x = rnd.Intn(FSize)
	y = rnd.Intn(FSize)
	return
}

func makeBotScreen() botScreen {
	var screen botScreen
	screen.Screen = makeScreen()
	return screen
}
