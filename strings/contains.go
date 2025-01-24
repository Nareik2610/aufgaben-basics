package strings

// Erwartet einen String s und einen Buchstaben c.
// Prüft, ob c in s vorkommt.
func Contains(s string, c byte) bool {
	// TODO

	for i := range s {

		if s[i] == c {
			return true
		}
	}
	return false
}
