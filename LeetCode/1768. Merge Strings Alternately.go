func mergeAlternately(word1 string, word2 string) string {    
    res := ""
    i, j := 0, 0

    for ; i < len(word1) && j < len(word2); i, j = i+1, j+1 {
        res += string(word1[i])
        res += string(word2[j])
    }

    res += word1[i:]
    res += word2[j:]

    return res
}
