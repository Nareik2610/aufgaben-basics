package rectangles

import "fmt"

// Erwartet eine Seitenlänge `length`.
// Zeichnet ein Quadrat mit dieser Seitenlänge auf der Konsole.
// Das Quadrat soll komplett mit `#`-Zeichen gefüllt sein.
func DrawSolidSquare(length int) {
	// TODO
	for row := 0; row < length; row++ {
		for col := 0; col < length; col++ {

			fmt.Print("#")

		}
		fmt.Println()
	}
}
