package main

func palindromeCHecker(s string) bool {
	var length = len(s)

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[length-1-i] {
			return false
		}
	}
	return true
}

func main() {
	palindromeCHecker("madam")
	palindromeCHecker("hello")
}
