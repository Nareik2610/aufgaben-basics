package rectangles

import "fmt"

// Erwartet zwei Seitenlängen `height` und `width`.
// Zeichnet ein Rechteck mit diesen Seitenlängen auf der Konsole.
// Die Zeichen für Rand und Füllung des Rechtecks werden als Parameter erwartet.
func DrawRectangle(width int, height int, randZeichen string, fuellZeichen string, inneresZeichen string) {
	// Zeichne die obere Grenze
	for i := 0; i < width; i++ {
		fmt.Print(randZeichen)
	}
	fmt.Println()

	// Zeichne die Seitenränder und die Füllung in der Mitte
	for i := 0; i < height-2; i++ {
		fmt.Print(randZeichen)
		for j := 0; j < width-2; j++ {
			// Inneres Muster hier: je nach Zeile und Spalte verschiedene Zeichen verwenden
			if i == (height-2)/2 && j == (width-2)/2 {
				fmt.Print(inneresZeichen) // Inneres Zeichen in der Mitte
			} else {
				fmt.Print(fuellZeichen)
			}
		}
		fmt.Println(randZeichen)
	}

	// Zeichne die untere Grenze
	if height > 1 {
		for i := 0; i < width; i++ {
			fmt.Print(randZeichen)
		}
		fmt.Println()
	}
}

func main() {
	randZeichen := "#"    // Zeichen für den Rand
	fuellZeichen := " "   // Zeichen für die Füllung
	inneresZeichen := "*" // Zeichen in der Mitte

	// Erstes Rechteck
	width1 := 7
	height1 := 5
	DrawRectangle(width1, height1, randZeichen, fuellZeichen, inneresZeichen)

	// Zweites Rechteck (3 x 5)
	width2 := 5
	height2 := 3
	DrawRectangle(width2, height2, randZeichen, fuellZeichen, inneresZeichen)
}

// REMARKS
// - Wenn Sie diese Aufgabe gelöst haben, können Sie die Aufgaben
//   für das Zeichnen von leeren und gefüllten Rechtecken sehr viel einfacher lösen.
