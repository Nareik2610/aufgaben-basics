package numbers

import "math"

// Erwartet eine Zahl n >= 1 und liefert die Anzahl der Teiler dieser Zahl zurück.
// TODO

func CountDivisors(n int) int {
	count := 0
	for i := 1; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			if i*i == n {
				count++
			} else {
				count += 2
			}
		}
	}
	return count
}
