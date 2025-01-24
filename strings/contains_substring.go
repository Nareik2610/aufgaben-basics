package strings

// Erwartet zwei Strings s und t und prüft, ob t in s als Teilstring vorkommt.
func ContainsSubstring(s, t string) bool {
	// TODO

	for i := 0; i < len(s)-len(t)+1; i++ {
		if s[i:i+len(t)] == t {
			return true
		}
	}
	return false

}
