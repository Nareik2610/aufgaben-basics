package lists

// Erwartet eine Liste von Zahlen und gibt das Minimum zurück.
// Wenn die Liste leer ist, wird 0 zurückgegeben.

func MinList(values []int) int {
	min := values[0]
	for _, wert := range values {

		if wert <= min {

			min = wert

		}

	}

}
