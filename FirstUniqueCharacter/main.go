package main

func main() {

}

// algorítimo quadratica
func FirstChar(s string) int {
outher:
	for i, r := range s {
		for ii, rr := range s {
			if i != ii && r == rr {
				continue outher
			}
		}
	}
	return -1
}

// algorítimo linear
func FirstUniqChar(s string) int {
	m := make(map[rune]int)

	for _, c := range s {
		m[c]++
	}

	for i, r := range s {
		if m[r] == 1 {
			return i
		}
	}

	return -1
}

// algorítimo linear porém sem as constantes big(O)
func FirstUniqChar2(s string) int {
	count := make([]int, 26)

	for _, c := range s {
		count[c-'a']++
	}

	for i, char := range s {
		if count[char-'a'] == 1 {
			return i
		}
	}

	return -1
}
