func longestCommonPrefix(strs []string) string {
    pref := strs[0]

    for i := 1; i < len(strs); i++ {
        count := 0
        for j := 0; j < len(strs[i]) && j < len(pref) && strs[i][j] == pref[j]; j++ {
            count++
        }
        pref = pref[:count]
    }

    return pref
}
