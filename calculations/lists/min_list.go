package lists

// Erwartet eine Liste von Zahlen und gibt das Minimum zurück.
// Wenn die Liste leer ist, wird 0 zurückgegeben.

func MinList(nums []int) int {
	// TODO
	if len(nums) == 0 {
		return 0
	}

	var min int = nums[0]

	for _, wert := range nums {

		if wert < min {

			min = wert

		}
	}

	return min
}
