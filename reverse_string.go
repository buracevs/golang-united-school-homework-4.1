package reverse_string

func ReverseString(input string) (output string) {
	caracters := []rune(input)
	first := 0
	last := len(caracters) - 1
	for first < last {
		caracters[first], caracters[last] = caracters[last], caracters[first]
		first++
		last--
	}
	output = string(caracters)
	return output
}
