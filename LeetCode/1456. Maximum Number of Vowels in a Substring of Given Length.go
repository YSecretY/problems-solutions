func maxVowels(s string, k int) int {
    max := countVowels(s[0:k])
    cur := max

    for l, r := 1, k+1; r <= len(s); l, r = l+1, r+1 {
        if isVowel(s[r-1]) { cur++ }
        if isVowel(s[l-1]) { cur-- }
        if cur > max { max = cur }
    }

    return max
}

func countVowels(s string) int {
    count := 0
    mp := map[rune]bool{'a': true, 'i': true, 'o': true, 'e': true, 'u': true}
    for _, r := range s {
        if _, ok := mp[r]; ok {
            count++
        }
    }

    return count
}

func isVowel(ch byte) bool {
    return ch == 'a' || ch == 'i' || ch == 'o' || ch == 'e' || ch == 'u'
}
