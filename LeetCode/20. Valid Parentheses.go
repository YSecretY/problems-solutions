func isValid(s string) bool {
    var stack []rune
    mp := map[rune]rune{
        ')':'(',
        ']':'[',
        '}':'{',
    }

    for _, el := range s {
        if _, ok := mp[el]; ok {
            if len(stack) != 0 && stack[len(stack)-1] == mp[el] {
                stack = stack[:len(stack)-1]
            } else {
                return false
            }
        } else {
            stack = append(stack, el)
        }
    }

    return len(stack) == 0
}
