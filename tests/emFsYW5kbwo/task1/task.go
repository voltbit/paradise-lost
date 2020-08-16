package main

import "fmt"

func Solution(s string) string {
	i := 0
	j := len(s) - 1
	if i == j {
		if s[0] == '?' {
			return "a"
		}
		return s
	}
	if j == -1 {
		return ""
	}
	pal := make([]rune, len(s))
	j = len(s) - 1
	for i := 0; i < len(s); i++ {
		if s[i] == s[j] && s[i] != '?' {
			pal[i], pal[j] = rune(s[i]), rune(s[j])
		} else if s[i] == '?' && s[j] != '?' {
			pal[i], pal[j] = rune(s[j]), rune(s[j])
		} else if s[i] != '?' && s[j] == '?' {
			pal[i], pal[j] = rune(s[i]), rune(s[i])
		} else if s[i] == '?' && s[j] == '?' {
			pal[i], pal[j] = 'a', 'a'
		} else {
			return "NO"
		}
		j--
	}
	return string(pal)
}

func main() {
	fmt.Println(Solution("?ab??a"))
	fmt.Println(Solution("bab??a"))
	fmt.Println(Solution("?a?"))
	fmt.Println(Solution("?"))
}
