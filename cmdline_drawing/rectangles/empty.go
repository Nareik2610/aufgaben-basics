package rectangles

import "fmt"

// Erwartet zwei Seitenlängen `height` und `width`.
// Zeichnet ein Rechteck mit diesen Seitenlängen auf der Konsole.
// Der Rand des Rechtecks soll aus `#`-Zeichen bestehen, der Innenraum soll leer sein.
func DrawEmptyRectangle(height, width int) {
	// TODO
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			if row == 0 || row == height-1 || col == 0 || col == width-1 {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

}

// REMARKS
// - Diese Aufgabe ist eine etwas komplexere Variante der Aufgabe "Gefülltes Rechteck zeichnen".
//   Man geht dabei ähnlich vor, muss allerdings zusätzlich prüfen, ob man sich am Rand des Rechtecks befindet.
