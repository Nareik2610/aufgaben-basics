package make_unique

import "fmt"

// Erwartet eine Liste von Strings.
// Hängt Zahlen an alle mehrfach vorkommenden Strings an, um sie eindeutig zu machen.
func MakeUnique(strings []string) {
	// TODO
	for i, el := range strings {
		count := 2
		for j := i + 1; j < len(strings); j++ {
			if strings[j] == el {
				strings[j] = fmt.Sprintf("%s_%d", el, count)
				count++
			}
		}
	}
}

// REMARKS
// - Dies ist eine ähnliche Aufgabe wie das Zählen von Vorkommen in einer Liste.
//   Die innere Schleife macht i.W. das gleiche wie die Funktion `Count` aus der Aufgabe `count`.
//   Zusätzlich wird der Wert von `count` an den String angehängt, um ihn eindeutig zu machen.
