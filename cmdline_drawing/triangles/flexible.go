package triangles

import "fmt"

// Erwartet eine Seitenlänge `length`.
// Zeichnet ein gleichschenkliges, rechtwinkliges Dreieck mit diesen Seitenlängen auf der Konsole.
// Der rechte Winkel soll links unten liegen und die Seiten sollen
// vertikal bzw. horizontal verlaufen.
// Die Zeichen für Rand und Füllung des Rechtecks werden als Parameter erwartet.
func DrawTriangle(length int, inner, outer string) {
	// TODO
	for row := 0; row < length; row++ {
		for col := 0; col <= row; col++ {
			if row == length-1 || col == 0 || col == row {
				fmt.Print(outer)
			} else {
				fmt.Print(inner)
			}
		}
		fmt.Println()
	}
}
