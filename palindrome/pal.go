package palindrome

func isPal(str string) bool {
	if len(str) <= 1 {
		return true
	}

	if str[0] == str[len(str)-1] {
		return isPal(remFirstLast(str))
	} else {
		return false
	}
}

func remFirstLast(str string) string {
	return str[1 : len(str)-1]
}
