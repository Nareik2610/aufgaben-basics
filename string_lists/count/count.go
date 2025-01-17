package count

// Erwartet eine Liste von Strings sowie einen String, der gezählt werden soll.
// Liefer die Anzahl der Vorkommen des gesuchten Strings in der Liste.
func Count(strings []string, search string) int {
	// TODO
	i := 0
	for _, count := range strings {

		if search == count {

			i++

		}
	}
	return i
}
