/* example code to scroll text on an x module where x is specified from the command line arg
and if the number of characters is less than x, the text will not scroll
*/

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/alinke/go-max7219-pixelcade"
)

func main() {
	var message string
	var loops, modules int

	flag.StringVar(&message, "message", "", "Message to display")
	flag.IntVar(&loops, "loops", 1, "Number of loops")
	flag.IntVar(&modules, "modules", 8, "Number of max7219 modules")
	flag.Parse()

	if message == "" {
		fmt.Println("Usage: -message <message> -loops <loops> -modules <modules>")
		os.Exit(1)
	}

	if loops < 1 {
		fmt.Println("Invalid loop count. Defaulting to 1.")
		loops = 1
	}

	if modules < 1 {
		fmt.Println("Invalid number of modules. Defaulting to 8.")
		modules = 8
	}

	mtx := max7219.NewMatrix(8, max7219.RotateClockwiseInvert)
	err := mtx.Open(1, 1, 7)
	if err != nil {
		log.Fatal(err)
	}

	if len(message) <= modules {
		//fmt.Println("Displaying without scrolling: ", message)
		startPos := (modules - len(message)) / 2
		for i, char := range message {
			mtx.OutputChar(i+startPos, max7219.FontLCD, char, true) // FontSinclair or FontLCD
		}
	} else {
		//fmt.Printf("Scrolling message: %s, Loops: %d, Modules: %d\n", message, loops, modules)
		for i := 0; i < loops; i++ {
			mtx.SlideMessage(message, max7219.FontLCD, true, 25*time.Millisecond)
		}
	}

	mtx.Close()
}
