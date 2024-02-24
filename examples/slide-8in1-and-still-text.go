/* example code to scroll text on an x module where x is specified from the command line arg
and if the number of characters is less than x, the text will not scroll
*/

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/alinke/go-max7219-pixelcade"
)

func main() {

	var err error
	var message string
	var loops int
	var modules int

	args := os.Args
	fmt.Println("Args length", len(args))
	if len(args) < 2 {
		fmt.Println("Usage: <message> <loops> <modules> [speed(ms)]")
		os.Exit(1)
	} else {
		if len(args) < 3 { // we only have message
			message = args[1]
			fmt.Println("Loops not entered, defaulting to 1")
			loops = 1
			fmt.Println("Num modules not entered, defaulting to 8")
			modules = 8
		} else { //let's assume we have loops and modules then
			message = args[1]

			loops, err = strconv.Atoi(args[2])
			if err != nil || loops < 1 {
				fmt.Println("Loops not entered, defaulting to 1")
				loops = 1
			}

			modules, err = strconv.Atoi(args[3])
			if err != nil || modules < 1 {
				fmt.Println("Num modules not entered, defaulting to 8")
				modules = 8
			}
		}
	}

	mtx := max7219.NewMatrix(8, max7219.RotateClockwiseInvert)
	err = mtx.Open(1, 1, 7)
	if err != nil {
		log.Fatal(err)
	}

	if len(message) <= modules {
		fmt.Println("Displaying without scrolling: ", message)
		startPos := (modules - len(message)) / 2
		for i, char := range message {
			mtx.OutputChar(i+startPos, max7219.FontLCD, char, true) //FontSinclair or FontLCD
		}
	} else {
		fmt.Printf("Scrolling message: %s, Loops: %d, Modules: %d, Delay: %dms\n", message, loops, modules)
		for i := 0; i < loops; i++ {
			mtx.SlideMessage(message, max7219.FontLCD, true, 25*time.Millisecond)
		}
	}

	mtx.Close()
}
