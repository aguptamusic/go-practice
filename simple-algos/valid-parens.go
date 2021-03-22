package algos

func ValidParens(s string) bool {
	matches := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	stack := make([]rune, 0)
	for _, c := range s {
		match, present := matches[c]
		if present {
			// character is closing brace

			// check that stack is not empty
			if len(stack) == 0 {
				return false
			}
			
			// pop matching open brace
			top := stack[len(stack) - 1]
			if top != match {
				// matching brace not found
				return false
			}
			stack = stack[:len(stack) - 1]
		} else {
			// character is open brace
			stack = append(stack, c)
		}
	}
	return len(stack) == 0 //every closing brace had a matching open brace
}
