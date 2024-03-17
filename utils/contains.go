package utils

func ArrayContain[T comparable](target T, inputs []T) bool {
	for _, input := range inputs {
		if target == input {
			return true
		}
	}
	return false
}
