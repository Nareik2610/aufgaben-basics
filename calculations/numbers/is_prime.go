package numbers

// Erwartet eine Zahl n und prüft, ob n eine Primzahl ist.
func IsPrime(n int) bool {
	// Zahl muss größer als 1 sein, um eine Primzahl zu sein
	if n <= 1 {
		return false
	}

	// Nur bis zur Quadratwurzel von n iterieren
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false // n ist teilbar, also keine Primzahl
		}
	}

	return true // n ist eine Primzahl
}
