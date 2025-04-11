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

	C256  string = ESC + "[38;5;"
	BG256 string = ESC + "[48;5;"

	Reset   string = ESC + "[0m"
	NoColor string = ""

	TopLeft      string = ESC + "[H"
	CLEAR_SCREEN string = ESC + "[2J"

	LF    string = "\r"
	Break string = LF + "\n"
)

type drawData struct {
	x, y int
	data string
}

func main() {
	clearScreen()

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

	schan := make(chan drawData)
	exit1 := make(chan string)
	exit2 := make(chan string)

	ycenter := 15
	xcenter := 40

	go func() {
		radius := 25

		for i := range 360 {
			x := float64(xcenter) + float64(radius)*math.Sin(float64(i))
			y := float64(ycenter) + 0.5*float64(radius)*math.Cos(float64(i))
			drawAt(int(x), int(y), colored('-', BrightGreen))
			time.Sleep(time.Millisecond * 33)
		}
		exit1 <- " Circle "
	}()

	go func() {
		for range 500 {
			randomDraws(25, 10)
			time.Sleep(time.Millisecond * 33)
		}
		exit2 <- "random " + <-exit1
	}()

	go func() {
		for s := range schan {
			drawAt(s.x, s.y, s.data)
		}
	}()

	drawAt(0, ycenter, <-exit2)

	w := 10
	h := 10

	mapData := mapGen()

	for y := range h {
		for x := range w {
			index := y*w + x
			char := mapData.char[index]
			bgColored := mapData.bg_color[index] + char
			drawAt(x+4, y+20, bgColored+Reset)
		}
	}

	drawAt(1, ycenter*3, Break)
}

func clearScreen() {
	fmt.Print(CLEAR_SCREEN, TopLeft)
}

func drawAt(x, y int, char string) {
	fmt.Printf("%s[%d;%dH%s", ESC, y, x, char)
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

func randomDraws(offset_x, offset_y int) {
	x := rand.IntN(30)
	y := rand.IntN(10)
	bg := rand.IntN(256)
	fg := rand.IntN(256)

	sbg := strconv.Itoa(bg)
	sfg := strconv.Itoa(fg)

	drawAt(x+offset_x, y+offset_y,
		BG256+sbg+"m"+
			C256+sfg+"m"+"Ÿ"+Reset)
}

type mapData struct {
	bg_color []string
	fg_color []string
	char     []string
}

func mapGen() mapData {
	mapT := mapData{
		char: []string{
			" 0 ", " 1 ", " 2 ", " 3 ", " 4 ", " 5 ", " 6 ", " 7 ", " 8 ", " 9 ",
			" 1 ", " 2 ", " 3 ", " 4 ", " 5 ", " 6 ", " 7 ", " 8 ", " 9 ", " 0 ",
			" 2 ", " 3 ", " 4 ", " 5 ", " 6 ", " 7 ", " 8 ", " 9 ", " 0 ", " 1 ",
			" 3 ", " 4 ", " 5 ", " 6 ", " 7 ", " 8 ", " 9 ", " 0 ", " 1 ", " 2 ",
			" 4", " 5 ", " 6 ", " 7 ", " 8 ", " 9 ", " 0 ", " 1 ", " 2 ", " 3 ",
			" 5", " 6 ", " 7 ", " 8 ", " 9 ", " 0 ", " 1 ", " 2 ", " 3 ", " 4 ",
			" 6", " 7 ", " 8 ", " 9 ", " 0 ", " 1 ", " 2 ", " 3 ", " 4 ", " 5 ",
			" 7", " 8 ", " 9 ", " 0 ", " 1 ", " 2 ", " 3 ", " 4 ", " 5 ", " 6 ",
			" 8", " 9 ", " 0 ", " 1 ", " 2 ", " 3 ", " 4 ", " 5 ", " 6 ", " 7 ",
			" 9", " 0 ", " 1 ", " 2 ", " 3 ", " 4 ", " 5 ", " 6 ", " 7 ", " 8 ",
		},
		bg_color: []string{
			BG256 + "149m", BG256 + "149m", BG256 + "149m", BG256 + "149m", BG256 + "149m",
			BG256 + "149m", BG256 + "149m", BG256 + "149m", BG256 + "149m", BG256 + "149m",
			BG256 + "149m", BG256 + "149m", BG256 + "149m", BG256 + "149m", BG256 + "149m",
			BG256 + "149m", BG256 + "149m", BG256 + "149m", BG256 + "149m", BG256 + "149m",
			BG256 + "143m", BG256 + "143m", BG256 + "143m", BG256 + "143m", BG256 + "143m",
			BG256 + "143m", BG256 + "143m", BG256 + "143m", BG256 + "143m", BG256 + "143m",
			BG256 + "143m", BG256 + "143m", BG256 + "63m", BG256 + "63m", BG256 + "143m",
			BG256 + "143m", BG256 + "143m", BG256 + "63m", BG256 + "63m", BG256 + "63m",
			BG256 + "143m", BG256 + "143m", BG256 + "63m", BG256 + "63m", BG256 + "63m",
			BG256 + "143m", BG256 + "143m", BG256 + "63m", BG256 + "63m", BG256 + "143m",
			BG256 + "143m", BG256 + "143m", BG256 + "143m", BG256 + "143m", BG256 + "143m",
			BG256 + "143m", BG256 + "143m", BG256 + "143m", BG256 + "143m", BG256 + "143m",
			BG256 + "143m", BG256 + "143m", BG256 + "143m", BG256 + "143m", BG256 + "143m",
			BG256 + "143m", BG256 + "143m", BG256 + "143m", BG256 + "143m", BG256 + "143m",
			BG256 + "143m", BG256 + "143m", BG256 + "143m", BG256 + "143m", BG256 + "143m",
			BG256 + "143m", BG256 + "143m", BG256 + "143m", BG256 + "143m", BG256 + "143m",
			BG256 + "143m", BG256 + "143m", BG256 + "143m", BG256 + "143m", BG256 + "143m",
			BG256 + "143m", BG256 + "143m", BG256 + "143m", BG256 + "143m", BG256 + "143m",
			BG256 + "143m", BG256 + "143m", BG256 + "143m", BG256 + "143m", BG256 + "143m",
			BG256 + "143m", BG256 + "143m", BG256 + "143m", BG256 + "143m", BG256 + "143m",
		},
		fg_color: []string{},
	}

	return mapT
}
