package contains

// Erwartet einen String und eine Liste von Strings.
// Gibt true zurück, wenn der String in der Liste enthalten ist, ansonsten false.
func Contains(strings []string, search string) bool {
	// TODO
	if len(strings) == 0 {
		return false
	}
	for _, s := range strings {

		if s == search {
			return true
		}

	}
	return false
}
