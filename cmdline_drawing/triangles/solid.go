package triangles

import "fmt"

// Erwartet eine Seitenlänge `length`.
// Zeichnet ein gleichschenkliges, rechtwinkliges Dreieck mit diesen Seitenlängen auf der Konsole.
// Das Dreieck soll komplett mit `#`-Zeichen gefüllt sein.
// Der rechte Winkel soll links unten liegen und die Seiten sollen
// vertikal bzw. horizontal verlaufen.
func DrawSolidTriangle(length int) {
	// TODO
	for row := 0; row < length; row++ {
		for col := 0; col <= row; col++ {
			fmt.Print("#")
		}
		fmt.Println()
	}
}
