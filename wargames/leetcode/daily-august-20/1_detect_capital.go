package main

func DetectCapital(s string) bool {
	count := 0
	for _, char := range s[1:] {
		if int(char) >= int('A') && int(char) <= int('Z') {
			count += 1
		}
	}

	if count == 0 {
		return true
	} else if (int(s[0]) >= int('A') && int(s[0]) <= int('Z')) && count == len(s)-1 {
		return true
	} else {
		return false
	}
}
