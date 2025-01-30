package strings

// Erwartet zwei strings s1 und s2.
// Liefert einen String, der abwechselnd aus den Buchstaben des einen und des anderen
// Strings besteht.
func Zip(s1, s2 string) string {
	// TODO
	result := ""
	for i := 0; i < len(s1) || i < len(s2); i++ {
		if i < len(s1) {
			result += string(s1[i])
		}
		if i < len(s2) {
			result += string(s2[i])
		}
	}
	return result
}
