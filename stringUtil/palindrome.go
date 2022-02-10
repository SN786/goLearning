package stringUtil

func IsPalindrome(s string) bool {
	var l int = len(s)
	i := 0
	j := l - 1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}