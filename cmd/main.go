package main

import (
	"fmt"
	"math"
	"math/rand/v2"
	"strconv"
	"time"
)

const (
	ESC string = "\x1b"

	Black         string = ESC + "[30m"
	BrightBlack   string = ESC + "[30;1m"
	Red           string = ESC + "[31m"
	BrightRed     string = ESC + "[31;1m"
	Green         string = ESC + "[32m"
	BrightGreen   string = ESC + "[32;1m"
	Yellow        string = ESC + "[33m"
	BrightYellow  string = ESC + "[33;1m"
	Blue          string = ESC + "[34m"
	BrightBlue    string = ESC + "[34;1m"
	Magenta       string = ESC + "[35m"
	BrightMagenta string = ESC + "[35;1m"
	Cyan          string = ESC + "[36m"
	BrightCyan    string = ESC + "[36;1m"
	White         string = ESC + "[37m"
	BrightWhite   string = ESC + "[37;1m"

	Color256 string = ESC + "[38;5;"
	BG256    string = ESC + "[48;5;"

	Reset   string = ESC + "[0m"
	NoColor string = ""

	TopLeft string = ESC + "[H"
	Break   string = "\n\r"
)

func main() {
	clearScreen()

	for range 500 {
		randomDraws(1)
		time.Sleep(time.Millisecond)
	}

	drawAt(1, 11, Break)

	// drawAt(2, 2, colored('Ÿ', Black))
	// drawAt(1, 2, colored('Ÿ', BrightBlack))
	// drawAt(4, 2, colored('Ÿ', Red))
	// drawAt(3, 2, colored('Ÿ', BrightRed))
	// drawAt(6, 2, colored('Ÿ', Green))
	// drawAt(5, 2, colored('Ÿ', BrightGreen))
	// drawAt(8, 2, colored('Ÿ', Yellow))
	// drawAt(7, 2, colored('Ÿ', BrightYellow))
	// drawAt(2, 3, colored('Ÿ', Blue))
	// drawAt(1, 3, colored('Ÿ', BrightBlue))
	// drawAt(4, 3, colored('Ÿ', Magenta))
	// drawAt(3, 3, colored('Ÿ', BrightMagenta))
	// drawAt(6, 3, colored('Ÿ', Cyan))
	// drawAt(5, 3, colored('Ÿ', BrightCyan))
	// drawAt(8, 3, colored('Ÿ', White))
	// drawAt(7, 3, colored('Ÿ', BrightWhite))
	// drawAt(5, 4, colored('Ÿ', NoColor))

	// fmt.Print("\n\n\r")

	// colorTest(Color256)

	// drawAt(20*4, 15,
	// 	BG256+"99m"+
	// 		Color256+"196m"+" 45"+
	// 		Reset+Break)

	// colorTest(BG256)

	radius := 25

	ycenter := 20
	xcenter := 70

	for i := range 360 {
		x := float64(xcenter) + float64(radius)*math.Sin(float64(i))
		y := float64(ycenter) + 0.5*float64(radius)*math.Cos(float64(i))
		drawAt(int(x), int(y), colored('-', BrightGreen))
		// time.Sleep(time.Millisecond * 100)
	}

	drawAt(1, ycenter*3, Break)
}

func clearScreen() {
	fmt.Print("\x1b[2J", TopLeft)
}

func drawAt(x, y int, char string) {
	fmt.Printf("\x1b[%d;%dH%s", y, x, char)
}

func colored(char rune, color string) string {
	if len(color) == 0 {
		return string(char)
	}
	return fmt.Sprintf("%s%c%s", color, char, Reset)
}

func colorTest(scape string) {
	for i := range 16 {
		for j := range 16 {
			index := i*16 + j
			code := strconv.Itoa(index)
			fmt.Print(scape + code + "m Ÿ")
		}
		fmt.Print(Reset, "\n\r")
	}
}

func randomDraws(offset int) {
	x := rand.IntN(30)
	y := rand.IntN(10)
	bg := rand.IntN(256)
	fg := rand.IntN(256)

	sbg := strconv.Itoa(bg)
	sfg := strconv.Itoa(fg)

	drawAt(x, y+offset,
		BG256+sbg+"m"+
			Color256+sfg+"m"+"Ÿ"+Reset)
}
