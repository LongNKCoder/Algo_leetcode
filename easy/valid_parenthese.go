package main

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	stackCheck := make([]string, 0)
	mapParentOpen := map[string]string{
		"{": "}",
		"(": ")",
		"[": "]",
	}
	mapParentClose := map[string]string{
		"}": "{",
		")": "(",
		"]": "[",
	}
	for _, char := range s {
		if _, ok := mapParentOpen[string(char)]; ok {
			stackCheck = append(stackCheck, string(char))
		}
		if open, ok := mapParentClose[string(char)]; ok {
			if len(stackCheck) == 0 {
				return false
			}
			value := stackCheck[len(stackCheck)-1]
			if value == open {
				stackCheck = stackCheck[:len(stackCheck)-1]
			} else {
				return false
			}
		}
	}
	if len(stackCheck) == 0 {
		return true
	}
	return false
}

func main() {
	println(isValid("()"))
	println(isValid("()[]{}"))
	println(isValid("(]"))
	println(isValid("([)]"))
	println(isValid("{[]}"))
	println(isValid("]"))
	println(isValid("[:)"))
	println(isValid("(){}}{"))
	println(isValid("([}}])"))
}
